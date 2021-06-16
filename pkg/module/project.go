package module

import (
	"context"
	"strings"

	"github.com/dop251/goja"
	operation "github.com/heww/xk6-harbor/pkg/harbor/client/project"
	"github.com/heww/xk6-harbor/pkg/harbor/client/repository"
	"github.com/heww/xk6-harbor/pkg/harbor/models"
	"github.com/heww/xk6-harbor/pkg/util"
	log "github.com/sirupsen/logrus"
	"go.k6.io/k6/js/common"
)

func (h *Harbor) CreateProject(ctx context.Context, body goja.Value) string {
	h.mustInitialized(ctx)

	rt := common.GetRuntime(ctx)
	var project models.ProjectReq
	err := rt.ExportTo(body, &project)
	Check(ctx, err)

	params := operation.NewCreateProjectParams()
	params.WithProject(&project).WithXResourceNameInLocation(&varTrue)

	res, err := h.api.Project.CreateProject(ctx, params)
	Checkf(ctx, err, "failed to create project %s", params.Project.ProjectName)

	return NameFromLocation(ctx, res.Location)
}

func (h *Harbor) DeleteProject(ctx context.Context, projectName string, args ...goja.Value) {
	h.mustInitialized(ctx)

	var force bool
	if len(args) > 0 {
		force = args[0].ToBoolean()
	}

	if force {
		pageSize := 20

		params := repository.NewListRepositoriesParams().WithProjectName(projectName)
		params.Page = util.Int64(1)
		params.PageSize = util.Int64(int64(pageSize))

		for {
			resp, err := h.api.Repository.ListRepositories(ctx, params)
			Checkf(ctx, err, "failed to list repositories for page %d", *params.Page)

			for _, repo := range resp.Payload {
				repoName := strings.TrimPrefix(repo.Name, projectName+"/")
				h.DeleteRepository(ctx, projectName, repoName)
			}

			if len(resp.Payload) < pageSize {
				break
			}
		}
	}

	params := operation.NewDeleteProjectParams()
	params.WithProjectNameOrID(projectName).WithXIsResourceName(&varTrue)

	_, err := h.api.Project.DeleteProject(ctx, params)
	Checkf(ctx, err, "failed to delete project %s", projectName)
}

func (h *Harbor) DeleteAllProjects(ctx context.Context, excludeProjects []string) {
	h.mustInitialized(ctx)

	rt := common.GetRuntime(ctx)

	m := make(map[string]bool, len(excludeProjects))
	for _, projectName := range excludeProjects {
		m[projectName] = true
	}

	page := 1
	pageSize := 10
	for {
		arg := map[string]interface{}{"page": page, "page_size": pageSize}

		result := h.ListProjects(ctx, rt.ToValue(arg))

		projects := result.Projects

		deleted := 0
		for _, project := range projects {
			if m[project.Name] {
				continue
			}

			func() {
				defer func() {
					if r := recover(); r != nil {
						log.Warnf("%v", r)
					}
				}()

				h.DeleteProject(ctx, project.Name, rt.ToValue(project.RepoCount+project.ChartCount > 0))
				deleted++
			}()
		}

		if len(projects) == 0 || len(projects) < pageSize {
			break
		}

		if deleted == 0 {
			page++
		}
	}
}

type ListProjectsResult struct {
	Projects []*models.Project `js:"projects"`
	Total    int64             `js:"total"`
}

func (h *Harbor) ListProjects(ctx context.Context, args ...goja.Value) ListProjectsResult {
	h.mustInitialized(ctx)

	params := operation.NewListProjectsParams()

	if len(args) > 0 {
		rt := common.GetRuntime(ctx)
		if err := rt.ExportTo(args[0], params); err != nil {
			common.Throw(common.GetRuntime(ctx), err)
		}
	}

	res, err := h.api.Project.ListProjects(ctx, params)
	Checkf(ctx, err, "failed to list projects")

	return ListProjectsResult{
		Projects: res.Payload,
		Total:    res.XTotalCount,
	}
}

type ListAuditLogsOfProjectResult struct {
	Logs  []*models.AuditLog `js:"logs"`
	Total int64              `js:"total"`
}

func (h *Harbor) ListAuditLogsOfProject(ctx context.Context, projectName string, args ...goja.Value) ListAuditLogsOfProjectResult {
	h.mustInitialized(ctx)

	params := operation.NewGetLogsParams().WithProjectName(projectName)

	if len(args) > 0 {
		rt := common.GetRuntime(ctx)
		if err := rt.ExportTo(args[0], params); err != nil {
			common.Throw(common.GetRuntime(ctx), err)
		}
	}

	res, err := h.api.Project.GetLogs(ctx, params)
	Checkf(ctx, err, "failed to list audit logs of project %s", projectName)

	return ListAuditLogsOfProjectResult{
		Logs:  res.Payload,
		Total: res.XTotalCount,
	}
}
