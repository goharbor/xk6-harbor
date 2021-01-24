package module

import (
	"context"

	"github.com/dop251/goja"
	"github.com/heww/xk6-harbor/pkg/harbor/client/products"
	operation "github.com/heww/xk6-harbor/pkg/harbor/client/replication"
	"github.com/heww/xk6-harbor/pkg/harbor/models"
	"github.com/loadimpact/k6/js/common"
)

func (h *Harbor) CreateReplicationPolicy(ctx context.Context, policy models.ReplicationPolicy) int64 {
	h.mustInitialized(ctx)

	params := products.NewPostReplicationPoliciesParams()
	params.WithPolicy(&policy)

	res, err := h.api.Products.PostReplicationPolicies(ctx, params)
	Checkf(ctx, err, "failed to create replication policy %s", params.Policy.Name)

	return IDFromLocation(ctx, res.Location)
}

func (h *Harbor) DeleteReplicationPolicy(ctx context.Context, id int64) {
	h.mustInitialized(ctx)

	params := products.NewDeleteReplicationPoliciesIDParams().WithID(id)

	_, err := h.api.Products.DeleteReplicationPoliciesID(ctx, params)
	Checkf(ctx, err, "failed to delete the replication policy %d", id)
}

type ListReplicationPoliciesResult struct {
	Policies []*models.ReplicationPolicy `js:"policies"`
	Total    int64                       `js:"total"`
}

func (h *Harbor) ListReplicationPolicies(ctx context.Context, args ...goja.Value) ListReplicationPoliciesResult {
	h.mustInitialized(ctx)

	params := products.NewGetReplicationPoliciesParams()
	if len(args) > 0 {
		rt := common.GetRuntime(ctx)
		if err := rt.ExportTo(args[0], params); err != nil {
			common.Throw(common.GetRuntime(ctx), err)
		}
	}

	res, err := h.api.Products.GetReplicationPolicies(ctx, params)
	Checkf(ctx, err, "failed to list replication policies	")

	return ListReplicationPoliciesResult{
		Policies: res.Payload,
		Total:    res.XTotalCount,
	}
}

func (h *Harbor) StartReplication(ctx context.Context, policyID int64) int64 {
	h.mustInitialized(ctx)

	params := operation.NewStartReplicationParams()
	params.WithExecution(&models.StartReplicationExecution{PolicyID: policyID})

	res, err := h.api.Replication.StartReplication(ctx, params)
	Checkf(ctx, err, "failed to start replication %d", policyID)

	return IDFromLocation(ctx, res.Location)
}

func (h *Harbor) GetReplicationExecution(ctx context.Context, executionID int64) *models.ReplicationExecution {
	h.mustInitialized(ctx)

	params := operation.NewGetReplicationExecutionParams()
	params.WithID(executionID)

	res, err := h.api.Replication.GetReplicationExecution(ctx, params)
	Checkf(ctx, err, "failed to get replication execution %d", executionID)

	return res.Payload
}
