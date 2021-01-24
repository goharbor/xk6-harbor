package module

import (
	"context"
	"fmt"

	operation "github.com/heww/xk6-harbor/pkg/harbor/client/products"
	"github.com/heww/xk6-harbor/pkg/harbor/models"
)

func (h *Harbor) CreateUser(ctx context.Context, username string, passwords ...string) int64 {
	h.mustInitialized(ctx)

	password := "Harbor12345"
	if len(passwords) > 0 {
		password = passwords[0]
	}

	params := operation.NewPostUsersParams()
	params.WithUser(&models.User{
		Username: username,
		Email:    fmt.Sprintf("%s@goharbor.io", username),
		Password: password,
		Realname: username,
	})

	res, err := h.api.Products.PostUsers(ctx, params)
	Checkf(ctx, err, "failed to create user %s", username)

	return IDFromLocation(ctx, res.Location)
}
