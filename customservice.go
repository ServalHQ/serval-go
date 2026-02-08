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

// CustomServiceService contains methods and other services that help with
// interacting with the serval API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomServiceService] method instead.
type CustomServiceService struct {
	Options []option.RequestOption
}

// NewCustomServiceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomServiceService(opts ...option.RequestOption) (r CustomServiceService) {
	r = CustomServiceService{}
	r.Options = opts
	return
}

// Create a new custom service for a team. Custom services represent external APIs
// or systems that are not built-in integrations.
func (r *CustomServiceService) New(ctx context.Context, body CustomServiceNewParams, opts ...option.RequestOption) (res *CustomService, err error) {
	var env CustomServiceNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "v2/custom-services"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Get a specific custom service by ID.
func (r *CustomServiceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *CustomService, err error) {
	var env CustomServiceGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/custom-services/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Update an existing custom service.
func (r *CustomServiceService) Update(ctx context.Context, id string, body CustomServiceUpdateParams, opts ...option.RequestOption) (res *CustomService, err error) {
	var env CustomServiceUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/custom-services/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// List all custom services for a team.
func (r *CustomServiceService) List(ctx context.Context, query CustomServiceListParams, opts ...option.RequestOption) (res *CustomServiceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/custom-services"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a custom service. This will not delete any app instances that use this
// custom service.
func (r *CustomServiceService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *CustomServiceDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/custom-services/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type CustomService struct {
	// The ID of the custom service.
	ID string `json:"id"`
	// The domain for branding/logo lookup (e.g., "hr.company.com").
	Domain string `json:"domain"`
	// The name of the custom service (e.g., "Internal HR System").
	Name string `json:"name"`
	// The ID of the team that the custom service belongs to.
	TeamID string `json:"teamId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Domain      respjson.Field
		Name        respjson.Field
		TeamID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomService) RawJSON() string { return r.JSON.raw }
func (r *CustomService) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomServiceListResponse struct {
	// The list of custom services.
	Data []CustomService `json:"data"`
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
func (r CustomServiceListResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomServiceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomServiceDeleteResponse = any

type CustomServiceNewParams struct {
	// The name of the custom service (e.g., "Internal HR System").
	Name string `json:"name,required"`
	// The ID of the team.
	TeamID string `json:"teamId,required"`
	// The domain for branding/logo lookup (e.g., "hr.company.com").
	Domain param.Opt[string] `json:"domain,omitzero"`
	paramObj
}

func (r CustomServiceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomServiceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomServiceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomServiceNewResponseEnvelope struct {
	// The created custom service.
	Data CustomService `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomServiceNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *CustomServiceNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomServiceGetResponseEnvelope struct {
	// The custom service.
	Data CustomService `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomServiceGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *CustomServiceGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomServiceUpdateParams struct {
	// The domain for branding/logo lookup.
	Domain param.Opt[string] `json:"domain,omitzero"`
	// The name of the custom service.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r CustomServiceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomServiceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomServiceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomServiceUpdateResponseEnvelope struct {
	// The updated custom service.
	Data CustomService `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomServiceUpdateResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *CustomServiceUpdateResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomServiceListParams struct {
	// Maximum number of results to return. Default is 10000, maximum is 10000.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Token for pagination. Leave empty for the first request.
	PageToken param.Opt[string] `query:"pageToken,omitzero" json:"-"`
	// The ID of the team.
	TeamID param.Opt[string] `query:"teamId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CustomServiceListParams]'s query parameters as
// `url.Values`.
func (r CustomServiceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
