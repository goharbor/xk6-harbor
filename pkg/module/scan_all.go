package module

import (
	"context"

	operation "github.com/heww/xk6-harbor/pkg/harbor/client/scan_all"
	"github.com/heww/xk6-harbor/pkg/harbor/models"
)

func (h *Harbor) StartScanAll(ctx context.Context) {
	h.mustInitialized(ctx)

	params := operation.NewCreateScanAllScheduleParams().
		WithSchedule(&models.Schedule{
			Schedule: &models.ScheduleObj{
				Type: models.ScheduleObjTypeManual,
			},
		})

	_, err := h.api.ScanAll.CreateScanAllSchedule(ctx, params)
	Checkf(ctx, err, "failed to start scan all")
}

func (h *Harbor) GetScanAllMetrics(ctx context.Context) *models.Stats {
	h.mustInitialized(ctx)

	parmas := operation.NewGetLatestScanAllMetricsParams()

	res, err := h.api.ScanAll.GetLatestScanAllMetrics(ctx, parmas)
	Checkf(ctx, err, "failed to get metrics of scan all")

	return res.Payload
}
