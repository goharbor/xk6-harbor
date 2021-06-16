package module

import (
	"context"

	"github.com/dop251/goja"
	operation "github.com/heww/xk6-harbor/pkg/harbor/client/member"
	"github.com/heww/xk6-harbor/pkg/harbor/models"
	"go.k6.io/k6/js/common"
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

type ListProjectMembersResult struct {
	ProjectMembers []*models.ProjectMemberEntity `js:"projectMembers"`
	Total          int64                         `js:"total"`
}

func (h *Harbor) ListProjectMembers(ctx context.Context, projectName string, args ...goja.Value) ListProjectMembersResult {
	h.mustInitialized(ctx)

	params := operation.NewListProjectMembersParams()
	params.WithProjectNameOrID(projectName).WithXIsResourceName(&varTrue)

	if len(args) > 0 {
		rt := common.GetRuntime(ctx)
		if err := rt.ExportTo(args[0], params); err != nil {
			common.Throw(common.GetRuntime(ctx), err)
		}
	}

	res, err := h.api.Member.ListProjectMembers(ctx, params)
	Checkf(ctx, err, "failed to list project members of project %s", projectName)

	return ListProjectMembersResult{
		ProjectMembers: res.Payload,
		Total:          res.XTotalCount,
	}
}
