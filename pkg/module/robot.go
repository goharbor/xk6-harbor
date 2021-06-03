package module

import (
	"context"

	operation "github.com/heww/xk6-harbor/pkg/harbor/client/robot"
	"github.com/heww/xk6-harbor/pkg/harbor/models"
)

func (h *Harbor) CreateRobot(ctx context.Context, robot models.RobotCreate) int64 {
	h.mustInitialized(ctx)

	params := operation.NewCreateRobotParams().WithRobot(&robot)

	res, err := h.api.Robot.CreateRobot(ctx, params)
	Checkf(ctx, err, "failed to create robot %s", robot.Name)

	return IDFromLocation(ctx, res.Location)
}
