// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package serval

import (
	"context"
	"encoding/json"
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
	// The ID of the workflow this approval procedure belongs to.
	WorkflowID string `json:"workflowId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Steps       respjson.Field
		WorkflowID  respjson.Field
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
	// Whether the step can be approved by the requester themselves. optional so server
	// can distinguish "not set" from "explicitly false" (DB defaults to TRUE; proto3
	// defaults bool to false)
	AllowSelfApproval bool `json:"allowSelfApproval" api:"nullable"`
	// Exactly one of approvers or custom_workflow must be set. Mutual exclusivity
	// validated server-side.
	Approvers []WorkflowApprovalProcedureStepApproverUnion `json:"approvers"`
	// Configuration for a custom workflow that determines approvers or auto-approves.
	CustomWorkflow WorkflowApprovalProcedureStepCustomWorkflow `json:"customWorkflow" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AllowSelfApproval respjson.Field
		Approvers         respjson.Field
		CustomWorkflow    respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowApprovalProcedureStep) RawJSON() string { return r.JSON.raw }
func (r *WorkflowApprovalProcedureStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WorkflowApprovalProcedureStepApproverUnion contains all possible properties and
// values from [WorkflowApprovalProcedureStepApproverAppOwner],
// [WorkflowApprovalProcedureStepApproverGroup],
// [WorkflowApprovalProcedureStepApproverManager],
// [WorkflowApprovalProcedureStepApproverUser].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WorkflowApprovalProcedureStepApproverUnion struct {
	// This field is from variant [WorkflowApprovalProcedureStepApproverAppOwner].
	AppOwner any  `json:"appOwner"`
	Notify   bool `json:"notify"`
	// This field is from variant [WorkflowApprovalProcedureStepApproverGroup].
	Group WorkflowApprovalProcedureStepApproverGroupGroup `json:"group"`
	// This field is from variant [WorkflowApprovalProcedureStepApproverManager].
	Manager any `json:"manager"`
	// This field is from variant [WorkflowApprovalProcedureStepApproverUser].
	User WorkflowApprovalProcedureStepApproverUserUser `json:"user"`
	JSON struct {
		AppOwner respjson.Field
		Notify   respjson.Field
		Group    respjson.Field
		Manager  respjson.Field
		User     respjson.Field
		raw      string
	} `json:"-"`
}

func (u WorkflowApprovalProcedureStepApproverUnion) AsAppOwner() (v WorkflowApprovalProcedureStepApproverAppOwner) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowApprovalProcedureStepApproverUnion) AsGroup() (v WorkflowApprovalProcedureStepApproverGroup) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowApprovalProcedureStepApproverUnion) AsManager() (v WorkflowApprovalProcedureStepApproverManager) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowApprovalProcedureStepApproverUnion) AsUser() (v WorkflowApprovalProcedureStepApproverUser) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WorkflowApprovalProcedureStepApproverUnion) RawJSON() string { return u.JSON.raw }

