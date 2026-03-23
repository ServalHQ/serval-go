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

// TagService contains methods and other services that help with interacting with
// the serval API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTagService] method instead.
type TagService struct {
	Options []option.RequestOption
}

// NewTagService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTagService(opts ...option.RequestOption) (r TagService) {
	r = TagService{}
	r.Options = opts
	return
}

// Create a new tag.
func (r *TagService) New(ctx context.Context, body TagNewParams, opts ...option.RequestOption) (res *Tag, err error) {
	var env TagNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "v2/tags"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// Get a specific tag by ID.
func (r *TagService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Tag, err error) {
	var env TagGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/tags/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// Update an existing tag.
func (r *TagService) Update(ctx context.Context, id string, body TagUpdateParams, opts ...option.RequestOption) (res *Tag, err error) {
	var env TagUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/tags/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// List all tags for a team.
func (r *TagService) List(ctx context.Context, query TagListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Tag], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v2/tags"
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

// List all tags for a team.
func (r *TagService) ListAutoPaging(ctx context.Context, query TagListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Tag] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a tag.
func (r *TagService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *TagDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/tags/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// A tag that can be associated with workflows.
type Tag struct {
	// The ID of the tag.
	ID string `json:"id"`
	// The color of the tag (CSS color string).
	Color string `json:"color" api:"nullable"`
	// The icon slug for the tag.
	IconSlug string `json:"iconSlug" api:"nullable"`
	// The name of the tag.
	Name string `json:"name"`
	// The ID of the team that owns this tag.
	TeamID string `json:"teamId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Color       respjson.Field
		IconSlug    respjson.Field
		Name        respjson.Field
		TeamID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Tag) RawJSON() string { return r.JSON.raw }
func (r *Tag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagDeleteResponse = any

type TagNewParams struct {
	// The name of the tag.
	Name string `json:"name" api:"required"`
	// The ID of the team to create the tag for.
	TeamID string `json:"teamId" api:"required"`
	// The color of the tag (CSS color string).
	Color param.Opt[string] `json:"color,omitzero"`
	// The icon slug for the tag.
	IconSlug param.Opt[string] `json:"iconSlug,omitzero"`
	paramObj
}

func (r TagNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TagNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagNewResponseEnvelope struct {
	// The created tag.
	Data Tag `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *TagNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagGetResponseEnvelope struct {
	// The tag.
	Data Tag `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *TagGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagUpdateParams struct {
	// The color of the tag (CSS color string).
	Color param.Opt[string] `json:"color,omitzero"`
	// The icon slug for the tag.
	IconSlug param.Opt[string] `json:"iconSlug,omitzero"`
	// The name of the tag.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r TagUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TagUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TagUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagUpdateResponseEnvelope struct {
	// The updated tag.
	Data Tag `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TagUpdateResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *TagUpdateResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TagListParams struct {
	// Maximum number of results to return. Default is 10000, maximum is 10000.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Token for pagination. Leave empty for the first request.
	PageToken param.Opt[string] `query:"pageToken,omitzero" json:"-"`
	// The ID of the team.
	TeamID param.Opt[string] `query:"teamId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TagListParams]'s query parameters as `url.Values`.
func (r TagListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
