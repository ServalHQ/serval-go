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

// AppResourceEntitlementService contains methods and other services that help with
// interacting with the serval API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppResourceEntitlementService] method instead.
type AppResourceEntitlementService struct {
	Options []option.RequestOption
}

// NewAppResourceEntitlementService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAppResourceEntitlementService(opts ...option.RequestOption) (r AppResourceEntitlementService) {
	r = AppResourceEntitlementService{}
	r.Options = opts
	return
}

// Create a new app resource entitlement for a resource.
func (r *AppResourceEntitlementService) New(ctx context.Context, body AppResourceEntitlementNewParams, opts ...option.RequestOption) (res *AppResourceEntitlement, err error) {
	var env AppResourceEntitlementNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "v2/app-resource-entitlements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Get a specific app resource entitlement by ID.
func (r *AppResourceEntitlementService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AppResourceEntitlement, err error) {
	var env AppResourceEntitlementGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/app-resource-entitlements/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Update an existing app resource entitlement.
func (r *AppResourceEntitlementService) Update(ctx context.Context, id string, body AppResourceEntitlementUpdateParams, opts ...option.RequestOption) (res *AppResourceEntitlement, err error) {
	var env AppResourceEntitlementUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/app-resource-entitlements/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// List all app resource entitlements for a resource.
func (r *AppResourceEntitlementService) List(ctx context.Context, query AppResourceEntitlementListParams, opts ...option.RequestOption) (res *[]AppResourceEntitlement, err error) {
	var env AppResourceEntitlementListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "v2/app-resource-entitlements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

type AppResourceEntitlement struct {
	// The ID of the entitlement.
	ID string `json:"id"`
	// The default access policy for the entitlement.
	AccessPolicyID string `json:"accessPolicyId,nullable"`
	// A description of the entitlement.
	Description string `json:"description"`
	// The name of the entitlement.
	Name string `json:"name"`
	// The provisioning method for the entitlement.
	ProvisioningMethod string `json:"provisioningMethod"`
	// Whether requests are enabled for the entitlement.
	RequestsEnabled bool `json:"requestsEnabled"`
	// The ID of the resource that the entitlement belongs to.
	ResourceID string `json:"resourceId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AccessPolicyID     respjson.Field
		Description        respjson.Field
		Name               respjson.Field
		ProvisioningMethod respjson.Field
		RequestsEnabled    respjson.Field
		ResourceID         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceEntitlement) RawJSON() string { return r.JSON.raw }
func (r *AppResourceEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceEntitlementNewParams struct {
	// The default access policy for the entitlement (optional).
	AccessPolicyID param.Opt[string] `json:"accessPolicyId,omitzero"`
	// A description of the entitlement.
	Description param.Opt[string] `json:"description,omitzero"`
	// The name of the entitlement.
	Name param.Opt[string] `json:"name,omitzero"`
	// The provisioning method for the entitlement.
	ProvisioningMethod param.Opt[string] `json:"provisioningMethod,omitzero"`
	// Whether requests are enabled for the entitlement.
	RequestsEnabled param.Opt[bool] `json:"requestsEnabled,omitzero"`
	// The ID of the resource.
	ResourceID param.Opt[string] `json:"resourceId,omitzero"`
	paramObj
}

func (r AppResourceEntitlementNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceEntitlementNewResponseEnvelope struct {
	// The created entitlement.
	Data AppResourceEntitlement `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceEntitlementNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AppResourceEntitlementNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceEntitlementGetResponseEnvelope struct {
	// The entitlement.
	Data AppResourceEntitlement `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceEntitlementGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AppResourceEntitlementGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceEntitlementUpdateParams struct {
	// The default access policy for the entitlement (optional).
	AccessPolicyID param.Opt[string] `json:"accessPolicyId,omitzero"`
	// A description of the entitlement.
	Description param.Opt[string] `json:"description,omitzero"`
	// The name of the entitlement.
	Name param.Opt[string] `json:"name,omitzero"`
	// The provisioning method for the entitlement.
	ProvisioningMethod param.Opt[string] `json:"provisioningMethod,omitzero"`
	// Whether requests are enabled for the entitlement.
	RequestsEnabled param.Opt[bool] `json:"requestsEnabled,omitzero"`
	paramObj
}

func (r AppResourceEntitlementUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceEntitlementUpdateResponseEnvelope struct {
	// The updated entitlement.
	Data AppResourceEntitlement `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceEntitlementUpdateResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AppResourceEntitlementUpdateResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceEntitlementListParams struct {
	// The ID of the resource.
	ResourceID param.Opt[string] `query:"resourceId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AppResourceEntitlementListParams]'s query parameters as
// `url.Values`.
func (r AppResourceEntitlementListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AppResourceEntitlementListResponseEnvelope struct {
	// The list of entitlements.
	Data []AppResourceEntitlement `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceEntitlementListResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AppResourceEntitlementListResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
