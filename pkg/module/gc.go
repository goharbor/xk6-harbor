package module

import (
	"context"
	"time"

	operation "github.com/heww/xk6-harbor/pkg/harbor/client/gc"
	"github.com/heww/xk6-harbor/pkg/harbor/models"
)

func (h *Harbor) StartGC(ctx context.Context) int64 {
	h.mustInitialized(ctx)

	params := operation.NewCreateGCScheduleParams().WithSchedule(&models.Schedule{
		Schedule: &models.ScheduleObj{Type: "Manual"},
		Parameters: map[string]interface{}{
			"dry_run":         false,
			"delete_untagged": true,
		},
	})

	res, err := h.api.GC.CreateGCSchedule(ctx, params)
	Checkf(ctx, err, "failed to start gc")

	return IDFromLocation(ctx, res.Location)
}

func (h *Harbor) GetGC(ctx context.Context, id int64) *models.GCHistory {
	h.mustInitialized(ctx)

	params := operation.NewGetGCParams().WithGCID(id)

	res, err := h.api.GC.GetGC(ctx, params)
	Checkf(ctx, err, "failed to get gc %d", id)

	return res.Payload
}

func (h *Harbor) StartGCAndWait(ctx context.Context) {
	jobID := h.StartGC(ctx)

	for {
		gc := h.GetGC(ctx, jobID)

		if gc.JobStatus == "Success" {
			break
		} else if gc.JobStatus == "Error" || gc.JobStatus == "Stopped" {
			Throwf(ctx, "expect Success but get %s for gc %d", gc.JobStatus, jobID)
		}

		time.Sleep(time.Second)
	}
}
