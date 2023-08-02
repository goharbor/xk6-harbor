package module

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/containerd/containerd/remotes"
	"github.com/containerd/containerd/remotes/docker"
	orascontent "github.com/deislabs/oras/pkg/content"
	"github.com/deislabs/oras/pkg/oras"
	"github.com/dop251/goja"
	"github.com/goharbor/xk6-harbor/pkg/util"
	"github.com/google/uuid"
	"github.com/opencontainers/go-digest"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	log "github.com/sirupsen/logrus"
	"go.k6.io/k6/js/common"
)

type GetCatalogQuery struct {
	N    int    `js:"n"`
	Last string `js:"last"`
}

func (h *Harbor) GetCatalog(args ...goja.Value) map[string]interface{} {
	h.mustInitialized()

	var param GetCatalogQuery
	if len(args) > 0 {
		rt := h.vu.Runtime()
		if err := rt.ExportTo(args[0], &param); err != nil {
			common.Throw(rt, err)
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
	Checkf(h.vu.Runtime(), err, "failed to get catalog")
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)

	m := map[string]interface{}{}
	Checkf(h.vu.Runtime(), dec.Decode(&m), "bad catalog")

	return m
}

func (h *Harbor) GetManifest(ref string) map[string]interface{} {
	h.mustInitialized()

	resolver := h.makeResolver()

	ref = h.getRef(ref)
	_, desc, err := resolver.Resolve(h.vu.Context(), ref)
	Checkf(h.vu.Runtime(), err, "failed to head the manifest")

	fetcher, err := resolver.Fetcher(h.vu.Context(), ref)
	Checkf(h.vu.Runtime(), err, "failed to create fetcher")

	rc, err := fetcher.Fetch(h.vu.Context(), desc)
	Checkf(h.vu.Runtime(), err, "failed to get the manifest")

	defer rc.Close()

	dec := json.NewDecoder(rc)

	m := map[string]interface{}{}
	Checkf(h.vu.Runtime(), dec.Decode(&m), "bad manifest")

	return m
}

type PullOption struct {
	Discard bool
}

func (h *Harbor) Pull(ref string, args ...goja.Value) {
	h.mustInitialized()

	params := PullOption{}
	ExportTo(h.vu.Runtime(), &params, args...)

	var store orascontent.ProvideIngester
	if params.Discard {
		store = newDiscardStore()
	} else {
		_, l := newLocalStore(h.vu.Runtime(), util.GenerateRandomString(8))
		store = l
	}

	pullOpts := []oras.PullOpt{
		oras.WithPullEmptyNameAllowed(),
		oras.WithContentProvideIngester(store),
	}

	ref = h.getRef(ref)
	_, _, err := oras.Pull(h.vu.Context(), h.makeResolver(args...), ref, store, pullOpts...)
	Checkf(h.vu.Runtime(), err, "failed to pull %s", ref)
}

type PushOption struct {
	Ref   string
	Store *ContentStore
	Blobs []ocispec.Descriptor
}

func (h *Harbor) Push(option PushOption, args ...goja.Value) string {
	h.mustInitialized()

	resolver := h.makeResolver(args...)
	ref := h.getRef(option.Ref)

	// this config makes the harbor identify the artifact as image
	configBytes, _ := json.Marshal(map[string]interface{}{"User": uuid.New().String()})
	config := ocispec.Descriptor{
		MediaType: ocispec.MediaTypeImageConfig,
		Digest:    digest.FromBytes(configBytes),
		Size:      int64(len(configBytes)),
	}

	_, err := writeBlob(option.Store.RootPath, configBytes)
	Checkf(h.vu.Runtime(), err, "faied to prepare the config for the %s", ref)

	manifest, err := oras.Push(h.vu.Context(), resolver, ref, option.Store.Store, option.Blobs, oras.WithConfig(config))
	Checkf(h.vu.Runtime(), err, "failed to push %s", ref)

	return manifest.Digest.String()
}

func (h *Harbor) getRef(ref string) string {
	if !strings.HasPrefix(ref, h.option.Host) {
		return h.option.Host + "/" + ref
	}

	return ref
}

func (h *Harbor) makeResolver(args ...goja.Value) remotes.Resolver {
	h.mustInitialized()

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
