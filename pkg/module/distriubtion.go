package module

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/containerd/containerd/content"
	"github.com/containerd/containerd/remotes"
	"github.com/containerd/containerd/remotes/docker"
	"github.com/deislabs/oras/pkg/oras"
	"github.com/dop251/goja"
	"github.com/google/uuid"
	"github.com/heww/xk6-harbor/pkg/util"
	"github.com/opencontainers/go-digest"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	log "github.com/sirupsen/logrus"
	"go.k6.io/k6/js/common"
)

type GetCatalogQuery struct {
	N    int    `js:"n"`
	Last string `js:"last"`
}

func (h *Harbor) GetCatalog(ctx context.Context, args ...goja.Value) map[string]interface{} {
	h.mustInitialized(ctx)

	var param GetCatalogQuery
	if len(args) > 0 {
		rt := common.GetRuntime(ctx)
		if err := rt.ExportTo(args[0], &param); err != nil {
			common.Throw(common.GetRuntime(ctx), err)
		}
	}

	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s://%s/v2/_catalog", h.option.Scheme, h.option.Host), nil)
	req.SetBasicAuth(h.option.Username, h.option.Password)

	q := req.URL.Query()
	if param.N != 0 {
		q.Add("n", strconv.Itoa(param.N))
	}

	if param.Last != "" {
		q.Add("last", param.Last)
	}

	req.URL.RawQuery = q.Encode()

	resp, err := h.httpClient.Do(req)
	Checkf(ctx, err, "failed to get catalog")
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	m := map[string]interface{}{}
	Checkf(ctx, dec.Decode(&m), "bad catalog")

	return m
}

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

type PullOption struct {
	Discard bool
}

func (h *Harbor) Pull(ctx context.Context, ref string, args ...goja.Value) {
	h.mustInitialized(ctx)

	params := PullOption{}
	ExportTo(ctx, &params, args...)

	var ingester content.Ingester
	if params.Discard {
		ingester = util.NewDiscardStore()
	} else {
		_, store := newLocalStore(ctx, util.GenerateRandomString(8))
		ingester = store
	}

	pullOpts := []oras.PullOpt{
		oras.WithPullEmptyNameAllowed(),
	}

	ref = h.getRef(ref)
	_, _, err := oras.Pull(ctx, h.makeResolver(ctx, args...), ref, ingester, pullOpts...)
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
