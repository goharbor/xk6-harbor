package module

import (
	"github.com/dop251/goja"
	operation "github.com/goharbor/xk6-harbor/pkg/harbor/client/quota"
	"github.com/goharbor/xk6-harbor/pkg/harbor/models"
	"go.k6.io/k6/js/common"
)

type ListQuotasResult struct {
	Quotas []*models.Quota `js:"quotas"`
	Total  int64           `js:"total"`
}

func (h *Harbor) ListQuotas(args ...goja.Value) ListQuotasResult {
	h.mustInitialized()

	params := operation.NewListQuotasParams()

	if len(args) > 0 {
		rt := h.vu.Runtime()
		if err := rt.ExportTo(args[0], params); err != nil {
			common.Throw(rt, err)
		}
	}

	res, err := h.api.Quota.ListQuotas(h.vu.Context(), params)
	Checkf(h.vu.Runtime(), err, "failed to list quotas")

	return ListQuotasResult{
		Quotas: res.Payload,
		Total:  res.XTotalCount,
	}
}
