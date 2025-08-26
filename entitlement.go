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

// EntitlementService contains methods and other services that help with
// interacting with the serval API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntitlementService] method instead.
type EntitlementService struct {
	Options []option.RequestOption
}

// NewEntitlementService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEntitlementService(opts ...option.RequestOption) (r EntitlementService) {
	r = EntitlementService{}
	r.Options = opts
	return
}

// Create a new entitlement for a resource.
func (r *EntitlementService) New(ctx context.Context, body EntitlementNewParams, opts ...option.RequestOption) (res *Entitlement, err error) {
	var env EntitlementNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "v2/entitlements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Get a specific entitlement by ID.
func (r *EntitlementService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Entitlement, err error) {
	var env EntitlementGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/entitlements/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Update an existing entitlement.
func (r *EntitlementService) Update(ctx context.Context, id string, body EntitlementUpdateParams, opts ...option.RequestOption) (res *Entitlement, err error) {
	var env EntitlementUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/entitlements/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// List all entitlements for a resource.
func (r *EntitlementService) List(ctx context.Context, query EntitlementListParams, opts ...option.RequestOption) (res *[]Entitlement, err error) {
	var env EntitlementListResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "v2/entitlements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Delete an entitlement.
func (r *EntitlementService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *EntitlementDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/entitlements/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Entitlement struct {
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
func (r Entitlement) RawJSON() string { return r.JSON.raw }
func (r *Entitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EntitlementDeleteResponse = any

type EntitlementNewParams struct {
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

func (r EntitlementNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EntitlementNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EntitlementNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EntitlementNewResponseEnvelope struct {
	// The created entitlement.
	Data Entitlement `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntitlementNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *EntitlementNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EntitlementGetResponseEnvelope struct {
	// The entitlement.
	Data Entitlement `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntitlementGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *EntitlementGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EntitlementUpdateParams struct {
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

func (r EntitlementUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EntitlementUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EntitlementUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EntitlementUpdateResponseEnvelope struct {
	// The updated entitlement.
	Data Entitlement `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntitlementUpdateResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *EntitlementUpdateResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EntitlementListParams struct {
	// The ID of the resource.
	ResourceID param.Opt[string] `query:"resourceId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EntitlementListParams]'s query parameters as `url.Values`.
func (r EntitlementListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EntitlementListResponseEnvelope struct {
	// The list of entitlements.
	Data []Entitlement `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntitlementListResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *EntitlementListResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
