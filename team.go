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

// TeamService contains methods and other services that help with interacting with
// the serval API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTeamService] method instead.
type TeamService struct {
	Options []option.RequestOption
	Users   TeamUserService
}

// NewTeamService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTeamService(opts ...option.RequestOption) (r TeamService) {
	r = TeamService{}
	r.Options = opts
	r.Users = NewTeamUserService(opts...)
	return
}

// Create a new team.
func (r *TeamService) New(ctx context.Context, body TeamNewParams, opts ...option.RequestOption) (res *Team, err error) {
	var env TeamNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "v2/teams"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Get a specific team by ID.
func (r *TeamService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Team, err error) {
	var env TeamGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/teams/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Update an existing team.
func (r *TeamService) Update(ctx context.Context, id string, body TeamUpdateParams, opts ...option.RequestOption) (res *Team, err error) {
	var env TeamUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/teams/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// List all teams.
func (r *TeamService) List(ctx context.Context, query TeamListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Team], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v2/teams"
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

// List all teams.
func (r *TeamService) ListAutoPaging(ctx context.Context, query TeamListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Team] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a team.
func (r *TeamService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *TeamDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/teams/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Team struct {
	ID string `json:"id"`
	// A timestamp in RFC 3339 format (e.g., "2017-01-15T01:30:15.01Z").
	CreatedAt      time.Time `json:"createdAt" format:"date-time"`
	Description    string    `json:"description"`
	Name           string    `json:"name"`
	OrganizationID string    `json:"organizationId"`
	Prefix         string    `json:"prefix"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		Description    respjson.Field
		Name           respjson.Field
		OrganizationID respjson.Field
		Prefix         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Team) RawJSON() string { return r.JSON.raw }
func (r *Team) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TeamDeleteResponse = any

type TeamNewParams struct {
	Name        string            `json:"name,required"`
	Description param.Opt[string] `json:"description,omitzero"`
	Prefix      param.Opt[string] `json:"prefix,omitzero"`
	paramObj
}

func (r TeamNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TeamNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TeamNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TeamNewResponseEnvelope struct {
	Data Team `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TeamNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *TeamNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TeamGetResponseEnvelope struct {
	Data Team `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TeamGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *TeamGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TeamUpdateParams struct {
	Description param.Opt[string] `json:"description,omitzero"`
	Name        param.Opt[string] `json:"name,omitzero"`
	Prefix      param.Opt[string] `json:"prefix,omitzero"`
	paramObj
}

func (r TeamUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TeamUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TeamUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TeamUpdateResponseEnvelope struct {
	Data Team `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TeamUpdateResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *TeamUpdateResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TeamListParams struct {
	// Maximum number of results to return. Default is 10000, maximum is 10000.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Token for pagination. Leave empty for the first request.
	PageToken param.Opt[string] `query:"pageToken,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TeamListParams]'s query parameters as `url.Values`.
func (r TeamListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
