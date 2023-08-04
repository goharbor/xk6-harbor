package module

import (
	"github.com/dop251/goja"
	operation "github.com/goharbor/xk6-harbor/pkg/harbor/client/registry"
	"github.com/goharbor/xk6-harbor/pkg/harbor/models"
	"go.k6.io/k6/js/common"
)

func (h *Harbor) CreateRegistry(r models.Registry) int64 {
	h.mustInitialized()

	params := operation.NewCreateRegistryParams().WithRegistry(&r)

	res, err := h.api.Registry.CreateRegistry(h.vu.Context(), params)
	Checkf(h.vu.Runtime(), err, "failed to create registry %s", params.Registry.Name)

	return IDFromLocation(h.vu.Runtime(), res.Location)
}

func (h *Harbor) DeleteRegistry(id int64) {
	h.mustInitialized()

	params := operation.NewDeleteRegistryParams().WithID(id)

	_, err := h.api.Registry.DeleteRegistry(h.vu.Context(), params)
	Checkf(h.vu.Runtime(), err, "failed to delete registry %d", id)
}

type ListRegistriesResult struct {
	Registries []*models.Registry `js:"registries"`
}

func (h *Harbor) ListRegistries(args ...goja.Value) ListRegistriesResult {
	h.mustInitialized()

	params := operation.NewListRegistriesParams()
	if len(args) > 0 {
		rt := h.vu.Runtime()
		if err := rt.ExportTo(args[0], params); err != nil {
			common.Throw(rt, err)
		}
	}

	res, err := h.api.Registry.ListRegistries(h.vu.Context(), params)
	Checkf(h.vu.Runtime(), err, "failed to list registries")

	return ListRegistriesResult{
		Registries: res.Payload,
	}
}
