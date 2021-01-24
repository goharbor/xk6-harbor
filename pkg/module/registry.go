package module

import (
	"context"

	"github.com/dop251/goja"
	operation "github.com/heww/xk6-harbor/pkg/harbor/client/products"
	"github.com/heww/xk6-harbor/pkg/harbor/models"
	"github.com/loadimpact/k6/js/common"
)

func (h *Harbor) CreateRegistry(ctx context.Context, r models.Registry) int64 {
	h.mustInitialized(ctx)

	params := operation.NewPostRegistriesParams().WithRegistry(&r)

	res, err := h.api.Products.PostRegistries(ctx, params)
	Checkf(ctx, err, "failed to create registry %s", params.Registry.Name)

	return IDFromLocation(ctx, res.Location)
}

func (h *Harbor) DeleteRegistry(ctx context.Context, id int64) {
	h.mustInitialized(ctx)

	params := operation.NewDeleteRegistriesIDParams().WithID(id)

	_, err := h.api.Products.DeleteRegistriesID(ctx, params)
	Checkf(ctx, err, "failed to delete registry %d", id)
}

type ListRegistriesResult struct {
	Registries []*models.Registry `js:"registries"`
}

func (h *Harbor) ListRegistries(ctx context.Context, args ...goja.Value) ListRegistriesResult {
	h.mustInitialized(ctx)

	params := operation.NewGetRegistriesParams()
	if len(args) > 0 {
		rt := common.GetRuntime(ctx)
		if err := rt.ExportTo(args[0], params); err != nil {
			common.Throw(common.GetRuntime(ctx), err)
		}
	}

	res, err := h.api.Products.GetRegistries(ctx, params)
	Checkf(ctx, err, "failed to list registries")

	return ListRegistriesResult{
		Registries: res.Payload,
	}
}
