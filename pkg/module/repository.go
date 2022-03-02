package module

import (
	"context"
	"net/url"

	"github.com/dop251/goja"
	operation "github.com/heww/xk6-harbor/pkg/harbor/client/repository"
	"github.com/heww/xk6-harbor/pkg/harbor/models"
	"go.k6.io/k6/js/common"
)

func (h *Harbor) DeleteRepository(ctx context.Context, projectName, repositoryName string) {
	h.mustInitialized(ctx)

	params := operation.NewDeleteRepositoryParams()
	params.WithProjectName(projectName).WithRepositoryName(url.PathEscape(repositoryName))

	_, err := h.api.Repository.DeleteRepository(ctx, params)
	Checkf(ctx, err, "failed to delete repository %s/%s", projectName, repositoryName)
}

func (h *Harbor) GetRepository(ctx context.Context, projectName, repositoryName string) *models.Repository {
	h.mustInitialized(ctx)

	params := operation.NewGetRepositoryParams()
	params.WithProjectName(projectName)
	params.WithRepositoryName(repositoryName)

	res, err := h.api.Repository.GetRepository(ctx, params)
	Checkf(ctx, err, "failed to get repository %s/%s", projectName, repositoryName)

	return res.Payload
}

type ListRepositoriesResult struct {
	Repositories []*models.Repository `js:"repositories"`
	Total        int64                `js:"total"`
}

func (h *Harbor) ListRepositories(ctx context.Context, projectName string, args ...goja.Value) ListRepositoriesResult {
	h.mustInitialized(ctx)

	params := operation.NewListRepositoriesParams()
	params.WithProjectName(projectName)

	if len(args) > 0 {
		rt := common.GetRuntime(ctx)
		if err := rt.ExportTo(args[0], params); err != nil {
			common.Throw(common.GetRuntime(ctx), err)
		}
	}

	res, err := h.api.Repository.ListRepositories(ctx, params)
	Checkf(ctx, err, "failed to list repositories of %s", projectName)

	return ListRepositoriesResult{
		Repositories: res.Payload,
		Total:        res.XTotalCount,
	}
}
