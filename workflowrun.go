// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package serval

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/ServalHQ/serval-go/internal/apijson"
	"github.com/ServalHQ/serval-go/internal/apiquery"
	"github.com/ServalHQ/serval-go/internal/requestconfig"
	"github.com/ServalHQ/serval-go/option"
	"github.com/ServalHQ/serval-go/packages/param"
	"github.com/ServalHQ/serval-go/packages/respjson"
)

// WorkflowRunService contains methods and other services that help with
// interacting with the serval API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowRunService] method instead.
type WorkflowRunService struct {
	Options []option.RequestOption
}

// NewWorkflowRunService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkflowRunService(opts ...option.RequestOption) (r WorkflowRunService) {
	r = WorkflowRunService{}
	r.Options = opts
	return
}

// Get a specific workflow run by ID.
func (r *WorkflowRunService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *WorkflowRun, err error) {
	var env WorkflowRunGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/workflow-runs/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// List workflow runs for a team. Filter by workflow ID, status, or time range.
func (r *WorkflowRunService) List(ctx context.Context, query WorkflowRunListParams, opts ...option.RequestOption) (res *WorkflowRunListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/workflow-runs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Search workflow runs with filters. Supports filtering by workflow ID, status,
// user IDs, time range, and linked ticket.
func (r *WorkflowRunService) Search(ctx context.Context, body WorkflowRunSearchParams, opts ...option.RequestOption) (res *WorkflowRunSearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/workflow-runs/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// WorkflowRun represents a single execution of a workflow.
type WorkflowRun struct {
	// The unique ID of the workflow run.
	ID string `json:"id"`
	// The timestamp when the workflow run completed (if applicable).
	CompletedAt string `json:"completedAt,nullable"`
	// The timestamp when the workflow run was created.
	CreatedAt string `json:"createdAt"`
	// The ID of the user who initiated the workflow run.
	InitiatedByUserID string `json:"initiatedByUserId"`
	// The inputs provided to the workflow (JSON string).
	Inputs string `json:"inputs,nullable"`
	// The ID of the linked ticket, if any.
	LinkedTicketID string `json:"linkedTicketId,nullable"`
	// The output of the workflow run (JSON string, available when completed or
	// failed).
	Output string `json:"output,nullable"`
	// The status of the workflow run.
	//
	// Any of "WORKFLOW_RUN_STATUS_UNSPECIFIED", "WORKFLOW_RUN_STATUS_PENDING",
	// "WORKFLOW_RUN_STATUS_RUNNING", "WORKFLOW_RUN_STATUS_COMPLETED",
	// "WORKFLOW_RUN_STATUS_FAILED", "WORKFLOW_RUN_STATUS_DENIED",
	// "WORKFLOW_RUN_STATUS_CANCELED".
	Status WorkflowRunStatus `json:"status"`
	// The ID of the target user for the workflow run (if different from initiator).
	TargetUserID string `json:"targetUserId,nullable"`
	// The ID of the team that the workflow belongs to.
	TeamID string `json:"teamId"`
	// The ID of the workflow that was run.
	WorkflowID string `json:"workflowId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CompletedAt       respjson.Field
		CreatedAt         respjson.Field
		InitiatedByUserID respjson.Field
		Inputs            respjson.Field
		LinkedTicketID    respjson.Field
		Output            respjson.Field
		Status            respjson.Field
		TargetUserID      respjson.Field
		TeamID            respjson.Field
		WorkflowID        respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowRun) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the workflow run.
type WorkflowRunStatus string

const (
	WorkflowRunStatusWorkflowRunStatusUnspecified WorkflowRunStatus = "WORKFLOW_RUN_STATUS_UNSPECIFIED"
	WorkflowRunStatusWorkflowRunStatusPending     WorkflowRunStatus = "WORKFLOW_RUN_STATUS_PENDING"
	WorkflowRunStatusWorkflowRunStatusRunning     WorkflowRunStatus = "WORKFLOW_RUN_STATUS_RUNNING"
	WorkflowRunStatusWorkflowRunStatusCompleted   WorkflowRunStatus = "WORKFLOW_RUN_STATUS_COMPLETED"
	WorkflowRunStatusWorkflowRunStatusFailed      WorkflowRunStatus = "WORKFLOW_RUN_STATUS_FAILED"
	WorkflowRunStatusWorkflowRunStatusDenied      WorkflowRunStatus = "WORKFLOW_RUN_STATUS_DENIED"
	WorkflowRunStatusWorkflowRunStatusCanceled    WorkflowRunStatus = "WORKFLOW_RUN_STATUS_CANCELED"
)

type WorkflowRunListResponse struct {
	// The list of workflow runs.
	Data []WorkflowRun `json:"data"`
	// Token for retrieving the next page of results. Empty if no more results.
	NextPageToken string `json:"nextPageToken,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data          respjson.Field
		NextPageToken respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowRunListResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRunListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowRunSearchResponse struct {
	// The list of workflow runs.
	Data []WorkflowRun `json:"data"`
	// Token for retrieving the next page of results. Empty if no more results.
	NextPageToken string `json:"nextPageToken,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data          respjson.Field
		NextPageToken respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowRunSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRunSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowRunGetResponseEnvelope struct {
	// The workflow run.
	Data WorkflowRun `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowRunGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *WorkflowRunGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowRunListParams struct {
	// Maximum number of results to return. Default is 100, maximum is 1000.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Token for pagination. Leave empty for the first request.
	PageToken param.Opt[string] `query:"pageToken,omitzero" json:"-"`
	// The ID of the team. Required.
	TeamID param.Opt[string] `query:"teamId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowRunListParams]'s query parameters as `url.Values`.
func (r WorkflowRunListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkflowRunSearchParams struct {
	// Filter by runs created after this timestamp (RFC3339 format).
	CreatedAfter param.Opt[string] `json:"createdAfter,omitzero"`
	// Filter by runs created before this timestamp (RFC3339 format).
	CreatedBefore param.Opt[string] `json:"createdBefore,omitzero"`
	// Filter by the user who initiated the workflow run.
	InitiatedByUserID param.Opt[string] `json:"initiatedByUserId,omitzero"`
	// Filter by linked ticket ID.
	LinkedTicketID param.Opt[string] `json:"linkedTicketId,omitzero"`
	// Maximum number of results to return. Default is 100, maximum is 1000.
	PageSize param.Opt[int64] `json:"pageSize,omitzero"`
	// Token for pagination. Leave empty for the first request.
	PageToken param.Opt[string] `json:"pageToken,omitzero"`
	// Filter by the target user for the workflow run.
	TargetUserID param.Opt[string] `json:"targetUserId,omitzero"`
	// Filter by workflow ID.
	WorkflowID param.Opt[string] `json:"workflowId,omitzero"`
	// The ID of the team. Required.
	TeamID param.Opt[string] `json:"teamId,omitzero"`
	// Filter by statuses (multiple allowed).
	//
	// Any of "WORKFLOW_RUN_STATUS_UNSPECIFIED", "WORKFLOW_RUN_STATUS_PENDING",
	// "WORKFLOW_RUN_STATUS_RUNNING", "WORKFLOW_RUN_STATUS_COMPLETED",
	// "WORKFLOW_RUN_STATUS_FAILED", "WORKFLOW_RUN_STATUS_DENIED",
	// "WORKFLOW_RUN_STATUS_CANCELED".
	Statuses []string `json:"statuses,omitzero"`
	paramObj
}

func (r WorkflowRunSearchParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowRunSearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowRunSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
