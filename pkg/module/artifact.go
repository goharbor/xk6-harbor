package module

import (
	"context"
	"fmt"
	"net/url"
	"sync"

	"github.com/dop251/goja"
	"github.com/dustin/go-humanize"
	operation "github.com/heww/xk6-harbor/pkg/harbor/client/artifact"
	"github.com/heww/xk6-harbor/pkg/harbor/models"
	"github.com/heww/xk6-harbor/pkg/util"
	"github.com/loadimpact/k6/js/common"
	ocispec "github.com/opencontainers/image-spec/specs-go/v1"
	ants "github.com/panjf2000/ants/v2"
	log "github.com/sirupsen/logrus"
)

const DefaultPoolSise = 300

type PrepareArtifactsOption struct {
	ProjectName    string     `js:"projectName"`
	RepositoryName string     `js:"repositoryName"`
	ArtifactSize   goja.Value `js:"artifactSize"`
	ArtifactsCount int        `js:"artifactsCount"`
}

func (h *Harbor) PrepareArtifacts(ctx context.Context, option PrepareArtifactsOption) {
	h.mustInitialized(ctx)

	if option.ArtifactsCount <= 0 {
		Throwf(ctx, "artifacts count must greater than 0")
	}

	store := newContentStore(ctx, util.GenerateRandomString(8))
	defer store.Free(ctx)

	descriptors, err := store.GenerateMany(option.ArtifactSize, option.ArtifactsCount)
	Check(ctx, err)

	var wg sync.WaitGroup

	poolSize := DefaultPoolSise
	if option.ArtifactsCount < poolSize {
		poolSize = option.ArtifactsCount
	}

	p, _ := ants.NewPoolWithFunc(poolSize, func(i interface{}) {
		defer wg.Done()

		defer func() {
			if r := recover(); r != nil {
				log.Errorf("%s", r)
			}
		}()

		ix := i.(int)

		h.Push(ctx, PushOption{
			Ref:   fmt.Sprintf("%s/%s:tag-%d", option.ProjectName, option.RepositoryName, ix),
			Store: store,
			Blobs: []ocispec.Descriptor{*descriptors[ix]},
		})
	})
	defer p.Release()

	for i := 0; i < option.ArtifactsCount; i++ {
		wg.Add(1)
		_ = p.Invoke(i)
	}

	wg.Wait()
}

type ListArtifactsResult struct {
	Artifacts []*models.Artifact `js:"artifacts"`
	Total     int64              `js:"total"`
}

func (h *Harbor) ListArtifacts(ctx context.Context, projectName, repositoryName string, args ...goja.Value) ListArtifactsResult {
	h.mustInitialized(ctx)

	params := operation.NewListArtifactsParams()

	if len(args) > 0 {
		err := common.GetRuntime(ctx).ExportTo(args[0], params)
		Check(ctx, err)
	}

	params.WithProjectName(projectName).WithRepositoryName(url.PathEscape(repositoryName))

	res, err := h.api.Artifact.ListArtifacts(ctx, params)
	Checkf(ctx, err, "failed to list artifacts %s/%s", projectName, repositoryName)

	return ListArtifactsResult{
		Artifacts: res.Payload,
		Total:     res.XTotalCount,
	}
}

type PrepareArtifactTagsOption struct {
	ProjectName    string     `js:"projectName"`
	RepositoryName string     `js:"repositoryName"`
	ArtifactSize   goja.Value `js:"artifactSize"`
	TagsCount      int        `js:"tagsCount"`
}

func (h *Harbor) PrepareArtifactTags(ctx context.Context, option PrepareArtifactTagsOption) string {
	h.mustInitialized(ctx)

	if option.TagsCount <= 0 {
		Throwf(ctx, "artifact tags count must greater than 0")
	}

	size, err := humanize.ParseBytes(option.ArtifactSize.String())
	Check(ctx, err)

	store := newContentStore(ctx, util.GenerateRandomString(8))
	defer store.Free(ctx)

	descriptor, err := store.Generate(size)
	Check(ctx, err)

	digest := h.Push(ctx, PushOption{
		Ref:   fmt.Sprintf("%s/%s:latest", option.ProjectName, option.RepositoryName),
		Store: store,
		Blobs: []ocispec.Descriptor{*descriptor},
	})

	if option.TagsCount == 1 {
		return digest
	}

	createdTagsCount := option.TagsCount - 1

	var wg sync.WaitGroup

	poolSize := DefaultPoolSise
	if createdTagsCount < poolSize {
		poolSize = createdTagsCount
	}

	p, _ := ants.NewPoolWithFunc(poolSize, func(i interface{}) {
		ix := i.(int)

		h.CreateArtifactTag(ctx,
			option.ProjectName,
			option.RepositoryName,
			digest,
			fmt.Sprintf("tag-%d", ix),
		)

		wg.Done()
	})
	defer p.Release()

	for i := 0; i < createdTagsCount; i++ {
		wg.Add(1)
		_ = p.Invoke(i)
	}

	wg.Wait()

	return digest
}

func (h *Harbor) CreateArtifactTag(ctx context.Context, projectName, repositoryName, digestOrTag, newTag string) string {
	h.mustInitialized(ctx)

	params := operation.NewCreateTagParams()

	params.WithProjectName(projectName)
	params.WithRepositoryName(url.PathEscape(repositoryName))
	params.WithReference(digestOrTag)
	params.WithTag(&models.Tag{Name: newTag})

	res, err := h.api.Artifact.CreateTag(ctx, params)
	Checkf(ctx, err, "failed to create new tag %s to %s/%s", newTag, projectName, repositoryName)

	return res.Location
}

func (h *Harbor) ListArtifactTags(ctx context.Context, projectName, repositoryName, digestOrTag string, args ...goja.Value) []*models.Tag {
	h.mustInitialized(ctx)

	params := operation.NewListTagsParams()

	if len(args) > 0 {
		err := common.GetRuntime(ctx).ExportTo(args[0], params)
		Check(ctx, err)
	}

	params.WithProjectName(projectName)
	params.WithRepositoryName(url.PathEscape(repositoryName))
	params.WithReference(digestOrTag)

	res, err := h.api.Artifact.ListTags(ctx, params)
	Checkf(ctx, err, "failed to list artifact tags %s/%s", projectName, repositoryName)

	return res.Payload
}
