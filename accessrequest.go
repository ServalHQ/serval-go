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

// AccessRequestService contains methods and other services that help with
// interacting with the serval API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessRequestService] method instead.
type AccessRequestService struct {
	Options []option.RequestOption
}

// NewAccessRequestService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessRequestService(opts ...option.RequestOption) (r AccessRequestService) {
	r = AccessRequestService{}
	r.Options = opts
	return
}

// Get a specific access request by ID.
func (r *AccessRequestService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AccessRequest, err error) {
	var env AccessRequestGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/access-requests/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// List access requests for a team. Filter by user, entitlement, app instance,
// status, or time range.
func (r *AccessRequestService) List(ctx context.Context, query AccessRequestListParams, opts ...option.RequestOption) (res *AccessRequestListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/access-requests"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Search access requests with filters. Supports filtering by entitlement, app
// instance, status, user IDs, time range, and linked ticket.
func (r *AccessRequestService) Search(ctx context.Context, body AccessRequestSearchParams, opts ...option.RequestOption) (res *AccessRequestSearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/access-requests/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// AccessRequest represents a request for access to an entitlement.
type AccessRequest struct {
	// The unique ID of the access request.
	ID string `json:"id"`
	// The number of minutes of access requested.
	AccessMinutes int64 `json:"accessMinutes"`
	// The business justification provided for the request.
	BusinessJustification string `json:"businessJustification,nullable"`
	// The timestamp when the access request was created.
	CreatedAt string `json:"createdAt"`
	// The timestamp when the access expires (if approved and active).
	ExpiresAt string `json:"expiresAt,nullable"`
	// The ID of the linked ticket, if any.
	LinkedTicketID string `json:"linkedTicketId,nullable"`
	// The ID of the user who requested access.
	RequestedByUserID string `json:"requestedByUserId"`
	// The ID of the requested role.
	RequestedRoleID string `json:"requestedRoleId"`
	// The status of the access request.
	//
	// Any of "ACCESS_REQUEST_STATUS_UNSPECIFIED", "ACCESS_REQUEST_STATUS_PENDING",
	// "ACCESS_REQUEST_STATUS_APPROVED", "ACCESS_REQUEST_STATUS_DENIED",
	// "ACCESS_REQUEST_STATUS_EXPIRED", "ACCESS_REQUEST_STATUS_REVOKED",
	// "ACCESS_REQUEST_STATUS_CANCELED", "ACCESS_REQUEST_STATUS_FAILED".
	Status AccessRequestStatus `json:"status"`
	// The ID of the target user for whom access was requested.
	TargetUserID string `json:"targetUserId"`
	// The ID of the team that the access request belongs to.
	TeamID string `json:"teamId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		AccessMinutes         respjson.Field
		BusinessJustification respjson.Field
		CreatedAt             respjson.Field
		ExpiresAt             respjson.Field
		LinkedTicketID        respjson.Field
		RequestedByUserID     respjson.Field
		RequestedRoleID       respjson.Field
		Status                respjson.Field
		TargetUserID          respjson.Field
		TeamID                respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessRequest) RawJSON() string { return r.JSON.raw }
func (r *AccessRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the access request.
type AccessRequestStatus string

const (
	AccessRequestStatusAccessRequestStatusUnspecified AccessRequestStatus = "ACCESS_REQUEST_STATUS_UNSPECIFIED"
	AccessRequestStatusAccessRequestStatusPending     AccessRequestStatus = "ACCESS_REQUEST_STATUS_PENDING"
	AccessRequestStatusAccessRequestStatusApproved    AccessRequestStatus = "ACCESS_REQUEST_STATUS_APPROVED"
	AccessRequestStatusAccessRequestStatusDenied      AccessRequestStatus = "ACCESS_REQUEST_STATUS_DENIED"
	AccessRequestStatusAccessRequestStatusExpired     AccessRequestStatus = "ACCESS_REQUEST_STATUS_EXPIRED"
	AccessRequestStatusAccessRequestStatusRevoked     AccessRequestStatus = "ACCESS_REQUEST_STATUS_REVOKED"
	AccessRequestStatusAccessRequestStatusCanceled    AccessRequestStatus = "ACCESS_REQUEST_STATUS_CANCELED"
	AccessRequestStatusAccessRequestStatusFailed      AccessRequestStatus = "ACCESS_REQUEST_STATUS_FAILED"
)

type AccessRequestListResponse struct {
	// The list of access requests.
	Data []AccessRequest `json:"data"`
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
func (r AccessRequestListResponse) RawJSON() string { return r.JSON.raw }
func (r *AccessRequestListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessRequestSearchResponse struct {
	// The list of access requests.
	Data []AccessRequest `json:"data"`
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
func (r AccessRequestSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *AccessRequestSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessRequestGetResponseEnvelope struct {
	// The access request.
	Data AccessRequest `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessRequestGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AccessRequestGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessRequestListParams struct {
	// Maximum number of results to return. Default is 100, maximum is 1000.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Token for pagination. Leave empty for the first request.
	PageToken param.Opt[string] `query:"pageToken,omitzero" json:"-"`
	// The ID of the team. Required.
	TeamID param.Opt[string] `query:"teamId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AccessRequestListParams]'s query parameters as
// `url.Values`.
func (r AccessRequestListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccessRequestSearchParams struct {
	// Filter by app instance ID.
	AppInstanceID param.Opt[string] `json:"appInstanceId,omitzero"`
	// Filter by requests created after this timestamp (RFC3339 format).
	CreatedAfter param.Opt[string] `json:"createdAfter,omitzero"`
	// Filter by requests created before this timestamp (RFC3339 format).
	CreatedBefore param.Opt[string] `json:"createdBefore,omitzero"`
	// Filter by linked ticket ID.
	LinkedTicketID param.Opt[string] `json:"linkedTicketId,omitzero"`
	// Maximum number of results to return. Default is 100, maximum is 1000.
	PageSize param.Opt[int64] `json:"pageSize,omitzero"`
	// Token for pagination. Leave empty for the first request.
	PageToken param.Opt[string] `json:"pageToken,omitzero"`
	// Filter by the user who requested access.
	RequestedByUserID param.Opt[string] `json:"requestedByUserId,omitzero"`
	// Filter by requested role ID.
	RequestedRoleID param.Opt[string] `json:"requestedRoleId,omitzero"`
	// Filter by the target user for whom access was requested.
	TargetUserID param.Opt[string] `json:"targetUserId,omitzero"`
	// The ID of the team. Required.
	TeamID param.Opt[string] `json:"teamId,omitzero"`
	// Filter by statuses (multiple allowed).
	//
	// Any of "ACCESS_REQUEST_STATUS_UNSPECIFIED", "ACCESS_REQUEST_STATUS_PENDING",
	// "ACCESS_REQUEST_STATUS_APPROVED", "ACCESS_REQUEST_STATUS_DENIED",
	// "ACCESS_REQUEST_STATUS_EXPIRED", "ACCESS_REQUEST_STATUS_REVOKED",
	// "ACCESS_REQUEST_STATUS_CANCELED", "ACCESS_REQUEST_STATUS_FAILED".
	Statuses []string `json:"statuses,omitzero"`
	paramObj
}

func (r AccessRequestSearchParams) MarshalJSON() (data []byte, err error) {
	type shadow AccessRequestSearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccessRequestSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
