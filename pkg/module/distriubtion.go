package module

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/containerd/containerd/remotes"
	"github.com/containerd/containerd/remotes/docker"
	"github.com/deislabs/oras/pkg/oras"
	"github.com/dop251/goja"
	"github.com/google/uuid"
	"github.com/heww/xk6-harbor/pkg/util"
	"github.com/opencontainers/go-digest"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	log "github.com/sirupsen/logrus"
)

func (h *Harbor) GetManifest(ctx context.Context, ref string) map[string]interface{} {
	h.mustInitialized(ctx)

	resolver := h.makeResolver(ctx)

	ref = h.getRef(ref)
	_, desc, err := resolver.Resolve(ctx, ref)
	Checkf(ctx, err, "failed to head the manifest")

	fetcher, err := resolver.Fetcher(ctx, ref)
	Checkf(ctx, err, "failed to create fetcher")

	rc, err := fetcher.Fetch(ctx, desc)
	Checkf(ctx, err, "failed to get the manifest")

	defer rc.Close()

	dec := json.NewDecoder(rc)

	m := map[string]interface{}{}
	Checkf(ctx, dec.Decode(&m), "bad manifest")

	return m
}

func (h *Harbor) Pull(ctx context.Context, ref string, args ...goja.Value) {
	h.mustInitialized(ctx)

	resolver := h.makeResolver(ctx, args...)
	store := newContentStore(ctx, util.GenerateRandomString(8))

	pullOpts := []oras.PullOpt{
		oras.WithPullEmptyNameAllowed(),
		oras.WithContentProvideIngester(store.Store),
	}

	ref = h.getRef(ref)
	_, _, err := oras.Pull(ctx, resolver, ref, store.Store, pullOpts...)
	Checkf(ctx, err, "failed to pull %s", ref)
}

type PushOption struct {
	Ref   string
	Store *ContentStore
	Blobs []ocispec.Descriptor
}

func (h *Harbor) Push(ctx context.Context, option PushOption, args ...goja.Value) string {
	h.mustInitialized(ctx)

	resolver := h.makeResolver(ctx, args...)
	ref := h.getRef(option.Ref)

	// this config makes the harbor identify the artifact as image
	configBytes, _ := json.Marshal(map[string]interface{}{"User": uuid.New().String()})
	config := ocispec.Descriptor{
		MediaType: ocispec.MediaTypeImageConfig,
		Digest:    digest.FromBytes(configBytes),
		Size:      int64(len(configBytes)),
	}

	_, err := writeBlob(option.Store.RootPath, configBytes)
	Checkf(ctx, err, "faied to prepare the config for the %s", ref)

	manifest, err := oras.Push(ctx, resolver, ref, option.Store.Store, option.Blobs, oras.WithConfig(config))
	Checkf(ctx, err, "failed to push %s", ref)

	return manifest.Digest.String()
}

func (h *Harbor) getRef(ref string) string {
	if !strings.HasPrefix(ref, h.option.Host) {
		return h.option.Host + "/" + ref
	}

	return ref
}

func (h *Harbor) makeResolver(ctx context.Context, args ...goja.Value) remotes.Resolver {
	h.mustInitialized(ctx)

	log.StandardLogger().SetLevel(log.ErrorLevel)

	var transport http.RoundTripper
	if h.option.Insecure {
		transport = util.NewInsecureTransport()
	} else {
		transport = util.NewDefaultTransport()
	}

	client := &http.Client{Transport: transport}

	authorizer := docker.NewAuthorizer(client, func(host string) (string, string, error) {
		if host == h.option.Host {
			return h.option.Username, h.option.Password, nil
		}

		return "", "", nil
	})

	plainHTTP := func(host string) (bool, error) {
		if host == h.option.Host {
			return h.option.Scheme == "http", nil
		}

		return false, nil // default is https
	}

	return docker.NewResolver(docker.ResolverOptions{
		Hosts: docker.ConfigureDefaultRegistries(
			docker.WithAuthorizer(authorizer),
			docker.WithClient(client),
			docker.WithPlainHTTP(plainHTTP),
		),
	})
}
