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

// TeamUserService contains methods and other services that help with interacting
// with the serval API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTeamUserService] method instead.
type TeamUserService struct {
	Options []option.RequestOption
}

// NewTeamUserService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTeamUserService(opts ...option.RequestOption) (r TeamUserService) {
	r = TeamUserService{}
	r.Options = opts
	return
}

// Add a user to a team with a specified role.
func (r *TeamUserService) New(ctx context.Context, body TeamUserNewParams, opts ...option.RequestOption) (res *TeamUser, err error) {
	var env TeamUserNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "v2/team-users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// Get a specific team user by team ID and user ID, or by composite ID.
func (r *TeamUserService) Get(ctx context.Context, id string, query TeamUserGetParams, opts ...option.RequestOption) (res *TeamUser, err error) {
	var env TeamUserGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/team-users/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// Update a team user's role.
func (r *TeamUserService) Update(ctx context.Context, id string, body TeamUserUpdateParams, opts ...option.RequestOption) (res *TeamUser, err error) {
	var env TeamUserUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/team-users/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// List team users. Filter by team_id and/or user_id.
func (r *TeamUserService) List(ctx context.Context, query TeamUserListParams, opts ...option.RequestOption) (res *pagination.CursorPage[TeamUser], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v2/team-users"
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

// List team users. Filter by team_id and/or user_id.
func (r *TeamUserService) ListAutoPaging(ctx context.Context, query TeamUserListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[TeamUser] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Remove a user from a team.
func (r *TeamUserService) Delete(ctx context.Context, id string, body TeamUserDeleteParams, opts ...option.RequestOption) (res *TeamUserDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/team-users/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

type TeamUser struct {
	// Composite identifier in the format {team_id}:{user_id}.
	ID string `json:"id"`
	// A timestamp in RFC 3339 format (e.g., "2025-01-15T01:30:15Z").
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Any of "TEAM_USER_ROLE_UNSPECIFIED", "TEAM_USER_ROLE_AGENT",
	// "TEAM_USER_ROLE_MANAGER", "TEAM_USER_ROLE_BUILDER", "TEAM_USER_ROLE_VIEWER",
	// "TEAM_USER_ROLE_CONTRIBUTOR".
	Role   TeamUserRole `json:"role"`
	TeamID string       `json:"teamId"`
	UserID string       `json:"userId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Role        respjson.Field
		TeamID      respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TeamUser) RawJSON() string { return r.JSON.raw }
func (r *TeamUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TeamUserRole string

const (
	TeamUserRoleTeamUserRoleUnspecified TeamUserRole = "TEAM_USER_ROLE_UNSPECIFIED"
	TeamUserRoleTeamUserRoleAgent       TeamUserRole = "TEAM_USER_ROLE_AGENT"
	TeamUserRoleTeamUserRoleManager     TeamUserRole = "TEAM_USER_ROLE_MANAGER"
	TeamUserRoleTeamUserRoleBuilder     TeamUserRole = "TEAM_USER_ROLE_BUILDER"
	TeamUserRoleTeamUserRoleViewer      TeamUserRole = "TEAM_USER_ROLE_VIEWER"
	TeamUserRoleTeamUserRoleContributor TeamUserRole = "TEAM_USER_ROLE_CONTRIBUTOR"
)

type TeamUserDeleteResponse = any

type TeamUserNewParams struct {
	// Any of "TEAM_USER_ROLE_UNSPECIFIED", "TEAM_USER_ROLE_AGENT",
	// "TEAM_USER_ROLE_MANAGER", "TEAM_USER_ROLE_BUILDER", "TEAM_USER_ROLE_VIEWER",
	// "TEAM_USER_ROLE_CONTRIBUTOR".
	Role   TeamUserNewParamsRole `json:"role,omitzero" api:"required"`
	TeamID string                `json:"teamId" api:"required"`
	UserID string                `json:"userId" api:"required"`
	paramObj
}

func (r TeamUserNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TeamUserNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TeamUserNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TeamUserNewParamsRole string

const (
	TeamUserNewParamsRoleTeamUserRoleUnspecified TeamUserNewParamsRole = "TEAM_USER_ROLE_UNSPECIFIED"
	TeamUserNewParamsRoleTeamUserRoleAgent       TeamUserNewParamsRole = "TEAM_USER_ROLE_AGENT"
	TeamUserNewParamsRoleTeamUserRoleManager     TeamUserNewParamsRole = "TEAM_USER_ROLE_MANAGER"
	TeamUserNewParamsRoleTeamUserRoleBuilder     TeamUserNewParamsRole = "TEAM_USER_ROLE_BUILDER"
	TeamUserNewParamsRoleTeamUserRoleViewer      TeamUserNewParamsRole = "TEAM_USER_ROLE_VIEWER"
	TeamUserNewParamsRoleTeamUserRoleContributor TeamUserNewParamsRole = "TEAM_USER_ROLE_CONTRIBUTOR"
)

type TeamUserNewResponseEnvelope struct {
	Data TeamUser `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TeamUserNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *TeamUserNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TeamUserGetParams struct {
	TeamID param.Opt[string] `query:"teamId,omitzero" json:"-"`
	UserID param.Opt[string] `query:"userId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TeamUserGetParams]'s query parameters as `url.Values`.
func (r TeamUserGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TeamUserGetResponseEnvelope struct {
	Data TeamUser `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TeamUserGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *TeamUserGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TeamUserUpdateParams struct {
	// Any of "TEAM_USER_ROLE_UNSPECIFIED", "TEAM_USER_ROLE_AGENT",
	// "TEAM_USER_ROLE_MANAGER", "TEAM_USER_ROLE_BUILDER", "TEAM_USER_ROLE_VIEWER",
	// "TEAM_USER_ROLE_CONTRIBUTOR".
	Role   TeamUserUpdateParamsRole `json:"role,omitzero" api:"required"`
	TeamID param.Opt[string]        `json:"teamId,omitzero"`
	UserID param.Opt[string]        `json:"userId,omitzero"`
	paramObj
}

func (r TeamUserUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TeamUserUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TeamUserUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TeamUserUpdateParamsRole string

const (
	TeamUserUpdateParamsRoleTeamUserRoleUnspecified TeamUserUpdateParamsRole = "TEAM_USER_ROLE_UNSPECIFIED"
	TeamUserUpdateParamsRoleTeamUserRoleAgent       TeamUserUpdateParamsRole = "TEAM_USER_ROLE_AGENT"
	TeamUserUpdateParamsRoleTeamUserRoleManager     TeamUserUpdateParamsRole = "TEAM_USER_ROLE_MANAGER"
	TeamUserUpdateParamsRoleTeamUserRoleBuilder     TeamUserUpdateParamsRole = "TEAM_USER_ROLE_BUILDER"
	TeamUserUpdateParamsRoleTeamUserRoleViewer      TeamUserUpdateParamsRole = "TEAM_USER_ROLE_VIEWER"
	TeamUserUpdateParamsRoleTeamUserRoleContributor TeamUserUpdateParamsRole = "TEAM_USER_ROLE_CONTRIBUTOR"
)

type TeamUserUpdateResponseEnvelope struct {
	Data TeamUser `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TeamUserUpdateResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *TeamUserUpdateResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TeamUserListParams struct {
	// Maximum number of results to return. Default is 10000, maximum is 10000.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Token for pagination. Leave empty for the first request.
	PageToken param.Opt[string] `query:"pageToken,omitzero" json:"-"`
	// Filter by team ID. Required when using /v2/teams/{team_id}/users path.
	TeamID param.Opt[string] `query:"teamId,omitzero" json:"-"`
	// Filter by user ID. If not provided, returns all users in the team.
	UserID param.Opt[string] `query:"userId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TeamUserListParams]'s query parameters as `url.Values`.
func (r TeamUserListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TeamUserDeleteParams struct {
	TeamID param.Opt[string] `query:"teamId,omitzero" json:"-"`
	UserID param.Opt[string] `query:"userId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TeamUserDeleteParams]'s query parameters as `url.Values`.
func (r TeamUserDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
