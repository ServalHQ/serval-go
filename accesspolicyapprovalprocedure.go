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

// AccessPolicyApprovalProcedureService contains methods and other services that
// help with interacting with the serval API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessPolicyApprovalProcedureService] method instead.
type AccessPolicyApprovalProcedureService struct {
	Options []option.RequestOption
}

// NewAccessPolicyApprovalProcedureService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccessPolicyApprovalProcedureService(opts ...option.RequestOption) (r AccessPolicyApprovalProcedureService) {
	r = AccessPolicyApprovalProcedureService{}
	r.Options = opts
	return
}

// Create a new approval procedure for an access policy.
func (r *AccessPolicyApprovalProcedureService) New(ctx context.Context, accessPolicyID string, body AccessPolicyApprovalProcedureNewParams, opts ...option.RequestOption) (res *AccessPolicyApprovalProcedure, err error) {
	var env AccessPolicyApprovalProcedureNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if accessPolicyID == "" {
		err = errors.New("missing required access_policy_id parameter")
		return
	}
	path := fmt.Sprintf("v2/access-policies/%s/approval-procedures", accessPolicyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Get a specific approval procedure by ID for an access policy.
func (r *AccessPolicyApprovalProcedureService) Get(ctx context.Context, id string, query AccessPolicyApprovalProcedureGetParams, opts ...option.RequestOption) (res *AccessPolicyApprovalProcedure, err error) {
	var env AccessPolicyApprovalProcedureGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccessPolicyID == "" {
		err = errors.New("missing required access_policy_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/access-policies/%s/approval-procedures/%s", query.AccessPolicyID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Update an existing approval procedure for an access policy.
func (r *AccessPolicyApprovalProcedureService) Update(ctx context.Context, id string, params AccessPolicyApprovalProcedureUpdateParams, opts ...option.RequestOption) (res *AccessPolicyApprovalProcedure, err error) {
	var env AccessPolicyApprovalProcedureUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccessPolicyID == "" {
		err = errors.New("missing required access_policy_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/access-policies/%s/approval-procedures/%s", params.AccessPolicyID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// List all approval procedures for an access policy.
func (r *AccessPolicyApprovalProcedureService) List(ctx context.Context, accessPolicyID string, opts ...option.RequestOption) (res *[]AccessPolicyApprovalProcedure, err error) {
	var env AccessPolicyApprovalProcedureListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if accessPolicyID == "" {
		err = errors.New("missing required access_policy_id parameter")
		return
	}
	path := fmt.Sprintf("v2/access-policies/%s/approval-procedures", accessPolicyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Delete an approval procedure for an access policy.
func (r *AccessPolicyApprovalProcedureService) Delete(ctx context.Context, id string, body AccessPolicyApprovalProcedureDeleteParams, opts ...option.RequestOption) (res *AccessPolicyApprovalProcedureDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccessPolicyID == "" {
		err = errors.New("missing required access_policy_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/access-policies/%s/approval-procedures/%s", body.AccessPolicyID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AccessPolicyApprovalProcedure struct {
	// The ID of the access policy approval procedure.
	ID string `json:"id"`
	// The ID of the access policy this approval procedure belongs to.
	AccessPolicyID string `json:"accessPolicyId"`
	// The steps in the approval procedure.
	Steps []AccessPolicyApprovalProcedureStep `json:"steps"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		AccessPolicyID respjson.Field
		Steps          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessPolicyApprovalProcedure) RawJSON() string { return r.JSON.raw }
func (r *AccessPolicyApprovalProcedure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessPolicyApprovalProcedureStep struct {
	// The ID of the approval step.
	ID string `json:"id"`
	// Whether the step can be approved by the requester themselves. optional so server
	// can distinguish "not set" from "explicitly false" (DB defaults to TRUE; proto3
	// defaults bool to false)
	AllowSelfApproval bool `json:"allowSelfApproval" api:"nullable"`
	// Exactly one of approvers or custom_workflow must be set. Mutual exclusivity
	// validated server-side.
	Approvers []AccessPolicyApprovalProcedureStepApproverUnion `json:"approvers"`
	// Configuration for a custom workflow that determines approvers or auto-approves.
	CustomWorkflow AccessPolicyApprovalProcedureStepCustomWorkflow `json:"customWorkflow" api:"nullable"`
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
func (r AccessPolicyApprovalProcedureStep) RawJSON() string { return r.JSON.raw }
func (r *AccessPolicyApprovalProcedureStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AccessPolicyApprovalProcedureStepApproverUnion contains all possible properties
// and values from [AccessPolicyApprovalProcedureStepApproverAppOwner],
// [AccessPolicyApprovalProcedureStepApproverGroup],
// [AccessPolicyApprovalProcedureStepApproverManager],
// [AccessPolicyApprovalProcedureStepApproverUser].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AccessPolicyApprovalProcedureStepApproverUnion struct {
	// This field is from variant [AccessPolicyApprovalProcedureStepApproverAppOwner].
	AppOwner any  `json:"appOwner"`
	Notify   bool `json:"notify"`
	// This field is from variant [AccessPolicyApprovalProcedureStepApproverGroup].
	Group AccessPolicyApprovalProcedureStepApproverGroupGroup `json:"group"`
	// This field is from variant [AccessPolicyApprovalProcedureStepApproverManager].
	Manager any `json:"manager"`
	// This field is from variant [AccessPolicyApprovalProcedureStepApproverUser].
	User AccessPolicyApprovalProcedureStepApproverUserUser `json:"user"`
	JSON struct {
		AppOwner respjson.Field
		Notify   respjson.Field
		Group    respjson.Field
		Manager  respjson.Field
		User     respjson.Field
		raw      string
	} `json:"-"`
}

func (u AccessPolicyApprovalProcedureStepApproverUnion) AsAppOwner() (v AccessPolicyApprovalProcedureStepApproverAppOwner) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccessPolicyApprovalProcedureStepApproverUnion) AsGroup() (v AccessPolicyApprovalProcedureStepApproverGroup) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccessPolicyApprovalProcedureStepApproverUnion) AsManager() (v AccessPolicyApprovalProcedureStepApproverManager) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AccessPolicyApprovalProcedureStepApproverUnion) AsUser() (v AccessPolicyApprovalProcedureStepApproverUser) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AccessPolicyApprovalProcedureStepApproverUnion) RawJSON() string { return u.JSON.raw }

func (r *AccessPolicyApprovalProcedureStepApproverUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// App owners as approvers. Only valid for access policy approval procedures.
type AccessPolicyApprovalProcedureStepApproverAppOwner struct {
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
func (r AccessPolicyApprovalProcedureStepApproverAppOwner) RawJSON() string { return r.JSON.raw }
func (r *AccessPolicyApprovalProcedureStepApproverAppOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Serval group as approvers.
type AccessPolicyApprovalProcedureStepApproverGroup struct {
	// A Serval group as approvers.
	Group AccessPolicyApprovalProcedureStepApproverGroupGroup `json:"group" api:"required"`
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
func (r AccessPolicyApprovalProcedureStepApproverGroup) RawJSON() string { return r.JSON.raw }
func (r *AccessPolicyApprovalProcedureStepApproverGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Serval group as approvers.
type AccessPolicyApprovalProcedureStepApproverGroupGroup struct {
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
func (r AccessPolicyApprovalProcedureStepApproverGroupGroup) RawJSON() string { return r.JSON.raw }
func (r *AccessPolicyApprovalProcedureStepApproverGroupGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The requester's manager as an approver.
type AccessPolicyApprovalProcedureStepApproverManager struct {
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
func (r AccessPolicyApprovalProcedureStepApproverManager) RawJSON() string { return r.JSON.raw }
func (r *AccessPolicyApprovalProcedureStepApproverManager) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A specific user as an approver.
type AccessPolicyApprovalProcedureStepApproverUser struct {
	// A specific user as an approver.
	User AccessPolicyApprovalProcedureStepApproverUserUser `json:"user" api:"required"`
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
func (r AccessPolicyApprovalProcedureStepApproverUser) RawJSON() string { return r.JSON.raw }
func (r *AccessPolicyApprovalProcedureStepApproverUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A specific user as an approver.
type AccessPolicyApprovalProcedureStepApproverUserUser struct {
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
func (r AccessPolicyApprovalProcedureStepApproverUserUser) RawJSON() string { return r.JSON.raw }
func (r *AccessPolicyApprovalProcedureStepApproverUserUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for a custom workflow that determines approvers or auto-approves.
type AccessPolicyApprovalProcedureStepCustomWorkflow struct {
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
func (r AccessPolicyApprovalProcedureStepCustomWorkflow) RawJSON() string { return r.JSON.raw }
func (r *AccessPolicyApprovalProcedureStepCustomWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessPolicyApprovalProcedureDeleteResponse = any

type AccessPolicyApprovalProcedureNewParams struct {
	// The approval steps for the procedure.
	Steps []AccessPolicyApprovalProcedureNewParamsStep `json:"steps,omitzero"`
	paramObj
}

func (r AccessPolicyApprovalProcedureNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AccessPolicyApprovalProcedureNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccessPolicyApprovalProcedureNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessPolicyApprovalProcedureNewParamsStep struct {
	// Whether the step can be approved by the requester themselves. optional so server
	// can distinguish "not set" from "explicitly false" (DB defaults to TRUE; proto3
	// defaults bool to false)
	AllowSelfApproval param.Opt[bool] `json:"allowSelfApproval,omitzero"`
	// Configuration for a custom workflow that determines approvers or auto-approves.
	CustomWorkflow AccessPolicyApprovalProcedureNewParamsStepCustomWorkflow `json:"customWorkflow,omitzero"`
	// Exactly one of approvers or custom_workflow must be set. Mutual exclusivity
	// validated server-side.
	Approvers []AccessPolicyApprovalProcedureNewParamsStepApproverUnion `json:"approvers,omitzero"`
	paramObj
}

func (r AccessPolicyApprovalProcedureNewParamsStep) MarshalJSON() (data []byte, err error) {
	type shadow AccessPolicyApprovalProcedureNewParamsStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccessPolicyApprovalProcedureNewParamsStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AccessPolicyApprovalProcedureNewParamsStepApproverUnion struct {
	OfAppOwner *AccessPolicyApprovalProcedureNewParamsStepApproverAppOwner `json:",omitzero,inline"`
	OfGroup    *AccessPolicyApprovalProcedureNewParamsStepApproverGroup    `json:",omitzero,inline"`
	OfManager  *AccessPolicyApprovalProcedureNewParamsStepApproverManager  `json:",omitzero,inline"`
	OfUser     *AccessPolicyApprovalProcedureNewParamsStepApproverUser     `json:",omitzero,inline"`
	paramUnion
}

func (u AccessPolicyApprovalProcedureNewParamsStepApproverUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAppOwner, u.OfGroup, u.OfManager, u.OfUser)
}
func (u *AccessPolicyApprovalProcedureNewParamsStepApproverUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AccessPolicyApprovalProcedureNewParamsStepApproverUnion) asAny() any {
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
func (u AccessPolicyApprovalProcedureNewParamsStepApproverUnion) GetAppOwner() *any {
	if vt := u.OfAppOwner; vt != nil {
		return &vt.AppOwner
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AccessPolicyApprovalProcedureNewParamsStepApproverUnion) GetGroup() *AccessPolicyApprovalProcedureNewParamsStepApproverGroupGroup {
	if vt := u.OfGroup; vt != nil {
		return &vt.Group
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AccessPolicyApprovalProcedureNewParamsStepApproverUnion) GetManager() *any {
	if vt := u.OfManager; vt != nil {
		return &vt.Manager
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AccessPolicyApprovalProcedureNewParamsStepApproverUnion) GetUser() *AccessPolicyApprovalProcedureNewParamsStepApproverUserUser {
	if vt := u.OfUser; vt != nil {
		return &vt.User
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AccessPolicyApprovalProcedureNewParamsStepApproverUnion) GetNotify() *bool {
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
type AccessPolicyApprovalProcedureNewParamsStepApproverAppOwner struct {
	// App owners as approvers. Only valid for access policy approval procedures.
	AppOwner any `json:"appOwner,omitzero" api:"required"`
	// Whether to notify this approver when the step is pending.
	Notify param.Opt[bool] `json:"notify,omitzero"`
	paramObj
}

func (r AccessPolicyApprovalProcedureNewParamsStepApproverAppOwner) MarshalJSON() (data []byte, err error) {
	type shadow AccessPolicyApprovalProcedureNewParamsStepApproverAppOwner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccessPolicyApprovalProcedureNewParamsStepApproverAppOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Serval group as approvers.
//
// The property Group is required.
type AccessPolicyApprovalProcedureNewParamsStepApproverGroup struct {
	// A Serval group as approvers.
	Group AccessPolicyApprovalProcedureNewParamsStepApproverGroupGroup `json:"group,omitzero" api:"required"`
	// Whether to notify this approver when the step is pending.
	Notify param.Opt[bool] `json:"notify,omitzero"`
	paramObj
}

func (r AccessPolicyApprovalProcedureNewParamsStepApproverGroup) MarshalJSON() (data []byte, err error) {
	type shadow AccessPolicyApprovalProcedureNewParamsStepApproverGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccessPolicyApprovalProcedureNewParamsStepApproverGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Serval group as approvers.
type AccessPolicyApprovalProcedureNewParamsStepApproverGroupGroup struct {
	// The ID of the Serval group.
	GroupID param.Opt[string] `json:"groupId,omitzero"`
	paramObj
}

func (r AccessPolicyApprovalProcedureNewParamsStepApproverGroupGroup) MarshalJSON() (data []byte, err error) {
	type shadow AccessPolicyApprovalProcedureNewParamsStepApproverGroupGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccessPolicyApprovalProcedureNewParamsStepApproverGroupGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The requester's manager as an approver.
//
// The property Manager is required.
type AccessPolicyApprovalProcedureNewParamsStepApproverManager struct {
	// The requester's manager as an approver.
	Manager any `json:"manager,omitzero" api:"required"`
	// Whether to notify this approver when the step is pending.
	Notify param.Opt[bool] `json:"notify,omitzero"`
	paramObj
}

func (r AccessPolicyApprovalProcedureNewParamsStepApproverManager) MarshalJSON() (data []byte, err error) {
	type shadow AccessPolicyApprovalProcedureNewParamsStepApproverManager
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccessPolicyApprovalProcedureNewParamsStepApproverManager) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A specific user as an approver.
//
// The property User is required.
type AccessPolicyApprovalProcedureNewParamsStepApproverUser struct {
	// A specific user as an approver.
	User AccessPolicyApprovalProcedureNewParamsStepApproverUserUser `json:"user,omitzero" api:"required"`
	// Whether to notify this approver when the step is pending.
	Notify param.Opt[bool] `json:"notify,omitzero"`
	paramObj
}

func (r AccessPolicyApprovalProcedureNewParamsStepApproverUser) MarshalJSON() (data []byte, err error) {
	type shadow AccessPolicyApprovalProcedureNewParamsStepApproverUser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccessPolicyApprovalProcedureNewParamsStepApproverUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A specific user as an approver.
type AccessPolicyApprovalProcedureNewParamsStepApproverUserUser struct {
	// The ID of the user.
	UserID param.Opt[string] `json:"userId,omitzero"`
	paramObj
}

func (r AccessPolicyApprovalProcedureNewParamsStepApproverUserUser) MarshalJSON() (data []byte, err error) {
	type shadow AccessPolicyApprovalProcedureNewParamsStepApproverUserUser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccessPolicyApprovalProcedureNewParamsStepApproverUserUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for a custom workflow that determines approvers or auto-approves.
type AccessPolicyApprovalProcedureNewParamsStepCustomWorkflow struct {
	// The ID of the workflow to execute.
	WorkflowID param.Opt[string] `json:"workflowId,omitzero"`
	paramObj
}

func (r AccessPolicyApprovalProcedureNewParamsStepCustomWorkflow) MarshalJSON() (data []byte, err error) {
	type shadow AccessPolicyApprovalProcedureNewParamsStepCustomWorkflow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccessPolicyApprovalProcedureNewParamsStepCustomWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessPolicyApprovalProcedureNewResponseEnvelope struct {
	// The created approval procedure.
	Data AccessPolicyApprovalProcedure `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessPolicyApprovalProcedureNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AccessPolicyApprovalProcedureNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessPolicyApprovalProcedureGetParams struct {
	// The ID of the access policy.
	AccessPolicyID string `path:"access_policy_id" api:"required" json:"-"`
	paramObj
}

type AccessPolicyApprovalProcedureGetResponseEnvelope struct {
	// The approval procedure.
	Data AccessPolicyApprovalProcedure `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessPolicyApprovalProcedureGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AccessPolicyApprovalProcedureGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessPolicyApprovalProcedureUpdateParams struct {
	// The ID of the access policy.
	AccessPolicyID string `path:"access_policy_id" api:"required" json:"-"`
	// The approval steps for the procedure.
	Steps []AccessPolicyApprovalProcedureUpdateParamsStep `json:"steps,omitzero"`
	paramObj
}

func (r AccessPolicyApprovalProcedureUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AccessPolicyApprovalProcedureUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccessPolicyApprovalProcedureUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessPolicyApprovalProcedureUpdateParamsStep struct {
	// Whether the step can be approved by the requester themselves. optional so server
	// can distinguish "not set" from "explicitly false" (DB defaults to TRUE; proto3
	// defaults bool to false)
	AllowSelfApproval param.Opt[bool] `json:"allowSelfApproval,omitzero"`
	// Configuration for a custom workflow that determines approvers or auto-approves.
	CustomWorkflow AccessPolicyApprovalProcedureUpdateParamsStepCustomWorkflow `json:"customWorkflow,omitzero"`
	// Exactly one of approvers or custom_workflow must be set. Mutual exclusivity
	// validated server-side.
	Approvers []AccessPolicyApprovalProcedureUpdateParamsStepApproverUnion `json:"approvers,omitzero"`
	paramObj
}

func (r AccessPolicyApprovalProcedureUpdateParamsStep) MarshalJSON() (data []byte, err error) {
	type shadow AccessPolicyApprovalProcedureUpdateParamsStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccessPolicyApprovalProcedureUpdateParamsStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AccessPolicyApprovalProcedureUpdateParamsStepApproverUnion struct {
	OfAppOwner *AccessPolicyApprovalProcedureUpdateParamsStepApproverAppOwner `json:",omitzero,inline"`
	OfGroup    *AccessPolicyApprovalProcedureUpdateParamsStepApproverGroup    `json:",omitzero,inline"`
	OfManager  *AccessPolicyApprovalProcedureUpdateParamsStepApproverManager  `json:",omitzero,inline"`
	OfUser     *AccessPolicyApprovalProcedureUpdateParamsStepApproverUser     `json:",omitzero,inline"`
	paramUnion
}

func (u AccessPolicyApprovalProcedureUpdateParamsStepApproverUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAppOwner, u.OfGroup, u.OfManager, u.OfUser)
}
func (u *AccessPolicyApprovalProcedureUpdateParamsStepApproverUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AccessPolicyApprovalProcedureUpdateParamsStepApproverUnion) asAny() any {
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
func (u AccessPolicyApprovalProcedureUpdateParamsStepApproverUnion) GetAppOwner() *any {
	if vt := u.OfAppOwner; vt != nil {
		return &vt.AppOwner
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AccessPolicyApprovalProcedureUpdateParamsStepApproverUnion) GetGroup() *AccessPolicyApprovalProcedureUpdateParamsStepApproverGroupGroup {
	if vt := u.OfGroup; vt != nil {
		return &vt.Group
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AccessPolicyApprovalProcedureUpdateParamsStepApproverUnion) GetManager() *any {
	if vt := u.OfManager; vt != nil {
		return &vt.Manager
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AccessPolicyApprovalProcedureUpdateParamsStepApproverUnion) GetUser() *AccessPolicyApprovalProcedureUpdateParamsStepApproverUserUser {
	if vt := u.OfUser; vt != nil {
		return &vt.User
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AccessPolicyApprovalProcedureUpdateParamsStepApproverUnion) GetNotify() *bool {
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
type AccessPolicyApprovalProcedureUpdateParamsStepApproverAppOwner struct {
	// App owners as approvers. Only valid for access policy approval procedures.
	AppOwner any `json:"appOwner,omitzero" api:"required"`
	// Whether to notify this approver when the step is pending.
	Notify param.Opt[bool] `json:"notify,omitzero"`
	paramObj
}

func (r AccessPolicyApprovalProcedureUpdateParamsStepApproverAppOwner) MarshalJSON() (data []byte, err error) {
	type shadow AccessPolicyApprovalProcedureUpdateParamsStepApproverAppOwner
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccessPolicyApprovalProcedureUpdateParamsStepApproverAppOwner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Serval group as approvers.
//
// The property Group is required.
type AccessPolicyApprovalProcedureUpdateParamsStepApproverGroup struct {
	// A Serval group as approvers.
	Group AccessPolicyApprovalProcedureUpdateParamsStepApproverGroupGroup `json:"group,omitzero" api:"required"`
	// Whether to notify this approver when the step is pending.
	Notify param.Opt[bool] `json:"notify,omitzero"`
	paramObj
}

func (r AccessPolicyApprovalProcedureUpdateParamsStepApproverGroup) MarshalJSON() (data []byte, err error) {
	type shadow AccessPolicyApprovalProcedureUpdateParamsStepApproverGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccessPolicyApprovalProcedureUpdateParamsStepApproverGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Serval group as approvers.
type AccessPolicyApprovalProcedureUpdateParamsStepApproverGroupGroup struct {
	// The ID of the Serval group.
	GroupID param.Opt[string] `json:"groupId,omitzero"`
	paramObj
}

func (r AccessPolicyApprovalProcedureUpdateParamsStepApproverGroupGroup) MarshalJSON() (data []byte, err error) {
	type shadow AccessPolicyApprovalProcedureUpdateParamsStepApproverGroupGroup
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccessPolicyApprovalProcedureUpdateParamsStepApproverGroupGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The requester's manager as an approver.
//
// The property Manager is required.
type AccessPolicyApprovalProcedureUpdateParamsStepApproverManager struct {
	// The requester's manager as an approver.
	Manager any `json:"manager,omitzero" api:"required"`
	// Whether to notify this approver when the step is pending.
	Notify param.Opt[bool] `json:"notify,omitzero"`
	paramObj
}

func (r AccessPolicyApprovalProcedureUpdateParamsStepApproverManager) MarshalJSON() (data []byte, err error) {
	type shadow AccessPolicyApprovalProcedureUpdateParamsStepApproverManager
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccessPolicyApprovalProcedureUpdateParamsStepApproverManager) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A specific user as an approver.
//
// The property User is required.
type AccessPolicyApprovalProcedureUpdateParamsStepApproverUser struct {
	// A specific user as an approver.
	User AccessPolicyApprovalProcedureUpdateParamsStepApproverUserUser `json:"user,omitzero" api:"required"`
	// Whether to notify this approver when the step is pending.
	Notify param.Opt[bool] `json:"notify,omitzero"`
	paramObj
}

func (r AccessPolicyApprovalProcedureUpdateParamsStepApproverUser) MarshalJSON() (data []byte, err error) {
	type shadow AccessPolicyApprovalProcedureUpdateParamsStepApproverUser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccessPolicyApprovalProcedureUpdateParamsStepApproverUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A specific user as an approver.
type AccessPolicyApprovalProcedureUpdateParamsStepApproverUserUser struct {
	// The ID of the user.
	UserID param.Opt[string] `json:"userId,omitzero"`
	paramObj
}

func (r AccessPolicyApprovalProcedureUpdateParamsStepApproverUserUser) MarshalJSON() (data []byte, err error) {
	type shadow AccessPolicyApprovalProcedureUpdateParamsStepApproverUserUser
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccessPolicyApprovalProcedureUpdateParamsStepApproverUserUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for a custom workflow that determines approvers or auto-approves.
type AccessPolicyApprovalProcedureUpdateParamsStepCustomWorkflow struct {
	// The ID of the workflow to execute.
	WorkflowID param.Opt[string] `json:"workflowId,omitzero"`
	paramObj
}

func (r AccessPolicyApprovalProcedureUpdateParamsStepCustomWorkflow) MarshalJSON() (data []byte, err error) {
	type shadow AccessPolicyApprovalProcedureUpdateParamsStepCustomWorkflow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccessPolicyApprovalProcedureUpdateParamsStepCustomWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessPolicyApprovalProcedureUpdateResponseEnvelope struct {
	// The updated approval procedure.
	Data AccessPolicyApprovalProcedure `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessPolicyApprovalProcedureUpdateResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AccessPolicyApprovalProcedureUpdateResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessPolicyApprovalProcedureListResponseEnvelope struct {
	// The list of approval procedures (typically 0 or 1).
	Data []AccessPolicyApprovalProcedure `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessPolicyApprovalProcedureListResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AccessPolicyApprovalProcedureListResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessPolicyApprovalProcedureDeleteParams struct {
	// The ID of the access policy.
	AccessPolicyID string `path:"access_policy_id" api:"required" json:"-"`
	paramObj
}
