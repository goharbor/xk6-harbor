package module

import (
	"github.com/dop251/goja"
	operation "github.com/goharbor/xk6-harbor/pkg/harbor/client/auditlog"
	"github.com/goharbor/xk6-harbor/pkg/harbor/models"
	"go.k6.io/k6/js/common"
)

type ListAuditLogsResult struct {
	AuditLogs []*models.AuditLog `js:"logs"`
	Total     int64              `js:"total"`
}

func (h *Harbor) ListAuditLogs(args ...goja.Value) ListAuditLogsResult {
	h.mustInitialized()

	params := operation.NewListAuditLogsParams()

	if len(args) > 0 {
		rt := h.vu.Runtime()
		if err := rt.ExportTo(args[0], params); err != nil {
			common.Throw(rt, err)
		}
	}

	res, err := h.api.Auditlog.ListAuditLogs(h.vu.Context(), params)
	Checkf(h.vu.Runtime(), err, "failed to list audit logs")

	return ListAuditLogsResult{
		AuditLogs: res.Payload,
		Total:     res.XTotalCount,
	}
}
