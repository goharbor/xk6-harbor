package module

import (
	"context"
	"fmt"

	"github.com/dop251/goja"
	operation "github.com/heww/xk6-harbor/pkg/harbor/client/user"
	"github.com/heww/xk6-harbor/pkg/harbor/models"
	"github.com/loadimpact/k6/js/common"
)

func (h *Harbor) CreateUser(ctx context.Context, username string, passwords ...string) int64 {
	h.mustInitialized(ctx)

	password := "Harbor12345"
	if len(passwords) > 0 {
		password = passwords[0]
	}

	params := operation.NewCreateUserParams()
	params.WithUserReq(&models.UserCreationReq{
		Username: username,
		Email:    fmt.Sprintf("%s@goharbor.io", username),
		Password: password,
		Realname: username,
	})

	res, err := h.api.User.CreateUser(ctx, params)
	Checkf(ctx, err, "failed to create user %s", username)

	return IDFromLocation(ctx, res.Location)
}

type ListUsersResult struct {
	Users []*models.UserResp `js:"users"`
	Total int64              `js:"total"`
}

func (h *Harbor) ListUsers(ctx context.Context, args ...goja.Value) ListUsersResult {
	h.mustInitialized(ctx)

	params := operation.NewListUsersParams()

	if len(args) > 0 {
		rt := common.GetRuntime(ctx)
		if err := rt.ExportTo(args[0], params); err != nil {
			common.Throw(common.GetRuntime(ctx), err)
		}
	}

	res, err := h.api.User.ListUsers(ctx, params)
	Checkf(ctx, err, "failed to list users")

	return ListUsersResult{
		Users: res.Payload,
		Total: res.XTotalCount,
	}
}

type SearchUsersResult struct {
	Users []*models.UserSearchRespItem `js:"users"`
	Total int64                        `js:"total"`
}

func (h *Harbor) SearchUsers(ctx context.Context, args ...goja.Value) SearchUsersResult {
	h.mustInitialized(ctx)

	params := operation.NewSearchUsersParams()

	if len(args) > 0 {
		rt := common.GetRuntime(ctx)
		if err := rt.ExportTo(args[0], params); err != nil {
			common.Throw(common.GetRuntime(ctx), err)
		}
	}

	res, err := h.api.User.SearchUsers(ctx, params)
	Checkf(ctx, err, "failed to list users")

	return SearchUsersResult{
		Users: res.Payload,
		Total: res.XTotalCount,
	}
}
