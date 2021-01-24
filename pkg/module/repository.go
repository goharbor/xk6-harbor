package module

import (
	"context"
	"net/url"

	operation "github.com/heww/xk6-harbor/pkg/harbor/client/repository"
	"github.com/heww/xk6-harbor/pkg/harbor/models"
)

func (h *Harbor) DeleteRepository(ctx context.Context, projectName, repositoryName string) {
	h.mustInitialized(ctx)

	params := operation.NewDeleteRepositoryParams()
	params.WithProjectName(projectName).WithRepositoryName(url.PathEscape(repositoryName))

	_, err := h.api.Repository.DeleteRepository(ctx, params)
	Checkf(ctx, err, "failed to delete repository %s/%s", projectName, repositoryName)
}

type ListRepositoriesResult struct {
	Repositories []*models.Repository `js:"repositories"`
	Total        int64                `js:"total"`
}

func (h *Harbor) ListRepositories(ctx context.Context, projectName string, query ...Query) ListRepositoriesResult {
	h.mustInitialized(ctx)

	params := operation.NewListRepositoriesParams()
	params.WithProjectName(projectName)

	if len(query) > 0 {
		q := query[0]

		params.Page = q.Page
		params.PageSize = q.PageSize
		params.Q = q.Q
	}

	res, err := h.api.Repository.ListRepositories(ctx, params)
	Checkf(ctx, err, "failed to list repositories of %s", projectName)

	return ListRepositoriesResult{
		Repositories: res.Payload,
		Total:        res.XTotalCount,
	}
}
