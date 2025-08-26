// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package serval

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/serval-go/internal/apijson"
	"github.com/stainless-sdks/serval-go/internal/requestconfig"
	"github.com/stainless-sdks/serval-go/option"
	"github.com/stainless-sdks/serval-go/packages/param"
	"github.com/stainless-sdks/serval-go/packages/respjson"
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
func (r *AccessPolicyApprovalProcedureService) New(ctx context.Context, accessPolicyID string, body AccessPolicyApprovalProcedureNewParams, opts ...option.RequestOption) (res *AccessPolicyApprovalProcedureNewResponse, err error) {
	var env AccessPolicyApprovalProcedureNewResponseEnvelope
	opts = append(r.Options[:], opts...)
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
func (r *AccessPolicyApprovalProcedureService) Get(ctx context.Context, id string, query AccessPolicyApprovalProcedureGetParams, opts ...option.RequestOption) (res *AccessPolicyApprovalProcedureGetResponse, err error) {
	var env AccessPolicyApprovalProcedureGetResponseEnvelope
	opts = append(r.Options[:], opts...)
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
func (r *AccessPolicyApprovalProcedureService) Update(ctx context.Context, id string, params AccessPolicyApprovalProcedureUpdateParams, opts ...option.RequestOption) (res *AccessPolicyApprovalProcedureUpdateResponse, err error) {
	var env AccessPolicyApprovalProcedureUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
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
func (r *AccessPolicyApprovalProcedureService) List(ctx context.Context, accessPolicyID string, opts ...option.RequestOption) (res *[]AccessPolicyApprovalProcedureListResponse, err error) {
	var env AccessPolicyApprovalProcedureListResponseEnvelope
	opts = append(r.Options[:], opts...)
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
	opts = append(r.Options[:], opts...)
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

// The created approval procedure.
type AccessPolicyApprovalProcedureNewResponse struct {
	// The ID of the access policy approval procedure.
	ID string `json:"id"`
	// The steps in the approval procedure.
	Steps []AccessPolicyApprovalProcedureNewResponseStep `json:"steps"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Steps       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessPolicyApprovalProcedureNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AccessPolicyApprovalProcedureNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessPolicyApprovalProcedureNewResponseStep struct {
	// The ID of the approval step.
	ID string `json:"id"`
	// Whether the step can be approved by the requester themselves.
	AllowSelfApproval bool `json:"allowSelfApproval"`
	// The IDs of the Serval groups that can approve the step.
	ServalGroupIDs []string `json:"servalGroupIds"`
	// The IDs of the specific users that can approve the step.
	SpecificUserIDs []string `json:"specificUserIds"`
	// The type of approval step.
	//
	// Any of "APPROVAL_PROCEDURE_STEP_TYPE_UNSPECIFIED", "SPECIFIC_USERS",
	// "SERVAL_GROUPS".
	StepType string `json:"stepType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AllowSelfApproval respjson.Field
		ServalGroupIDs    respjson.Field
		SpecificUserIDs   respjson.Field
		StepType          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessPolicyApprovalProcedureNewResponseStep) RawJSON() string { return r.JSON.raw }
func (r *AccessPolicyApprovalProcedureNewResponseStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The approval procedure.
type AccessPolicyApprovalProcedureGetResponse struct {
	// The ID of the access policy approval procedure.
	ID string `json:"id"`
	// The steps in the approval procedure.
	Steps []AccessPolicyApprovalProcedureGetResponseStep `json:"steps"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Steps       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessPolicyApprovalProcedureGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AccessPolicyApprovalProcedureGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessPolicyApprovalProcedureGetResponseStep struct {
	// The ID of the approval step.
	ID string `json:"id"`
	// Whether the step can be approved by the requester themselves.
	AllowSelfApproval bool `json:"allowSelfApproval"`
	// The IDs of the Serval groups that can approve the step.
	ServalGroupIDs []string `json:"servalGroupIds"`
	// The IDs of the specific users that can approve the step.
	SpecificUserIDs []string `json:"specificUserIds"`
	// The type of approval step.
	//
	// Any of "APPROVAL_PROCEDURE_STEP_TYPE_UNSPECIFIED", "SPECIFIC_USERS",
	// "SERVAL_GROUPS".
	StepType string `json:"stepType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AllowSelfApproval respjson.Field
		ServalGroupIDs    respjson.Field
		SpecificUserIDs   respjson.Field
		StepType          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessPolicyApprovalProcedureGetResponseStep) RawJSON() string { return r.JSON.raw }
func (r *AccessPolicyApprovalProcedureGetResponseStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The updated approval procedure.
type AccessPolicyApprovalProcedureUpdateResponse struct {
	// The ID of the access policy approval procedure.
	ID string `json:"id"`
	// The steps in the approval procedure.
	Steps []AccessPolicyApprovalProcedureUpdateResponseStep `json:"steps"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Steps       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessPolicyApprovalProcedureUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *AccessPolicyApprovalProcedureUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessPolicyApprovalProcedureUpdateResponseStep struct {
	// The ID of the approval step.
	ID string `json:"id"`
	// Whether the step can be approved by the requester themselves.
	AllowSelfApproval bool `json:"allowSelfApproval"`
	// The IDs of the Serval groups that can approve the step.
	ServalGroupIDs []string `json:"servalGroupIds"`
	// The IDs of the specific users that can approve the step.
	SpecificUserIDs []string `json:"specificUserIds"`
	// The type of approval step.
	//
	// Any of "APPROVAL_PROCEDURE_STEP_TYPE_UNSPECIFIED", "SPECIFIC_USERS",
	// "SERVAL_GROUPS".
	StepType string `json:"stepType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AllowSelfApproval respjson.Field
		ServalGroupIDs    respjson.Field
		SpecificUserIDs   respjson.Field
		StepType          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessPolicyApprovalProcedureUpdateResponseStep) RawJSON() string { return r.JSON.raw }
func (r *AccessPolicyApprovalProcedureUpdateResponseStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessPolicyApprovalProcedureListResponse struct {
	// The ID of the access policy approval procedure.
	ID string `json:"id"`
	// The steps in the approval procedure.
	Steps []AccessPolicyApprovalProcedureListResponseStep `json:"steps"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Steps       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessPolicyApprovalProcedureListResponse) RawJSON() string { return r.JSON.raw }
func (r *AccessPolicyApprovalProcedureListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessPolicyApprovalProcedureListResponseStep struct {
	// The ID of the approval step.
	ID string `json:"id"`
	// Whether the step can be approved by the requester themselves.
	AllowSelfApproval bool `json:"allowSelfApproval"`
	// The IDs of the Serval groups that can approve the step.
	ServalGroupIDs []string `json:"servalGroupIds"`
	// The IDs of the specific users that can approve the step.
	SpecificUserIDs []string `json:"specificUserIds"`
	// The type of approval step.
	//
	// Any of "APPROVAL_PROCEDURE_STEP_TYPE_UNSPECIFIED", "SPECIFIC_USERS",
	// "SERVAL_GROUPS".
	StepType string `json:"stepType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AllowSelfApproval respjson.Field
		ServalGroupIDs    respjson.Field
		SpecificUserIDs   respjson.Field
		StepType          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessPolicyApprovalProcedureListResponseStep) RawJSON() string { return r.JSON.raw }
func (r *AccessPolicyApprovalProcedureListResponseStep) UnmarshalJSON(data []byte) error {
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
	// The ID of the approval step.
	ID param.Opt[string] `json:"id,omitzero"`
	// Whether the step can be approved by the requester themselves.
	AllowSelfApproval param.Opt[bool] `json:"allowSelfApproval,omitzero"`
	// The IDs of the Serval groups that can approve the step.
	ServalGroupIDs []string `json:"servalGroupIds,omitzero"`
	// The IDs of the specific users that can approve the step.
	SpecificUserIDs []string `json:"specificUserIds,omitzero"`
	// The type of approval step.
	//
	// Any of "APPROVAL_PROCEDURE_STEP_TYPE_UNSPECIFIED", "SPECIFIC_USERS",
	// "SERVAL_GROUPS".
	StepType string `json:"stepType,omitzero"`
	paramObj
}

func (r AccessPolicyApprovalProcedureNewParamsStep) MarshalJSON() (data []byte, err error) {
	type shadow AccessPolicyApprovalProcedureNewParamsStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccessPolicyApprovalProcedureNewParamsStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AccessPolicyApprovalProcedureNewParamsStep](
		"stepType", "APPROVAL_PROCEDURE_STEP_TYPE_UNSPECIFIED", "SPECIFIC_USERS", "SERVAL_GROUPS",
	)
}

type AccessPolicyApprovalProcedureNewResponseEnvelope struct {
	// The created approval procedure.
	Data AccessPolicyApprovalProcedureNewResponse `json:"data"`
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
	AccessPolicyID string `path:"access_policy_id,required" json:"-"`
	paramObj
}

type AccessPolicyApprovalProcedureGetResponseEnvelope struct {
	// The approval procedure.
	Data AccessPolicyApprovalProcedureGetResponse `json:"data"`
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
	AccessPolicyID string `path:"access_policy_id,required" json:"-"`
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
	// The ID of the approval step.
	ID param.Opt[string] `json:"id,omitzero"`
	// Whether the step can be approved by the requester themselves.
	AllowSelfApproval param.Opt[bool] `json:"allowSelfApproval,omitzero"`
	// The IDs of the Serval groups that can approve the step.
	ServalGroupIDs []string `json:"servalGroupIds,omitzero"`
	// The IDs of the specific users that can approve the step.
	SpecificUserIDs []string `json:"specificUserIds,omitzero"`
	// The type of approval step.
	//
	// Any of "APPROVAL_PROCEDURE_STEP_TYPE_UNSPECIFIED", "SPECIFIC_USERS",
	// "SERVAL_GROUPS".
	StepType string `json:"stepType,omitzero"`
	paramObj
}

func (r AccessPolicyApprovalProcedureUpdateParamsStep) MarshalJSON() (data []byte, err error) {
	type shadow AccessPolicyApprovalProcedureUpdateParamsStep
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccessPolicyApprovalProcedureUpdateParamsStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AccessPolicyApprovalProcedureUpdateParamsStep](
		"stepType", "APPROVAL_PROCEDURE_STEP_TYPE_UNSPECIFIED", "SPECIFIC_USERS", "SERVAL_GROUPS",
	)
}

type AccessPolicyApprovalProcedureUpdateResponseEnvelope struct {
	// The updated approval procedure.
	Data AccessPolicyApprovalProcedureUpdateResponse `json:"data"`
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
	// The list of approval procedures.
	Data []AccessPolicyApprovalProcedureListResponse `json:"data"`
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
	AccessPolicyID string `path:"access_policy_id,required" json:"-"`
	paramObj
}
