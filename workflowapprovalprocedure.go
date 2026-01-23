// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package serval

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/ServalHQ/serval-go/internal/apijson"
	"github.com/ServalHQ/serval-go/internal/requestconfig"
	"github.com/ServalHQ/serval-go/option"
	"github.com/ServalHQ/serval-go/packages/param"
	"github.com/ServalHQ/serval-go/packages/respjson"
)

// WorkflowApprovalProcedureService contains methods and other services that help
// with interacting with the serval API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowApprovalProcedureService] method instead.
type WorkflowApprovalProcedureService struct {
	Options []option.RequestOption
}

// NewWorkflowApprovalProcedureService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWorkflowApprovalProcedureService(opts ...option.RequestOption) (r WorkflowApprovalProcedureService) {
	r = WorkflowApprovalProcedureService{}
	r.Options = opts
	return
}

// Create a new approval procedure for a workflow.
func (r *WorkflowApprovalProcedureService) New(ctx context.Context, workflowID string, body WorkflowApprovalProcedureNewParams, opts ...option.RequestOption) (res *WorkflowApprovalProcedure, err error) {
	var env WorkflowApprovalProcedureNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if workflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return
	}
	path := fmt.Sprintf("v2/workflows/%s/approval-procedures", workflowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Get a specific approval procedure by ID for a workflow.
func (r *WorkflowApprovalProcedureService) Get(ctx context.Context, id string, query WorkflowApprovalProcedureGetParams, opts ...option.RequestOption) (res *WorkflowApprovalProcedure, err error) {
	var env WorkflowApprovalProcedureGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.WorkflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/workflows/%s/approval-procedures/%s", query.WorkflowID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Update an existing approval procedure for a workflow.
func (r *WorkflowApprovalProcedureService) Update(ctx context.Context, id string, params WorkflowApprovalProcedureUpdateParams, opts ...option.RequestOption) (res *WorkflowApprovalProcedure, err error) {
	var env WorkflowApprovalProcedureUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.WorkflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/workflows/%s/approval-procedures/%s", params.WorkflowID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// List all approval procedures for a workflow.
func (r *WorkflowApprovalProcedureService) List(ctx context.Context, workflowID string, opts ...option.RequestOption) (res *[]WorkflowApprovalProcedure, err error) {
	var env WorkflowApprovalProcedureListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if workflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return
	}
	path := fmt.Sprintf("v2/workflows/%s/approval-procedures", workflowID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Delete an approval procedure for a workflow.
func (r *WorkflowApprovalProcedureService) Delete(ctx context.Context, id string, body WorkflowApprovalProcedureDeleteParams, opts ...option.RequestOption) (res *WorkflowApprovalProcedureDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.WorkflowID == "" {
		err = errors.New("missing required workflow_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/workflows/%s/approval-procedures/%s", body.WorkflowID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type WorkflowApprovalProcedure struct {
	// The ID of the workflow approval procedure.
	ID string `json:"id"`
	// The steps in the approval procedure.
	Steps []WorkflowApprovalProcedureStep `json:"steps"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Steps       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowApprovalProcedure) RawJSON() string { return r.JSON.raw }
func (r *WorkflowApprovalProcedure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowApprovalProcedureStep struct {
	// The ID of the approval step.
	ID string `json:"id"`
	// Whether the step can be approved by the requester themselves.
	AllowSelfApproval bool `json:"allowSelfApproval"`
	// A workflow ID to execute to determine the approvers for this step (or to
	// auto-approve the step).
	CustomWorkflowID string `json:"customWorkflowId"`
	// The IDs of the Serval groups that can approve the step.
	ServalGroupIDs []string `json:"servalGroupIds"`
	// The IDs of the specific users that can approve the step.
	SpecificUserIDs []string `json:"specificUserIds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AllowSelfApproval respjson.Field
		CustomWorkflowID  respjson.Field
		ServalGroupIDs    respjson.Field
		SpecificUserIDs   respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowApprovalProcedureStep) RawJSON() string { return r.JSON.raw }
func (r *WorkflowApprovalProcedureStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowApprovalProcedureDeleteResponse = any

type WorkflowApprovalProcedureNewParams struct {
	// The approval steps for the procedure.
	Steps []WorkflowApprovalProcedureNewParamsStep `json:"steps,omitzero"`
	paramObj
}

func (r WorkflowApprovalProcedureNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowApprovalProcedureNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowApprovalProcedureNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowApprovalProcedureNewParamsStep struct {
	// Whether the step can be approved by the requester themselves.
	AllowSelfApproval param.Opt[bool] `json:"allowSelfApproval,omitzero"`
	// A workflow ID to execute to determine the approvers for this step (or to
	// auto-approve the step).
	CustomWorkflowID param.Opt[string] `json:"customWorkflowId,omitzero"`
	// The IDs of the Serval groups that can approve the step.
	ServalGroupIDs []string `json:"servalGroupIds,omitzero"`
	// The IDs of the specific users that can approve the step.
	SpecificUserIDs []string `json:"specificUserIds,omitzero"`
	paramObj
}

func (r WorkflowApprovalProcedureNewParamsStep) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowApprovalProcedureNewParamsStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowApprovalProcedureNewParamsStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowApprovalProcedureNewResponseEnvelope struct {
	// The created approval procedure.
	Data WorkflowApprovalProcedure `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowApprovalProcedureNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *WorkflowApprovalProcedureNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowApprovalProcedureGetParams struct {
	// The ID of the workflow.
	WorkflowID string `path:"workflow_id,required" json:"-"`
	paramObj
}

type WorkflowApprovalProcedureGetResponseEnvelope struct {
	// The approval procedure.
	Data WorkflowApprovalProcedure `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowApprovalProcedureGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *WorkflowApprovalProcedureGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowApprovalProcedureUpdateParams struct {
	// The ID of the workflow.
	WorkflowID string `path:"workflow_id,required" json:"-"`
	// The approval steps for the procedure.
	Steps []WorkflowApprovalProcedureUpdateParamsStep `json:"steps,omitzero"`
	paramObj
}

func (r WorkflowApprovalProcedureUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowApprovalProcedureUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowApprovalProcedureUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowApprovalProcedureUpdateParamsStep struct {
	// Whether the step can be approved by the requester themselves.
	AllowSelfApproval param.Opt[bool] `json:"allowSelfApproval,omitzero"`
	// A workflow ID to execute to determine the approvers for this step (or to
	// auto-approve the step).
	CustomWorkflowID param.Opt[string] `json:"customWorkflowId,omitzero"`
	// The IDs of the Serval groups that can approve the step.
	ServalGroupIDs []string `json:"servalGroupIds,omitzero"`
	// The IDs of the specific users that can approve the step.
	SpecificUserIDs []string `json:"specificUserIds,omitzero"`
	paramObj
}

func (r WorkflowApprovalProcedureUpdateParamsStep) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowApprovalProcedureUpdateParamsStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowApprovalProcedureUpdateParamsStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowApprovalProcedureUpdateResponseEnvelope struct {
	// The updated approval procedure.
	Data WorkflowApprovalProcedure `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowApprovalProcedureUpdateResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *WorkflowApprovalProcedureUpdateResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowApprovalProcedureListResponseEnvelope struct {
	// The list of approval procedures (typically 0 or 1).
	Data []WorkflowApprovalProcedure `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowApprovalProcedureListResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *WorkflowApprovalProcedureListResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowApprovalProcedureDeleteParams struct {
	// The ID of the workflow.
	WorkflowID string `path:"workflow_id,required" json:"-"`
	paramObj
}
