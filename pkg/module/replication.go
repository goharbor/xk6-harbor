package module

import (
	"github.com/dop251/goja"
	operation "github.com/goharbor/xk6-harbor/pkg/harbor/client/replication"
	"github.com/goharbor/xk6-harbor/pkg/harbor/models"
	"go.k6.io/k6/js/common"
)

func (h *Harbor) CreateReplicationPolicy(policy models.ReplicationPolicy) int64 {
	h.mustInitialized()

	params := operation.NewCreateReplicationPolicyParams()
	params.WithPolicy(&policy)

	res, err := h.api.Replication.CreateReplicationPolicy(h.vu.Context(), params)
	Checkf(h.vu.Runtime(), err, "failed to create replication policy %s", params.Policy.Name)

	return IDFromLocation(h.vu.Runtime(), res.Location)
}

func (h *Harbor) DeleteReplicationPolicy(id int64) {
	h.mustInitialized()

	params := operation.NewDeleteReplicationPolicyParams().WithID(id)

	_, err := h.api.Replication.DeleteReplicationPolicy(h.vu.Context(), params)
	Checkf(h.vu.Runtime(), err, "failed to delete the replication policy %d", id)
}

type ListReplicationPoliciesResult struct {
	Policies []*models.ReplicationPolicy `js:"policies"`
	Total    int64                       `js:"total"`
}

func (h *Harbor) ListReplicationPolicies(args ...goja.Value) ListReplicationPoliciesResult {
	h.mustInitialized()

	params := operation.NewListReplicationPoliciesParams()
	if len(args) > 0 {
		rt := h.vu.Runtime()
		if err := rt.ExportTo(args[0], params); err != nil {
			common.Throw(rt, err)
		}
	}

	res, err := h.api.Replication.ListReplicationPolicies(h.vu.Context(), params)
	Checkf(h.vu.Runtime(), err, "failed to list replication policies	")

	return ListReplicationPoliciesResult{
		Policies: res.Payload,
		Total:    res.XTotalCount,
	}
}

func (h *Harbor) StartReplication(policyID int64) int64 {
	h.mustInitialized()

	params := operation.NewStartReplicationParams()
	params.WithExecution(&models.StartReplicationExecution{PolicyID: policyID})

	res, err := h.api.Replication.StartReplication(h.vu.Context(), params)
	Checkf(h.vu.Runtime(), err, "failed to start replication %d", policyID)

	return IDFromLocation(h.vu.Runtime(), res.Location)
}

func (h *Harbor) GetReplicationExecution(executionID int64) *models.ReplicationExecution {
	h.mustInitialized()

	params := operation.NewGetReplicationExecutionParams()
	params.WithID(executionID)

	res, err := h.api.Replication.GetReplicationExecution(h.vu.Context(), params)
	Checkf(h.vu.Runtime(), err, "failed to get replication execution %d", executionID)

	return res.Payload
}
