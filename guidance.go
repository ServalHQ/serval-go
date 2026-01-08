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
func (r *GuidanceService) List(ctx context.Context, query GuidanceListParams, opts ...option.RequestOption) (res *[]Guidance, err error) {
	var env GuidanceListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "v2/guidances"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
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
	// Whether there are unpublished changes to the guidance.
	HasUnpublishedChanges bool `json:"hasUnpublishedChanges"`
	// Whether the guidance has been published at least once.
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
	// The content of the guidance (optional).
	Content param.Opt[string] `json:"content,omitzero"`
	// Whether this guidance should always be used (optional, defaults to false).
	ShouldAlwaysUse param.Opt[bool] `json:"shouldAlwaysUse,omitzero"`
	// A description of the guidance.
	Description param.Opt[string] `json:"description,omitzero"`
	// The name of the guidance.
	Name param.Opt[string] `json:"name,omitzero"`
	// The ID of the team.
	TeamID param.Opt[string] `json:"teamId,omitzero"`
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

type GuidanceListResponseEnvelope struct {
	// The list of guidances.
	Data []Guidance `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GuidanceListResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *GuidanceListResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
