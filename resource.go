// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package serval

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/serval-go/internal/apijson"
	"github.com/stainless-sdks/serval-go/internal/apiquery"
	"github.com/stainless-sdks/serval-go/internal/requestconfig"
	"github.com/stainless-sdks/serval-go/option"
	"github.com/stainless-sdks/serval-go/packages/param"
	"github.com/stainless-sdks/serval-go/packages/respjson"
)

// ResourceService contains methods and other services that help with interacting
// with the serval API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewResourceService] method instead.
type ResourceService struct {
	Options []option.RequestOption
}

// NewResourceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewResourceService(opts ...option.RequestOption) (r ResourceService) {
	r = ResourceService{}
	r.Options = opts
	return
}

// Create a new resource for an app instance.
func (r *ResourceService) New(ctx context.Context, body ResourceNewParams, opts ...option.RequestOption) (res *Resource, err error) {
	var env ResourceNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "v2/resources"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Get a specific resource by ID.
func (r *ResourceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Resource, err error) {
	var env ResourceGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/resources/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Update an existing resource.
func (r *ResourceService) Update(ctx context.Context, id string, body ResourceUpdateParams, opts ...option.RequestOption) (res *Resource, err error) {
	var env ResourceUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/resources/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// List all resources for an app instance.
func (r *ResourceService) List(ctx context.Context, query ResourceListParams, opts ...option.RequestOption) (res *[]Resource, err error) {
	var env ResourceListResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "v2/resources"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Delete a resource.
func (r *ResourceService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *ResourceDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/resources/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Resource struct {
	// The ID of the resource.
	ID string `json:"id"`
	// The ID of the app instance that the resource belongs to.
	AppInstanceID string `json:"appInstanceId"`
	// A description of the resource.
	Description string `json:"description"`
	// The external ID of the resource.
	ExternalID string `json:"externalId,nullable"`
	// The name of the resource.
	Name string `json:"name"`
	// The type of the resource.
	ResourceType string `json:"resourceType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		AppInstanceID respjson.Field
		Description   respjson.Field
		ExternalID    respjson.Field
		Name          respjson.Field
		ResourceType  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Resource) RawJSON() string { return r.JSON.raw }
func (r *Resource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResourceDeleteResponse = any

type ResourceNewParams struct {
	// The external ID of the resource (optional).
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// The ID of the app instance.
	AppInstanceID param.Opt[string] `json:"appInstanceId,omitzero"`
	// A description of the resource.
	Description param.Opt[string] `json:"description,omitzero"`
	// The name of the resource.
	Name param.Opt[string] `json:"name,omitzero"`
	// The type of the resource.
	ResourceType param.Opt[string] `json:"resourceType,omitzero"`
	paramObj
}

func (r ResourceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ResourceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResourceNewResponseEnvelope struct {
	// The created resource.
	Data Resource `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResourceNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *ResourceNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResourceGetResponseEnvelope struct {
	// The resource.
	Data Resource `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResourceGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *ResourceGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResourceUpdateParams struct {
	// The external ID of the resource (optional).
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// A description of the resource.
	Description param.Opt[string] `json:"description,omitzero"`
	// The name of the resource.
	Name param.Opt[string] `json:"name,omitzero"`
	// The type of the resource.
	ResourceType param.Opt[string] `json:"resourceType,omitzero"`
	paramObj
}

func (r ResourceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ResourceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ResourceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResourceUpdateResponseEnvelope struct {
	// The updated resource.
	Data Resource `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResourceUpdateResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *ResourceUpdateResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResourceListParams struct {
	// The ID of the app instance.
	AppInstanceID param.Opt[string] `query:"appInstanceId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ResourceListParams]'s query parameters as `url.Values`.
func (r ResourceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ResourceListResponseEnvelope struct {
	// The list of resources.
	Data []Resource `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResourceListResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *ResourceListResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
