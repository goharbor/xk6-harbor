package module

import (
	"net/url"

	"github.com/dop251/goja"
	operation "github.com/goharbor/xk6-harbor/pkg/harbor/client/repository"
	"github.com/goharbor/xk6-harbor/pkg/harbor/models"
	"go.k6.io/k6/js/common"
)

func (h *Harbor) DeleteRepository(projectName, repositoryName string) {
	h.mustInitialized()

	params := operation.NewDeleteRepositoryParams()
	params.WithProjectName(projectName).WithRepositoryName(url.PathEscape(repositoryName))

	_, err := h.api.Repository.DeleteRepository(h.vu.Context(), params)
	Checkf(h.vu.Runtime(), err, "failed to delete repository %s/%s", projectName, repositoryName)
}

func (h *Harbor) GetRepository(projectName, repositoryName string) *models.Repository {
	h.mustInitialized()

	params := operation.NewGetRepositoryParams()
	params.WithProjectName(projectName)
	params.WithRepositoryName(repositoryName)

	res, err := h.api.Repository.GetRepository(h.vu.Context(), params)
	Checkf(h.vu.Runtime(), err, "failed to get repository %s/%s", projectName, repositoryName)

	return res.Payload
}

type ListRepositoriesResult struct {
	Repositories []*models.Repository `js:"repositories"`
	Total        int64                `js:"total"`
}

func (h *Harbor) ListRepositories(projectName string, args ...goja.Value) ListRepositoriesResult {
	h.mustInitialized()

	params := operation.NewListRepositoriesParams()
	params.WithProjectName(projectName)

	if len(args) > 0 {
		rt := h.vu.Runtime()
		if err := rt.ExportTo(args[0], params); err != nil {
			common.Throw(rt, err)
		}
	}

	res, err := h.api.Repository.ListRepositories(h.vu.Context(), params)
	Checkf(h.vu.Runtime(), err, "failed to list repositories of %s", projectName)

	return ListRepositoriesResult{
		Repositories: res.Payload,
		Total:        res.XTotalCount,
	}
}
