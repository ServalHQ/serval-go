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
	"github.com/ServalHQ/serval-go/packages/pagination"
	"github.com/ServalHQ/serval-go/packages/param"
	"github.com/ServalHQ/serval-go/packages/respjson"
)

// GuidanceService contains methods and other services that help with interacting
// with the serval API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGuidanceService] method instead.
type GuidanceService struct {
	Options []option.RequestOption
}

// NewGuidanceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewGuidanceService(opts ...option.RequestOption) (r GuidanceService) {
	r = GuidanceService{}
	r.Options = opts
	return
}

// Create a new guidance for a team.
func (r *GuidanceService) New(ctx context.Context, body GuidanceNewParams, opts ...option.RequestOption) (res *Guidance, err error) {
	var env GuidanceNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "v2/guidances"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Get a specific guidance by ID.
func (r *GuidanceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Guidance, err error) {
	var env GuidanceGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/guidances/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Update an existing guidance.
func (r *GuidanceService) Update(ctx context.Context, id string, body GuidanceUpdateParams, opts ...option.RequestOption) (res *Guidance, err error) {
	var env GuidanceUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/guidances/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// List all guidances for a team.
func (r *GuidanceService) List(ctx context.Context, query GuidanceListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Guidance], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v2/guidances"
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

// List all guidances for a team.
func (r *GuidanceService) ListAutoPaging(ctx context.Context, query GuidanceListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Guidance] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a guidance.
func (r *GuidanceService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *GuidanceDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/guidances/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Guidance struct {
	// The ID of the guidance.
	ID string `json:"id"`
	// The content of the guidance.
	Content string `json:"content"`
	// A description of the guidance.
	Description string `json:"description"`
	// Whether there are unpublished changes to the guidance (computed by server).
	HasUnpublishedChanges bool `json:"hasUnpublishedChanges"`
	// Whether the guidance is published. Set to true to publish the guidance.
	IsPublished bool `json:"isPublished"`
	// The name of the guidance.
	Name string `json:"name"`
	// Whether this guidance should always be used (skipping LLM selection).
	ShouldAlwaysUse bool `json:"shouldAlwaysUse"`
	// The ID of the team that the guidance belongs to.
	TeamID string `json:"teamId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		Content               respjson.Field
		Description           respjson.Field
		HasUnpublishedChanges respjson.Field
		IsPublished           respjson.Field
		Name                  respjson.Field
		ShouldAlwaysUse       respjson.Field
		TeamID                respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Guidance) RawJSON() string { return r.JSON.raw }
func (r *Guidance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GuidanceDeleteResponse = any

type GuidanceNewParams struct {
	// The content of the guidance.
	Content string `json:"content,required"`
	// The name of the guidance.
	Name string `json:"name,required"`
	// The ID of the team.
	TeamID string `json:"teamId,required"`
	// Whether to publish the guidance after creation (optional).
	IsPublished param.Opt[bool] `json:"isPublished,omitzero"`
	// Whether this guidance should always be used (optional, defaults to false).
	ShouldAlwaysUse param.Opt[bool] `json:"shouldAlwaysUse,omitzero"`
	// A description of the guidance.
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r GuidanceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow GuidanceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GuidanceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GuidanceNewResponseEnvelope struct {
	// The created guidance.
	Data Guidance `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GuidanceNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *GuidanceNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GuidanceGetResponseEnvelope struct {
	// The guidance.
	Data Guidance `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GuidanceGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *GuidanceGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GuidanceUpdateParams struct {
	// Whether the guidance is published. Set to true to publish the guidance.
	IsPublished param.Opt[bool] `json:"isPublished,omitzero"`
	// Whether this guidance should always be used (optional).
	ShouldAlwaysUse param.Opt[bool] `json:"shouldAlwaysUse,omitzero"`
	// The content of the guidance.
	Content param.Opt[string] `json:"content,omitzero"`
	// A description of the guidance.
	Description param.Opt[string] `json:"description,omitzero"`
	// The name of the guidance.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r GuidanceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow GuidanceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GuidanceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GuidanceUpdateResponseEnvelope struct {
	// The updated guidance.
	Data Guidance `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GuidanceUpdateResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *GuidanceUpdateResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GuidanceListParams struct {
	// Maximum number of results to return. Default is 1000, maximum is 2000.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Token for pagination. Leave empty for the first request.
	PageToken param.Opt[string] `query:"pageToken,omitzero" json:"-"`
	// The ID of the team.
	TeamID param.Opt[string] `query:"teamId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GuidanceListParams]'s query parameters as `url.Values`.
func (r GuidanceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
