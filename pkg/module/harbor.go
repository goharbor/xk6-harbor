package module

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/dop251/goja"
	rtclient "github.com/go-openapi/runtime/client"
	"github.com/goharbor/xk6-harbor/pkg/harbor/client"
	"github.com/goharbor/xk6-harbor/pkg/util"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
)

var DefaultRootPath = filepath.Join(os.TempDir(), "harbor")

func init() {
	modules.Register("k6/x/harbor", New())

	rootPath := os.Getenv("HARBOR_ROOT")
	if rootPath != "" {
		DefaultRootPath = rootPath
	}

	if err := os.MkdirAll(DefaultRootPath, 0755); err != nil {
		panic(err)
	}
}

var (
	varTrue = true
)

type (
	Option struct {
		Scheme   string // http or https
		Host     string
		Username string
		Password string
		Insecure bool // Allow insecure server connections when using SSL
	}
	Harbor struct {
		vu          modules.VU
		httpClient  *http.Client
		api         *client.HarborAPI
		option      *Option
		initialized bool
		once        sync.Once
	}

	RootModule struct{}
	Module     struct {
		*Harbor
	}
)

var (
	_ modules.Instance = &Module{}
	_ modules.Module   = &RootModule{}
)

// New creates a new instance of the root module.
func New() *RootModule {
	return &RootModule{}
}

func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &Module{
		Harbor: &Harbor{vu: vu},
	}
}

func (m *Module) Exports() modules.Exports {
	return modules.Exports{
		Named: map[string]interface{}{
			"Harbor":       m.newHarbor,
			"ContentStore": m.newContentStore,
		},
	}
}

func (m *Module) newHarbor(c goja.ConstructorCall) *goja.Object {
	rt := m.vu.Runtime()

	if len(c.Arguments) > 0 {
		m.Initialize(c)
	}

	return rt.ToValue(m.Harbor).ToObject(rt)
}

func (m *Module) newContentStore(c goja.ConstructorCall) *goja.Object {
	rt := m.vu.Runtime()

	store := newContentStore(m.vu.Runtime(), c.Arguments[0].String())

	return rt.ToValue(store).ToObject(rt)
}

func (h *Harbor) Initialize(c goja.ConstructorCall) {
	if h.initialized {
		common.Throw(h.vu.Runtime(), errors.New("harbor module initialized"))
	}

	h.once.Do(func() {
		opt := &Option{
			Scheme:   util.GetEnv("HARBOR_SCHEME", "https"),
			Host:     util.GetEnv("HARBOR_HOST", ""),
			Username: util.GetEnv("HARBOR_USERNAME", "admin"),
			Password: util.GetEnv("HARBOR_PASSWORD", "Harbor12345"),
			Insecure: false,
		}

		if len(c.Arguments) > 0 {
			if c.Arguments[0] != nil && !goja.IsUndefined(c.Arguments[0]) && !goja.IsNull(c.Arguments[0]) {
				rt := h.vu.Runtime()

				err := rt.ExportTo(c.Arguments[0], opt)
				Checkf(h.vu.Runtime(), err, "failed to parse the option")
			}
		}

		if opt.Host == "" {
			h.vu.Runtime().Interrupt("harbor host is required in initialization")
			return
		}

		opt.Scheme = strings.ToLower(opt.Scheme)
		if opt.Scheme != "https" && opt.Scheme != "http" {
			h.vu.Runtime().Interrupt(fmt.Sprintf("invalid harbor scheme %s", opt.Scheme))
			return
		}

		opt.Host = strings.TrimSuffix(opt.Host, "/")

		rawURL := fmt.Sprintf("%s://%s%s", opt.Scheme, opt.Host, client.DefaultBasePath)
		u, err := url.Parse(rawURL)
		if err != nil {
			common.Throw(h.vu.Runtime(), err)
		}

		config := client.Config{URL: u}

		if opt.Username != "" && opt.Password != "" {
			config.AuthInfo = rtclient.BasicAuth(opt.Username, opt.Password)
		}

		if opt.Insecure {
			config.Transport = util.NewInsecureTransport()
		} else {
			config.Transport = util.NewDefaultTransport()
		}

		h.api = client.New(config)
		h.option = opt
		h.httpClient = &http.Client{Transport: config.Transport}
		h.initialized = true
	})
}

func (h *Harbor) Free() {
	err := os.RemoveAll(DefaultRootPath)
	if err != nil {
		panic(h.vu.Runtime().NewGoError(err))
	}
}

func (h *Harbor) mustInitialized() {
	if !h.initialized {
		common.Throw(h.vu.Runtime(), errors.New("harbor module not initialized"))
	}
}
