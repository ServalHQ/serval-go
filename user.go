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

// UserService contains methods and other services that help with interacting with
// the serval API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r UserService) {
	r = UserService{}
	r.Options = opts
	return
}

// Create a new user.
func (r *UserService) New(ctx context.Context, body UserNewParams, opts ...option.RequestOption) (res *User, err error) {
	var env UserNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "v2/users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Get a specific user by ID.
func (r *UserService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *User, err error) {
	var env UserGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/users/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Update an existing user.
func (r *UserService) Update(ctx context.Context, id string, body UserUpdateParams, opts ...option.RequestOption) (res *User, err error) {
	var env UserUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/users/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// List all users.
func (r *UserService) List(ctx context.Context, query UserListParams, opts ...option.RequestOption) (res *UserListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a user.
func (r *UserService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *UserDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/users/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type User struct {
	ID        string `json:"id"`
	AvatarURL string `json:"avatarUrl"`
	// A timestamp in RFC 3339 format (e.g., "2017-01-15T01:30:15.01Z").
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// A timestamp in RFC 3339 format (e.g., "2017-01-15T01:30:15.01Z").
	DeactivatedAt time.Time `json:"deactivatedAt,nullable" format:"date-time"`
	Email         string    `json:"email"`
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	Name          string    `json:"name"`
	// Any of "USER_ROLE_UNSPECIFIED", "USER_ROLE_ORG_MEMBER", "USER_ROLE_ORG_ADMIN",
	// "USER_ROLE_ORG_GUEST".
	Role UserRole `json:"role"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		AvatarURL     respjson.Field
		CreatedAt     respjson.Field
		DeactivatedAt respjson.Field
		Email         respjson.Field
		FirstName     respjson.Field
		LastName      respjson.Field
		Name          respjson.Field
		Role          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r User) RawJSON() string { return r.JSON.raw }
func (r *User) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserRole string

const (
	UserRoleUserRoleUnspecified UserRole = "USER_ROLE_UNSPECIFIED"
	UserRoleUserRoleOrgMember   UserRole = "USER_ROLE_ORG_MEMBER"
	UserRoleUserRoleOrgAdmin    UserRole = "USER_ROLE_ORG_ADMIN"
	UserRoleUserRoleOrgGuest    UserRole = "USER_ROLE_ORG_GUEST"
)

type UserListResponse struct {
	// The list of users.
	Data []User `json:"data"`
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
func (r UserListResponse) RawJSON() string { return r.JSON.raw }
func (r *UserListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserDeleteResponse = any

type UserNewParams struct {
	Email     string            `json:"email,required"`
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	LastName  param.Opt[string] `json:"lastName,omitzero"`
	// Any of "USER_ROLE_UNSPECIFIED", "USER_ROLE_ORG_MEMBER", "USER_ROLE_ORG_ADMIN",
	// "USER_ROLE_ORG_GUEST".
	Role UserNewParamsRole `json:"role,omitzero"`
	paramObj
}

func (r UserNewParams) MarshalJSON() (data []byte, err error) {
	type shadow UserNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserNewParamsRole string

const (
	UserNewParamsRoleUserRoleUnspecified UserNewParamsRole = "USER_ROLE_UNSPECIFIED"
	UserNewParamsRoleUserRoleOrgMember   UserNewParamsRole = "USER_ROLE_ORG_MEMBER"
	UserNewParamsRoleUserRoleOrgAdmin    UserNewParamsRole = "USER_ROLE_ORG_ADMIN"
	UserNewParamsRoleUserRoleOrgGuest    UserNewParamsRole = "USER_ROLE_ORG_GUEST"
)

type UserNewResponseEnvelope struct {
	Data User `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *UserNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGetResponseEnvelope struct {
	Data User `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *UserGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserUpdateParams struct {
	AvatarURL param.Opt[string] `json:"avatarUrl,omitzero"`
	Email     param.Opt[string] `json:"email,omitzero"`
	FirstName param.Opt[string] `json:"firstName,omitzero"`
	LastName  param.Opt[string] `json:"lastName,omitzero"`
	// Any of "USER_ROLE_UNSPECIFIED", "USER_ROLE_ORG_MEMBER", "USER_ROLE_ORG_ADMIN",
	// "USER_ROLE_ORG_GUEST".
	Role UserUpdateParamsRole `json:"role,omitzero"`
	paramObj
}

func (r UserUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow UserUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserUpdateParamsRole string

const (
	UserUpdateParamsRoleUserRoleUnspecified UserUpdateParamsRole = "USER_ROLE_UNSPECIFIED"
	UserUpdateParamsRoleUserRoleOrgMember   UserUpdateParamsRole = "USER_ROLE_ORG_MEMBER"
	UserUpdateParamsRoleUserRoleOrgAdmin    UserUpdateParamsRole = "USER_ROLE_ORG_ADMIN"
	UserUpdateParamsRoleUserRoleOrgGuest    UserUpdateParamsRole = "USER_ROLE_ORG_GUEST"
)

type UserUpdateResponseEnvelope struct {
	Data User `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserUpdateResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *UserUpdateResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserListParams struct {
	// Whether to include deactivated users in the response.
	IncludeDeactivated param.Opt[bool] `query:"includeDeactivated,omitzero" json:"-"`
	// Maximum number of results to return. Default is 1000, maximum is 1000.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Token for pagination. Leave empty for the first request.
	PageToken param.Opt[string] `query:"pageToken,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserListParams]'s query parameters as `url.Values`.
func (r UserListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
