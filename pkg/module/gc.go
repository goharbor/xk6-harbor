package module

import (
	"time"

	operation "github.com/goharbor/xk6-harbor/pkg/harbor/client/gc"
	"github.com/goharbor/xk6-harbor/pkg/harbor/models"
)

func (h *Harbor) StartGC() int64 {
	h.mustInitialized()

	params := operation.NewCreateGCScheduleParams().WithSchedule(&models.Schedule{
		Schedule: &models.ScheduleObj{Type: "Manual"},
		Parameters: map[string]interface{}{
			"dry_run":         false,
			"delete_untagged": true,
		},
	})

	res, err := h.api.GC.CreateGCSchedule(h.vu.Context(), params)
	Checkf(h.vu.Runtime(), err, "failed to start gc")

	return IDFromLocation(h.vu.Runtime(), res.Location)
}

func (h *Harbor) GetGC(id int64) *models.GCHistory {
	h.mustInitialized()

	params := operation.NewGetGCParams().WithGCID(id)

	res, err := h.api.GC.GetGC(h.vu.Context(), params)
	Checkf(h.vu.Runtime(), err, "failed to get gc %d", id)

	return res.Payload
}

func (h *Harbor) StartGCAndWait() {
	jobID := h.StartGC()

	for {
		gc := h.GetGC(jobID)

		if gc.JobStatus == "Success" {
			break
		} else if gc.JobStatus == "Error" || gc.JobStatus == "Stopped" {
			Throwf(h.vu.Runtime(), "expect Success but get %s for gc %d", gc.JobStatus, jobID)
		}

		time.Sleep(time.Second)
	}
}
