package module

import (
	operation "github.com/goharbor/xk6-harbor/pkg/harbor/client/scan_all"
	"github.com/goharbor/xk6-harbor/pkg/harbor/models"
)

func (h *Harbor) StartScanAll() {
	h.mustInitialized()

	params := operation.NewCreateScanAllScheduleParams().
		WithSchedule(&models.Schedule{
			Schedule: &models.ScheduleObj{
				Type: models.ScheduleObjTypeManual,
			},
		})

	_, err := h.api.ScanAll.CreateScanAllSchedule(h.vu.Context(), params)
	Checkf(h.vu.Runtime(), err, "failed to start scan all")
}

func (h *Harbor) GetScanAllMetrics() *models.Stats {
	h.mustInitialized()

	parmas := operation.NewGetLatestScanAllMetricsParams()

	res, err := h.api.ScanAll.GetLatestScanAllMetrics(h.vu.Context(), parmas)
	Checkf(h.vu.Runtime(), err, "failed to get metrics of scan all")

	return res.Payload
}
