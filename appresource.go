// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package serval

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/ServalHQ/serval-go/internal/apijson"
	"github.com/ServalHQ/serval-go/internal/requestconfig"
	"github.com/ServalHQ/serval-go/option"
	"github.com/ServalHQ/serval-go/packages/param"
	"github.com/ServalHQ/serval-go/packages/respjson"
)

// AppResourceService contains methods and other services that help with
// interacting with the serval API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppResourceService] method instead.
type AppResourceService struct {
	Options []option.RequestOption
}

// NewAppResourceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAppResourceService(opts ...option.RequestOption) (r AppResourceService) {
	r = AppResourceService{}
	r.Options = opts
	return
}

// Create a new app resource for an app instance.
func (r *AppResourceService) New(ctx context.Context, body AppResourceNewParams, opts ...option.RequestOption) (res *AppResource, err error) {
	var env AppResourceNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "v2/app-resources"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Get a specific app resource by ID.
func (r *AppResourceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AppResource, err error) {
	var env AppResourceGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/app-resources/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Update an existing app resource.
func (r *AppResourceService) Update(ctx context.Context, id string, body AppResourceUpdateParams, opts ...option.RequestOption) (res *AppResource, err error) {
	var env AppResourceUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/app-resources/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Delete an app resource.
func (r *AppResourceService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *AppResourceDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/app-resources/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AppResource struct {
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
func (r AppResource) RawJSON() string { return r.JSON.raw }
func (r *AppResource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceDeleteResponse = any

type AppResourceNewParams struct {
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

func (r AppResourceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceNewResponseEnvelope struct {
	// The created resource.
	Data AppResource `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AppResourceNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceGetResponseEnvelope struct {
	// The resource.
	Data AppResource `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AppResourceGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceUpdateParams struct {
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

func (r AppResourceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceUpdateResponseEnvelope struct {
	// The updated resource.
	Data AppResource `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceUpdateResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AppResourceUpdateResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
