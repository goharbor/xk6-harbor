package module

import (
	"context"

	"github.com/dop251/goja"
	operation "github.com/heww/xk6-harbor/pkg/harbor/client/auditlog"
	"github.com/heww/xk6-harbor/pkg/harbor/models"
	"go.k6.io/k6/js/common"
)

type ListAuditLogsResult struct {
	AuditLogs []*models.AuditLog `js:"logs"`
	Total     int64              `js:"total"`
}

func (h *Harbor) ListAuditLogs(ctx context.Context, args ...goja.Value) ListAuditLogsResult {
	h.mustInitialized(ctx)

	params := operation.NewListAuditLogsParams()

	if len(args) > 0 {
		rt := common.GetRuntime(ctx)
		if err := rt.ExportTo(args[0], params); err != nil {
			common.Throw(common.GetRuntime(ctx), err)
		}
	}

	res, err := h.api.Auditlog.ListAuditLogs(ctx, params)
	Checkf(ctx, err, "failed to list audit logs")

	return ListAuditLogsResult{
		AuditLogs: res.Payload,
		Total:     res.XTotalCount,
	}
}
