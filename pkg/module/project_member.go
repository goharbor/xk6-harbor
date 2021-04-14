package module

import (
	"context"

	operation "github.com/heww/xk6-harbor/pkg/harbor/client/member"
	"github.com/heww/xk6-harbor/pkg/harbor/models"
)

func (h *Harbor) CreateProjectMember(ctx context.Context, projectName string, userID int64, roleIDs ...int64) string {
	h.mustInitialized(ctx)

	roleID := int64(1)
	if len(roleIDs) > 0 {
		roleID = roleIDs[0]
	}

	params := operation.NewCreateProjectMemberParams()
	params.WithProjectNameOrID(projectName).WithXIsResourceName(&varTrue).WithProjectMember(&models.ProjectMember{
		MemberUser: &models.UserEntity{UserID: userID},
		RoleID:     roleID,
	})

	res, err := h.api.Member.CreateProjectMember(ctx, params)
	Checkf(ctx, err, "failed to create project member for project %s", projectName)

	return res.Location
}