func (r *WorkflowApprovalProcedureStepApproverUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// App owners as approvers. Only valid for access policy approval procedures.
type WorkflowApprovalProcedureStepApproverAppOwner struct {
	// App owners as approvers. Only valid for access policy approval procedures.
	AppOwner any `json:"appOwner" api:"required"`
	// Whether to notify this approver when the step is pending.
	Notify bool `json:"notify"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AppOwner    respjson.Field
		Notify      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowApprovalProcedureStepApproverAppOwner) RawJSON() string { return r.JSON.raw }
func (r *WorkflowApprovalProcedureStepApproverAppOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Serval group as approvers.
type WorkflowApprovalProcedureStepApproverGroup struct {
	// A Serval group as approvers.
	Group WorkflowApprovalProcedureStepApproverGroupGroup `json:"group" api:"required"`
	// Whether to notify this approver when the step is pending.
	Notify bool `json:"notify"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Group       respjson.Field
		Notify      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowApprovalProcedureStepApproverGroup) RawJSON() string { return r.JSON.raw }
func (r *WorkflowApprovalProcedureStepApproverGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Serval group as approvers.
type WorkflowApprovalProcedureStepApproverGroupGroup struct {
	// The ID of the Serval group.
	GroupID string `json:"groupId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		GroupID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowApprovalProcedureStepApproverGroupGroup) RawJSON() string { return r.JSON.raw }
func (r *WorkflowApprovalProcedureStepApproverGroupGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The requester's manager as an approver.
type WorkflowApprovalProcedureStepApproverManager struct {
	// The requester's manager as an approver.
	Manager any `json:"manager" api:"required"`
	// Whether to notify this approver when the step is pending.
	Notify bool `json:"notify"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Manager     respjson.Field
		Notify      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowApprovalProcedureStepApproverManager) RawJSON() string { return r.JSON.raw }
func (r *WorkflowApprovalProcedureStepApproverManager) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A specific user as an approver.
type WorkflowApprovalProcedureStepApproverUser struct {
	// A specific user as an approver.
	User WorkflowApprovalProcedureStepApproverUserUser `json:"user" api:"required"`
	// Whether to notify this approver when the step is pending.
	Notify bool `json:"notify"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		User        respjson.Field
		Notify      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowApprovalProcedureStepApproverUser) RawJSON() string { return r.JSON.raw }
func (r *WorkflowApprovalProcedureStepApproverUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A specific user as an approver.
type WorkflowApprovalProcedureStepApproverUserUser struct {
	// The ID of the user.
	UserID string `json:"userId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowApprovalProcedureStepApproverUserUser) RawJSON() string { return r.JSON.raw }
func (r *WorkflowApprovalProcedureStepApproverUserUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for a custom workflow that determines approvers or auto-approves.
type WorkflowApprovalProcedureStepCustomWorkflow struct {
	// The ID of the workflow to execute.
	WorkflowID string `json:"workflowId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		WorkflowID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowApprovalProcedureStepCustomWorkflow) RawJSON() string { return r.JSON.raw }
func (r *WorkflowApprovalProcedureStepCustomWorkflow) UnmarshalJSON(data []byte) error {
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
	// Whether the step can be approved by the requester themselves. optional so server
	// can distinguish "not set" from "explicitly false" (DB defaults to TRUE; proto3
	// defaults bool to false)
	AllowSelfApproval param.Opt[bool] `json:"allowSelfApproval,omitzero"`
	// Configuration for a custom workflow that determines approvers or auto-approves.
	CustomWorkflow WorkflowApprovalProcedureNewParamsStepCustomWorkflow `json:"customWorkflow,omitzero"`
	// Exactly one of approvers or custom_workflow must be set. Mutual exclusivity
	// validated server-side.
	Approvers []WorkflowApprovalProcedureNewParamsStepApproverUnion `json:"approvers,omitzero"`
	paramObj
}

func (r WorkflowApprovalProcedureNewParamsStep) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowApprovalProcedureNewParamsStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowApprovalProcedureNewParamsStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WorkflowApprovalProcedureNewParamsStepApproverUnion struct {
	OfAppOwner *WorkflowApprovalProcedureNewParamsStepApproverAppOwner `json:",omitzero,inline"`
	OfGroup    *WorkflowApprovalProcedureNewParamsStepApproverGroup    `json:",omitzero,inline"`
	OfManager  *WorkflowApprovalProcedureNewParamsStepApproverManager  `json:",omitzero,inline"`
	OfUser     *WorkflowApprovalProcedureNewParamsStepApproverUser     `json:",omitzero,inline"`
	paramUnion
}

func (u WorkflowApprovalProcedureNewParamsStepApproverUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAppOwner, u.OfGroup, u.OfManager, u.OfUser)
}
func (u *WorkflowApprovalProcedureNewParamsStepApproverUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WorkflowApprovalProcedureNewParamsStepApproverUnion) asAny() any {
	if !param.IsOmitted(u.OfAppOwner) {
		return u.OfAppOwner
	} else if !param.IsOmitted(u.OfGroup) {
		return u.OfGroup
	} else if !param.IsOmitted(u.OfManager) {
		return u.OfManager
	} else if !param.IsOmitted(u.OfUser) {
		return u.OfUser
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowApprovalProcedureNewParamsStepApproverUnion) GetAppOwner() *any {
	if vt := u.OfAppOwner; vt != nil {
		return &vt.AppOwner
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowApprovalProcedureNewParamsStepApproverUnion) GetGroup() *WorkflowApprovalProcedureNewParamsStepApproverGroupGroup {
	if vt := u.OfGroup; vt != nil {
		return &vt.Group
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowApprovalProcedureNewParamsStepApproverUnion) GetManager() *any {
	if vt := u.OfManager; vt != nil {
		return &vt.Manager
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowApprovalProcedureNewParamsStepApproverUnion) GetUser() *WorkflowApprovalProcedureNewParamsStepApproverUserUser {
	if vt := u.OfUser; vt != nil {
		return &vt.User
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowApprovalProcedureNewParamsStepApproverUnion) GetNotify() *bool {
	if vt := u.OfAppOwner; vt != nil && vt.Notify.Valid() {
		return &vt.Notify.Value
	} else if vt := u.OfGroup; vt != nil && vt.Notify.Valid() {
		return &vt.Notify.Value
	} else if vt := u.OfManager; vt != nil && vt.Notify.Valid() {
		return &vt.Notify.Value
	} else if vt := u.OfUser; vt != nil && vt.Notify.Valid() {
		return &vt.Notify.Value
	}
	return nil
}

// App owners as approvers. Only valid for access policy approval procedures.
//
// The property AppOwner is required.
type WorkflowApprovalProcedureNewParamsStepApproverAppOwner struct {
	// App owners as approvers. Only valid for access policy approval procedures.
	AppOwner any `json:"appOwner,omitzero" api:"required"`
	// Whether to notify this approver when the step is pending.
	Notify param.Opt[bool] `json:"notify,omitzero"`
	paramObj
}

func (r WorkflowApprovalProcedureNewParamsStepApproverAppOwner) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowApprovalProcedureNewParamsStepApproverAppOwner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowApprovalProcedureNewParamsStepApproverAppOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Serval group as approvers.
//
// The property Group is required.
type WorkflowApprovalProcedureNewParamsStepApproverGroup struct {
	// A Serval group as approvers.
	Group WorkflowApprovalProcedureNewParamsStepApproverGroupGroup `json:"group,omitzero" api:"required"`
	// Whether to notify this approver when the step is pending.
	Notify param.Opt[bool] `json:"notify,omitzero"`
	paramObj
}

func (r WorkflowApprovalProcedureNewParamsStepApproverGroup) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowApprovalProcedureNewParamsStepApproverGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowApprovalProcedureNewParamsStepApproverGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Serval group as approvers.
type WorkflowApprovalProcedureNewParamsStepApproverGroupGroup struct {
	// The ID of the Serval group.
	GroupID param.Opt[string] `json:"groupId,omitzero"`
	paramObj
}

func (r WorkflowApprovalProcedureNewParamsStepApproverGroupGroup) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowApprovalProcedureNewParamsStepApproverGroupGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowApprovalProcedureNewParamsStepApproverGroupGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The requester's manager as an approver.
//
// The property Manager is required.
type WorkflowApprovalProcedureNewParamsStepApproverManager struct {
	// The requester's manager as an approver.
	Manager any `json:"manager,omitzero" api:"required"`
	// Whether to notify this approver when the step is pending.
	Notify param.Opt[bool] `json:"notify,omitzero"`
	paramObj
}

func (r WorkflowApprovalProcedureNewParamsStepApproverManager) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowApprovalProcedureNewParamsStepApproverManager
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowApprovalProcedureNewParamsStepApproverManager) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A specific user as an approver.
//
// The property User is required.
type WorkflowApprovalProcedureNewParamsStepApproverUser struct {
	// A specific user as an approver.
	User WorkflowApprovalProcedureNewParamsStepApproverUserUser `json:"user,omitzero" api:"required"`
	// Whether to notify this approver when the step is pending.
	Notify param.Opt[bool] `json:"notify,omitzero"`
	paramObj
}

func (r WorkflowApprovalProcedureNewParamsStepApproverUser) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowApprovalProcedureNewParamsStepApproverUser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowApprovalProcedureNewParamsStepApproverUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A specific user as an approver.
type WorkflowApprovalProcedureNewParamsStepApproverUserUser struct {
	// The ID of the user.
	UserID param.Opt[string] `json:"userId,omitzero"`
	paramObj
}

func (r WorkflowApprovalProcedureNewParamsStepApproverUserUser) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowApprovalProcedureNewParamsStepApproverUserUser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowApprovalProcedureNewParamsStepApproverUserUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for a custom workflow that determines approvers or auto-approves.
type WorkflowApprovalProcedureNewParamsStepCustomWorkflow struct {
	// The ID of the workflow to execute.
	WorkflowID param.Opt[string] `json:"workflowId,omitzero"`
	paramObj
}

func (r WorkflowApprovalProcedureNewParamsStepCustomWorkflow) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowApprovalProcedureNewParamsStepCustomWorkflow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowApprovalProcedureNewParamsStepCustomWorkflow) UnmarshalJSON(data []byte) error {
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
	WorkflowID string `path:"workflow_id" api:"required" json:"-"`
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
	WorkflowID string `path:"workflow_id" api:"required" json:"-"`
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
	// Whether the step can be approved by the requester themselves. optional so server
	// can distinguish "not set" from "explicitly false" (DB defaults to TRUE; proto3
	// defaults bool to false)
	AllowSelfApproval param.Opt[bool] `json:"allowSelfApproval,omitzero"`
	// Configuration for a custom workflow that determines approvers or auto-approves.
	CustomWorkflow WorkflowApprovalProcedureUpdateParamsStepCustomWorkflow `json:"customWorkflow,omitzero"`
	// Exactly one of approvers or custom_workflow must be set. Mutual exclusivity
	// validated server-side.
	Approvers []WorkflowApprovalProcedureUpdateParamsStepApproverUnion `json:"approvers,omitzero"`
	paramObj
}

func (r WorkflowApprovalProcedureUpdateParamsStep) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowApprovalProcedureUpdateParamsStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowApprovalProcedureUpdateParamsStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WorkflowApprovalProcedureUpdateParamsStepApproverUnion struct {
	OfAppOwner *WorkflowApprovalProcedureUpdateParamsStepApproverAppOwner `json:",omitzero,inline"`
	OfGroup    *WorkflowApprovalProcedureUpdateParamsStepApproverGroup    `json:",omitzero,inline"`
	OfManager  *WorkflowApprovalProcedureUpdateParamsStepApproverManager  `json:",omitzero,inline"`
	OfUser     *WorkflowApprovalProcedureUpdateParamsStepApproverUser     `json:",omitzero,inline"`
	paramUnion
}

func (u WorkflowApprovalProcedureUpdateParamsStepApproverUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAppOwner, u.OfGroup, u.OfManager, u.OfUser)
}
func (u *WorkflowApprovalProcedureUpdateParamsStepApproverUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WorkflowApprovalProcedureUpdateParamsStepApproverUnion) asAny() any {
	if !param.IsOmitted(u.OfAppOwner) {
		return u.OfAppOwner
	} else if !param.IsOmitted(u.OfGroup) {
		return u.OfGroup
	} else if !param.IsOmitted(u.OfManager) {
		return u.OfManager
	} else if !param.IsOmitted(u.OfUser) {
		return u.OfUser
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowApprovalProcedureUpdateParamsStepApproverUnion) GetAppOwner() *any {
	if vt := u.OfAppOwner; vt != nil {
		return &vt.AppOwner
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowApprovalProcedureUpdateParamsStepApproverUnion) GetGroup() *WorkflowApprovalProcedureUpdateParamsStepApproverGroupGroup {
	if vt := u.OfGroup; vt != nil {
		return &vt.Group
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowApprovalProcedureUpdateParamsStepApproverUnion) GetManager() *any {
	if vt := u.OfManager; vt != nil {
		return &vt.Manager
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowApprovalProcedureUpdateParamsStepApproverUnion) GetUser() *WorkflowApprovalProcedureUpdateParamsStepApproverUserUser {
	if vt := u.OfUser; vt != nil {
		return &vt.User
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowApprovalProcedureUpdateParamsStepApproverUnion) GetNotify() *bool {
	if vt := u.OfAppOwner; vt != nil && vt.Notify.Valid() {
		return &vt.Notify.Value
	} else if vt := u.OfGroup; vt != nil && vt.Notify.Valid() {
		return &vt.Notify.Value
	} else if vt := u.OfManager; vt != nil && vt.Notify.Valid() {
		return &vt.Notify.Value
	} else if vt := u.OfUser; vt != nil && vt.Notify.Valid() {
		return &vt.Notify.Value
	}
	return nil
}

// App owners as approvers. Only valid for access policy approval procedures.
//
// The property AppOwner is required.
type WorkflowApprovalProcedureUpdateParamsStepApproverAppOwner struct {
	// App owners as approvers. Only valid for access policy approval procedures.
	AppOwner any `json:"appOwner,omitzero" api:"required"`
	// Whether to notify this approver when the step is pending.
	Notify param.Opt[bool] `json:"notify,omitzero"`
	paramObj
}

func (r WorkflowApprovalProcedureUpdateParamsStepApproverAppOwner) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowApprovalProcedureUpdateParamsStepApproverAppOwner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowApprovalProcedureUpdateParamsStepApproverAppOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Serval group as approvers.
//
// The property Group is required.
type WorkflowApprovalProcedureUpdateParamsStepApproverGroup struct {
	// A Serval group as approvers.
	Group WorkflowApprovalProcedureUpdateParamsStepApproverGroupGroup `json:"group,omitzero" api:"required"`
	// Whether to notify this approver when the step is pending.
	Notify param.Opt[bool] `json:"notify,omitzero"`
	paramObj
}

func (r WorkflowApprovalProcedureUpdateParamsStepApproverGroup) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowApprovalProcedureUpdateParamsStepApproverGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowApprovalProcedureUpdateParamsStepApproverGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Serval group as approvers.
type WorkflowApprovalProcedureUpdateParamsStepApproverGroupGroup struct {
	// The ID of the Serval group.
	GroupID param.Opt[string] `json:"groupId,omitzero"`
	paramObj
}

func (r WorkflowApprovalProcedureUpdateParamsStepApproverGroupGroup) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowApprovalProcedureUpdateParamsStepApproverGroupGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowApprovalProcedureUpdateParamsStepApproverGroupGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The requester's manager as an approver.
//
// The property Manager is required.
type WorkflowApprovalProcedureUpdateParamsStepApproverManager struct {
	// The requester's manager as an approver.
	Manager any `json:"manager,omitzero" api:"required"`
	// Whether to notify this approver when the step is pending.
	Notify param.Opt[bool] `json:"notify,omitzero"`
	paramObj
}

func (r WorkflowApprovalProcedureUpdateParamsStepApproverManager) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowApprovalProcedureUpdateParamsStepApproverManager
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowApprovalProcedureUpdateParamsStepApproverManager) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A specific user as an approver.
//
// The property User is required.
type WorkflowApprovalProcedureUpdateParamsStepApproverUser struct {
	// A specific user as an approver.
	User WorkflowApprovalProcedureUpdateParamsStepApproverUserUser `json:"user,omitzero" api:"required"`
	// Whether to notify this approver when the step is pending.
	Notify param.Opt[bool] `json:"notify,omitzero"`
	paramObj
}

func (r WorkflowApprovalProcedureUpdateParamsStepApproverUser) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowApprovalProcedureUpdateParamsStepApproverUser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowApprovalProcedureUpdateParamsStepApproverUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A specific user as an approver.
type WorkflowApprovalProcedureUpdateParamsStepApproverUserUser struct {
	// The ID of the user.
	UserID param.Opt[string] `json:"userId,omitzero"`
	paramObj
}

func (r WorkflowApprovalProcedureUpdateParamsStepApproverUserUser) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowApprovalProcedureUpdateParamsStepApproverUserUser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowApprovalProcedureUpdateParamsStepApproverUserUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for a custom workflow that determines approvers or auto-approves.
type WorkflowApprovalProcedureUpdateParamsStepCustomWorkflow struct {
	// The ID of the workflow to execute.
	WorkflowID param.Opt[string] `json:"workflowId,omitzero"`
	paramObj
}

func (r WorkflowApprovalProcedureUpdateParamsStepCustomWorkflow) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowApprovalProcedureUpdateParamsStepCustomWorkflow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowApprovalProcedureUpdateParamsStepCustomWorkflow) UnmarshalJSON(data []byte) error {
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
	WorkflowID string `path:"workflow_id" api:"required" json:"-"`
	paramObj
}
