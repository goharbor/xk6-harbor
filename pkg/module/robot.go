package module

import (
	operation "github.com/goharbor/xk6-harbor/pkg/harbor/client/robot"
	"github.com/goharbor/xk6-harbor/pkg/harbor/models"
)

func (h *Harbor) CreateRobot(robot models.RobotCreate) int64 {
	h.mustInitialized()

	params := operation.NewCreateRobotParams().WithRobot(&robot)

	res, err := h.api.Robot.CreateRobot(h.vu.Context(), params)
	Checkf(h.vu.Runtime(), err, "failed to create robot %s", robot.Name)

	return IDFromLocation(h.vu.Runtime(), res.Location)
}
