package module

import (
	"context"

	operation "github.com/heww/xk6-harbor/pkg/harbor/client/products"
	"github.com/heww/xk6-harbor/pkg/harbor/models"
)

func (h *Harbor) CreateProjectMember(ctx context.Context, projectID int64, userID int64, roleIDs ...int64) string {
	h.mustInitialized(ctx)

	roleID := int64(1)
	if len(roleIDs) > 0 {
		roleID = roleIDs[0]
	}

	params := operation.NewPostProjectsProjectIDMembersParams()
	params.WithProjectID(projectID).WithProjectMember(&models.ProjectMember{
		MemberUser: &models.UserEntity{UserID: userID},
		RoleID:     roleID,
	})

	res, err := h.api.Products.PostProjectsProjectIDMembers(ctx, params)
	Checkf(ctx, err, "failed to create project member for project %d", projectID)

	return res.Location
}
