// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package serval

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/ServalHQ/serval-go/internal/apijson"
	"github.com/ServalHQ/serval-go/internal/apiquery"
	"github.com/ServalHQ/serval-go/internal/requestconfig"
	"github.com/ServalHQ/serval-go/option"
	"github.com/ServalHQ/serval-go/packages/pagination"
	"github.com/ServalHQ/serval-go/packages/param"
	"github.com/ServalHQ/serval-go/packages/respjson"
)

// ApprovalDelegationService contains methods and other services that help with
// interacting with the serval API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewApprovalDelegationService] method instead.
type ApprovalDelegationService struct {
	Options []option.RequestOption
}

// NewApprovalDelegationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewApprovalDelegationService(opts ...option.RequestOption) (r ApprovalDelegationService) {
	r = ApprovalDelegationService{}
	r.Options = opts
	return
}

// Create a new approval delegation rule. Optionally specify a delegator_user_id to
// create on behalf of another user.
func (r *ApprovalDelegationService) New(ctx context.Context, body ApprovalDelegationNewParams, opts ...option.RequestOption) (res *ApprovalDelegation, err error) {
	var env ApprovalDelegationNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "v2/approval-delegations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// Get an approval delegation rule by its ID.
func (r *ApprovalDelegationService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *ApprovalDelegation, err error) {
	var env ApprovalDelegationGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/approval-delegations/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// Update an existing approval delegation rule.
func (r *ApprovalDelegationService) Update(ctx context.Context, id string, body ApprovalDelegationUpdateParams, opts ...option.RequestOption) (res *ApprovalDelegation, err error) {
	var env ApprovalDelegationUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/approval-delegations/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// List approval delegation rules. Optionally filter by delegator_user_id.
func (r *ApprovalDelegationService) List(ctx context.Context, query ApprovalDelegationListParams, opts ...option.RequestOption) (res *pagination.CursorPage[ApprovalDelegation], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v2/approval-delegations"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List approval delegation rules. Optionally filter by delegator_user_id.
func (r *ApprovalDelegationService) ListAutoPaging(ctx context.Context, query ApprovalDelegationListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[ApprovalDelegation] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Delete an approval delegation rule.
func (r *ApprovalDelegationService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *ApprovalDelegationDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/approval-delegations/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// An approval delegation rule that allows a user to delegate their approval
// responsibilities.
type ApprovalDelegation struct {
	// The ID of the approval delegation.
	ID string `json:"id"`
	// The timestamp when the delegation was created.
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// The delegates who can approve on behalf of the delegator.
	Delegates []ApprovalDelegationDelegate `json:"delegates"`
	// The ID of the user who is delegating their approvals.
	DelegatorUserID string `json:"delegatorUserId"`
	// A description of the delegation rule.
	Description string `json:"description"`
	// The priority of the delegation rule (lower values are higher priority).
	Priority int64 `json:"priority"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		Delegates       respjson.Field
		DelegatorUserID respjson.Field
		Description     respjson.Field
		Priority        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApprovalDelegation) RawJSON() string { return r.JSON.raw }
func (r *ApprovalDelegation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A delegate who can approve on behalf of a delegator.
type ApprovalDelegationDelegate struct {
	// The ID of the delegate (user ID or group ID, depending on type).
	ID string `json:"id"`
	// The type of delegate (user or group).
	//
	// Any of "APPROVAL_DELEGATE_TYPE_UNSPECIFIED", "APPROVAL_DELEGATE_TYPE_USER",
	// "APPROVAL_DELEGATE_TYPE_GROUP".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApprovalDelegationDelegate) RawJSON() string { return r.JSON.raw }
func (r *ApprovalDelegationDelegate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApprovalDelegationDeleteResponse = any

type ApprovalDelegationNewParams struct {
	// The delegates who can approve on behalf of the delegator.
	Delegates []ApprovalDelegationNewParamsDelegate `json:"delegates,omitzero" api:"required"`
	// The ID of the user who is delegating their approvals. When omitted, defaults to
	// the authenticated user.
	DelegatorUserID param.Opt[string] `json:"delegatorUserId,omitzero"`
	// A description of the delegation rule.
	Description param.Opt[string] `json:"description,omitzero"`
	// The priority of the delegation rule (lower values are higher priority).
	Priority param.Opt[int64] `json:"priority,omitzero"`
	paramObj
}

func (r ApprovalDelegationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ApprovalDelegationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ApprovalDelegationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A delegate who can approve on behalf of a delegator.
type ApprovalDelegationNewParamsDelegate struct {
	// The ID of the delegate (user ID or group ID, depending on type).
	ID param.Opt[string] `json:"id,omitzero"`
	// The type of delegate (user or group).
	//
	// Any of "APPROVAL_DELEGATE_TYPE_UNSPECIFIED", "APPROVAL_DELEGATE_TYPE_USER",
	// "APPROVAL_DELEGATE_TYPE_GROUP".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ApprovalDelegationNewParamsDelegate) MarshalJSON() (data []byte, err error) {
	type shadow ApprovalDelegationNewParamsDelegate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ApprovalDelegationNewParamsDelegate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ApprovalDelegationNewParamsDelegate](
		"type", "APPROVAL_DELEGATE_TYPE_UNSPECIFIED", "APPROVAL_DELEGATE_TYPE_USER", "APPROVAL_DELEGATE_TYPE_GROUP",
	)
}

type ApprovalDelegationNewResponseEnvelope struct {
	// An approval delegation rule that allows a user to delegate their approval
	// responsibilities.
	Data ApprovalDelegation `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApprovalDelegationNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *ApprovalDelegationNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApprovalDelegationGetResponseEnvelope struct {
	// An approval delegation rule that allows a user to delegate their approval
	// responsibilities.
	Data ApprovalDelegation `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApprovalDelegationGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *ApprovalDelegationGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApprovalDelegationUpdateParams struct {
	// A description of the delegation rule.
	Description param.Opt[string] `json:"description,omitzero"`
	// The priority of the delegation rule (lower values are higher priority).
	Priority param.Opt[int64] `json:"priority,omitzero"`
	// The delegates who can approve on behalf of the delegator.
	Delegates []ApprovalDelegationUpdateParamsDelegate `json:"delegates,omitzero"`
	paramObj
}

func (r ApprovalDelegationUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ApprovalDelegationUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ApprovalDelegationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A delegate who can approve on behalf of a delegator.
type ApprovalDelegationUpdateParamsDelegate struct {
	// The ID of the delegate (user ID or group ID, depending on type).
	ID param.Opt[string] `json:"id,omitzero"`
	// The type of delegate (user or group).
	//
	// Any of "APPROVAL_DELEGATE_TYPE_UNSPECIFIED", "APPROVAL_DELEGATE_TYPE_USER",
	// "APPROVAL_DELEGATE_TYPE_GROUP".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ApprovalDelegationUpdateParamsDelegate) MarshalJSON() (data []byte, err error) {
	type shadow ApprovalDelegationUpdateParamsDelegate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ApprovalDelegationUpdateParamsDelegate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ApprovalDelegationUpdateParamsDelegate](
		"type", "APPROVAL_DELEGATE_TYPE_UNSPECIFIED", "APPROVAL_DELEGATE_TYPE_USER", "APPROVAL_DELEGATE_TYPE_GROUP",
	)
}

type ApprovalDelegationUpdateResponseEnvelope struct {
	// An approval delegation rule that allows a user to delegate their approval
	// responsibilities.
	Data ApprovalDelegation `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApprovalDelegationUpdateResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *ApprovalDelegationUpdateResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApprovalDelegationListParams struct {
	// Optional user ID to list delegations for a specific user.
	DelegatorUserID param.Opt[string] `query:"delegatorUserId,omitzero" json:"-"`
	// Maximum number of results to return.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Token for pagination. Leave empty for the first request.
	PageToken param.Opt[string] `query:"pageToken,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ApprovalDelegationListParams]'s query parameters as
// `url.Values`.
func (r ApprovalDelegationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
