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
	"github.com/ServalHQ/serval-go/packages/param"
	"github.com/ServalHQ/serval-go/packages/respjson"
)

// GroupService contains methods and other services that help with interacting with
// the serval API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGroupService] method instead.
type GroupService struct {
	Options []option.RequestOption
}

// NewGroupService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewGroupService(opts ...option.RequestOption) (r GroupService) {
	r = GroupService{}
	r.Options = opts
	return
}

// Create a new group.
func (r *GroupService) New(ctx context.Context, body GroupNewParams, opts ...option.RequestOption) (res *Group, err error) {
	var env GroupNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "v2/groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Get a specific group by ID.
func (r *GroupService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Group, err error) {
	var env GroupGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/groups/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Update an existing group.
func (r *GroupService) Update(ctx context.Context, id string, body GroupUpdateParams, opts ...option.RequestOption) (res *Group, err error) {
	var env GroupUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/groups/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// List all groups.
func (r *GroupService) List(ctx context.Context, query GroupListParams, opts ...option.RequestOption) (res *GroupListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a group.
func (r *GroupService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *GroupDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/groups/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Group struct {
	ID string `json:"id"`
	// A timestamp in RFC 3339 format (e.g., "2017-01-15T01:30:15.01Z").
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// A timestamp in RFC 3339 format (e.g., "2017-01-15T01:30:15.01Z").
	DeletedAt      time.Time `json:"deletedAt,nullable" format:"date-time"`
	Name           string    `json:"name"`
	OrganizationID string    `json:"organizationId"`
	UserIDs        []string  `json:"userIds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		DeletedAt      respjson.Field
		Name           respjson.Field
		OrganizationID respjson.Field
		UserIDs        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Group) RawJSON() string { return r.JSON.raw }
func (r *Group) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GroupListResponse struct {
	// The list of groups.
	Data []Group `json:"data"`
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
func (r GroupListResponse) RawJSON() string { return r.JSON.raw }
func (r *GroupListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GroupDeleteResponse = any

type GroupNewParams struct {
	Name    param.Opt[string] `json:"name,omitzero"`
	UserIDs []string          `json:"userIds,omitzero"`
	paramObj
}

func (r GroupNewParams) MarshalJSON() (data []byte, err error) {
	type shadow GroupNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GroupNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GroupNewResponseEnvelope struct {
	Data Group `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GroupNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *GroupNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GroupGetResponseEnvelope struct {
	Data Group `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GroupGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *GroupGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GroupUpdateParams struct {
	Name    param.Opt[string] `json:"name,omitzero"`
	UserIDs []string          `json:"userIds,omitzero"`
	paramObj
}

func (r GroupUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow GroupUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GroupUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GroupUpdateResponseEnvelope struct {
	Data Group `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GroupUpdateResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *GroupUpdateResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GroupListParams struct {
	// Maximum number of results to return. Default is 5000, maximum is 5000.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Token for pagination. Leave empty for the first request.
	PageToken param.Opt[string] `query:"pageToken,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GroupListParams]'s query parameters as `url.Values`.
func (r GroupListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
