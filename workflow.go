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

// WorkflowService contains methods and other services that help with interacting
// with the serval API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkflowService] method instead.
type WorkflowService struct {
	Options            []option.RequestOption
	ApprovalProcedures WorkflowApprovalProcedureService
}

// NewWorkflowService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkflowService(opts ...option.RequestOption) (r WorkflowService) {
	r = WorkflowService{}
	r.Options = opts
	r.ApprovalProcedures = NewWorkflowApprovalProcedureService(opts...)
	return
}

// Create a new workflow for a team.
func (r *WorkflowService) New(ctx context.Context, body WorkflowNewParams, opts ...option.RequestOption) (res *Workflow, err error) {
	var env WorkflowNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "v2/workflows"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Get a specific workflow by ID.
func (r *WorkflowService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Workflow, err error) {
	var env WorkflowGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/workflows/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Update an existing workflow.
func (r *WorkflowService) Update(ctx context.Context, id string, body WorkflowUpdateParams, opts ...option.RequestOption) (res *Workflow, err error) {
	var env WorkflowUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/workflows/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// List all workflows for a team.
func (r *WorkflowService) List(ctx context.Context, query WorkflowListParams, opts ...option.RequestOption) (res *WorkflowListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/workflows"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a workflow.
func (r *WorkflowService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *WorkflowDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/workflows/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Workflow struct {
	// The ID of the workflow.
	ID string `json:"id"`
	// The content/code of the workflow.
	Content string `json:"content"`
	// A description of the workflow.
	Description string `json:"description"`
	// The execution scope of the workflow.
	//
	// Any of "WORKFLOW_EXECUTION_SCOPE_UNSPECIFIED", "TEAM_PRIVATE", "TEAM_PUBLIC".
	ExecutionScope WorkflowExecutionScope `json:"executionScope"`
	// Whether there are unpublished changes to the workflow (computed by server).
	HasUnpublishedChanges bool `json:"hasUnpublishedChanges"`
	// Whether the workflow is published. Set to true to publish the workflow.
	IsPublished bool `json:"isPublished"`
	// Whether the workflow is temporary.
	IsTemporary bool `json:"isTemporary"`
	// The name of the workflow.
	Name string `json:"name"`
	// The parameters schema of the workflow (JSON).
	Parameters string `json:"parameters"`
	// Whether the workflow requires form confirmation.
	RequireFormConfirmation bool `json:"requireFormConfirmation"`
	// IDs of tags associated with this workflow.
	TagIDs []string `json:"tagIds"`
	// The ID of the team that the workflow belongs to.
	TeamID string `json:"teamId"`
	// The type of the workflow.
	//
	// Any of "WORKFLOW_TYPE_UNSPECIFIED", "EXECUTABLE", "GUIDANCE".
	Type WorkflowType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		Content                 respjson.Field
		Description             respjson.Field
		ExecutionScope          respjson.Field
		HasUnpublishedChanges   respjson.Field
		IsPublished             respjson.Field
		IsTemporary             respjson.Field
		Name                    respjson.Field
		Parameters              respjson.Field
		RequireFormConfirmation respjson.Field
		TagIDs                  respjson.Field
		TeamID                  respjson.Field
		Type                    respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Workflow) RawJSON() string { return r.JSON.raw }
func (r *Workflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The execution scope of the workflow.
type WorkflowExecutionScope string

const (
	WorkflowExecutionScopeWorkflowExecutionScopeUnspecified WorkflowExecutionScope = "WORKFLOW_EXECUTION_SCOPE_UNSPECIFIED"
	WorkflowExecutionScopeTeamPrivate                       WorkflowExecutionScope = "TEAM_PRIVATE"
	WorkflowExecutionScopeTeamPublic                        WorkflowExecutionScope = "TEAM_PUBLIC"
)

// The type of the workflow.
type WorkflowType string

const (
	WorkflowTypeWorkflowTypeUnspecified WorkflowType = "WORKFLOW_TYPE_UNSPECIFIED"
	WorkflowTypeExecutable              WorkflowType = "EXECUTABLE"
	WorkflowTypeGuidance                WorkflowType = "GUIDANCE"
)

type WorkflowListResponse struct {
	// The list of workflows.
	Data []Workflow `json:"data"`
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
func (r WorkflowListResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkflowListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowDeleteResponse = any

type WorkflowNewParams struct {
	// The content/code of the workflow.
	Content string `json:"content,required"`
	// The name of the workflow.
	Name string `json:"name,required"`
	// The ID of the team.
	TeamID string `json:"teamId,required"`
	// The type of the workflow.
	//
	// Any of "WORKFLOW_TYPE_UNSPECIFIED", "EXECUTABLE", "GUIDANCE".
	Type WorkflowNewParamsType `json:"type,omitzero,required"`
	// Whether to publish the workflow after creation (optional).
	IsPublished param.Opt[bool] `json:"isPublished,omitzero"`
	// Whether the workflow is temporary (optional).
	IsTemporary param.Opt[bool] `json:"isTemporary,omitzero"`
	// The parameters schema of the workflow (JSON, optional).
	Parameters param.Opt[string] `json:"parameters,omitzero"`
	// Whether the workflow requires form confirmation (optional).
	RequireFormConfirmation param.Opt[bool] `json:"requireFormConfirmation,omitzero"`
	// A description of the workflow.
	Description param.Opt[string] `json:"description,omitzero"`
	// The execution scope of the workflow.
	//
	// Any of "WORKFLOW_EXECUTION_SCOPE_UNSPECIFIED", "TEAM_PRIVATE", "TEAM_PUBLIC".
	ExecutionScope WorkflowNewParamsExecutionScope `json:"executionScope,omitzero"`
	paramObj
}

func (r WorkflowNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the workflow.
type WorkflowNewParamsType string

const (
	WorkflowNewParamsTypeWorkflowTypeUnspecified WorkflowNewParamsType = "WORKFLOW_TYPE_UNSPECIFIED"
	WorkflowNewParamsTypeExecutable              WorkflowNewParamsType = "EXECUTABLE"
	WorkflowNewParamsTypeGuidance                WorkflowNewParamsType = "GUIDANCE"
)

// The execution scope of the workflow.
type WorkflowNewParamsExecutionScope string

const (
	WorkflowNewParamsExecutionScopeWorkflowExecutionScopeUnspecified WorkflowNewParamsExecutionScope = "WORKFLOW_EXECUTION_SCOPE_UNSPECIFIED"
	WorkflowNewParamsExecutionScopeTeamPrivate                       WorkflowNewParamsExecutionScope = "TEAM_PRIVATE"
	WorkflowNewParamsExecutionScopeTeamPublic                        WorkflowNewParamsExecutionScope = "TEAM_PUBLIC"
)

type WorkflowNewResponseEnvelope struct {
	// The created workflow.
	Data Workflow `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *WorkflowNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowGetResponseEnvelope struct {
	// The workflow.
	Data Workflow `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *WorkflowGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowUpdateParams struct {
	// Whether the workflow is published. Set to true to publish the workflow.
	IsPublished param.Opt[bool] `json:"isPublished,omitzero"`
	// The content/code of the workflow.
	Content param.Opt[string] `json:"content,omitzero"`
	// A description of the workflow.
	Description param.Opt[string] `json:"description,omitzero"`
	// Whether the workflow is temporary.
	IsTemporary param.Opt[bool] `json:"isTemporary,omitzero"`
	// The name of the workflow.
	Name param.Opt[string] `json:"name,omitzero"`
	// The parameters schema of the workflow (JSON).
	Parameters param.Opt[string] `json:"parameters,omitzero"`
	// Whether the workflow requires form confirmation.
	RequireFormConfirmation param.Opt[bool] `json:"requireFormConfirmation,omitzero"`
	// The execution scope of the workflow.
	//
	// Any of "WORKFLOW_EXECUTION_SCOPE_UNSPECIFIED", "TEAM_PRIVATE", "TEAM_PUBLIC".
	ExecutionScope WorkflowUpdateParamsExecutionScope `json:"executionScope,omitzero"`
	// The type of the workflow.
	//
	// Any of "WORKFLOW_TYPE_UNSPECIFIED", "EXECUTABLE", "GUIDANCE".
	Type WorkflowUpdateParamsType `json:"type,omitzero"`
	paramObj
}

func (r WorkflowUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The execution scope of the workflow.
type WorkflowUpdateParamsExecutionScope string

const (
	WorkflowUpdateParamsExecutionScopeWorkflowExecutionScopeUnspecified WorkflowUpdateParamsExecutionScope = "WORKFLOW_EXECUTION_SCOPE_UNSPECIFIED"
	WorkflowUpdateParamsExecutionScopeTeamPrivate                       WorkflowUpdateParamsExecutionScope = "TEAM_PRIVATE"
	WorkflowUpdateParamsExecutionScopeTeamPublic                        WorkflowUpdateParamsExecutionScope = "TEAM_PUBLIC"
)

// The type of the workflow.
type WorkflowUpdateParamsType string

const (
	WorkflowUpdateParamsTypeWorkflowTypeUnspecified WorkflowUpdateParamsType = "WORKFLOW_TYPE_UNSPECIFIED"
	WorkflowUpdateParamsTypeExecutable              WorkflowUpdateParamsType = "EXECUTABLE"
	WorkflowUpdateParamsTypeGuidance                WorkflowUpdateParamsType = "GUIDANCE"
)

type WorkflowUpdateResponseEnvelope struct {
	// The updated workflow.
	Data Workflow `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowUpdateResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *WorkflowUpdateResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowListParams struct {
	// Whether to include temporary workflows (optional, defaults to false).
	IncludeTemporary param.Opt[bool] `query:"includeTemporary,omitzero" json:"-"`
	// Maximum number of results to return. Default is 1000, maximum is 1000.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Token for pagination. Leave empty for the first request.
	PageToken param.Opt[string] `query:"pageToken,omitzero" json:"-"`
	// The ID of the team.
	TeamID param.Opt[string] `query:"teamId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkflowListParams]'s query parameters as `url.Values`.
func (r WorkflowListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
