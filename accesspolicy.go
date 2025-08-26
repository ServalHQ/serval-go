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

// AccessPolicyService contains methods and other services that help with
// interacting with the serval API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessPolicyService] method instead.
type AccessPolicyService struct {
	Options []option.RequestOption
}

// NewAccessPolicyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessPolicyService(opts ...option.RequestOption) (r AccessPolicyService) {
	r = AccessPolicyService{}
	r.Options = opts
	return
}

// Create a new access policy for a team.
func (r *AccessPolicyService) New(ctx context.Context, body AccessPolicyNewParams, opts ...option.RequestOption) (res *AccessPolicy, err error) {
	var env AccessPolicyNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "v2/access-policies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Get a specific access policy by ID.
func (r *AccessPolicyService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AccessPolicy, err error) {
	var env AccessPolicyGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/access-policies/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Update an existing access policy.
func (r *AccessPolicyService) Update(ctx context.Context, id string, body AccessPolicyUpdateParams, opts ...option.RequestOption) (res *AccessPolicy, err error) {
	var env AccessPolicyUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/access-policies/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// List all access policies for a team.
func (r *AccessPolicyService) List(ctx context.Context, query AccessPolicyListParams, opts ...option.RequestOption) (res *[]AccessPolicy, err error) {
	var env AccessPolicyListResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "v2/access-policies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Delete an access policy.
func (r *AccessPolicyService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *AccessPolicyDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/access-policies/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AccessPolicy struct {
	// The ID of the access policy.
	ID string `json:"id"`
	// A description of the access policy.
	Description string `json:"description"`
	// The maximum number of minutes that access can be granted for.
	MaxAccessMinutes int64 `json:"maxAccessMinutes"`
	// The name of the access policy.
	Name string `json:"name"`
	// Whether a business justification is required when requesting access.
	RequireBusinessJustification bool `json:"requireBusinessJustification"`
	// The ID of the team that the access policy belongs to.
	TeamID string `json:"teamId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                           respjson.Field
		Description                  respjson.Field
		MaxAccessMinutes             respjson.Field
		Name                         respjson.Field
		RequireBusinessJustification respjson.Field
		TeamID                       respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessPolicy) RawJSON() string { return r.JSON.raw }
func (r *AccessPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessPolicyDeleteResponse = any

type AccessPolicyNewParams struct {
	// The maximum number of minutes that access can be granted for (optional).
	MaxAccessMinutes param.Opt[int64] `json:"maxAccessMinutes,omitzero"`
	// Whether a business justification is required when requesting access (optional).
	RequireBusinessJustification param.Opt[bool] `json:"requireBusinessJustification,omitzero"`
	// A description of the access policy.
	Description param.Opt[string] `json:"description,omitzero"`
	// The name of the access policy.
	Name param.Opt[string] `json:"name,omitzero"`
	// The ID of the team.
	TeamID param.Opt[string] `json:"teamId,omitzero"`
	paramObj
}

func (r AccessPolicyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AccessPolicyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccessPolicyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessPolicyNewResponseEnvelope struct {
	// The created access policy.
	Data AccessPolicy `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessPolicyNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AccessPolicyNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessPolicyGetResponseEnvelope struct {
	// The access policy.
	Data AccessPolicy `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessPolicyGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AccessPolicyGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessPolicyUpdateParams struct {
	// A description of the access policy.
	Description param.Opt[string] `json:"description,omitzero"`
	// The maximum number of minutes that access can be granted for.
	MaxAccessMinutes param.Opt[int64] `json:"maxAccessMinutes,omitzero"`
	// The name of the access policy.
	Name param.Opt[string] `json:"name,omitzero"`
	// Whether a business justification is required when requesting access.
	RequireBusinessJustification param.Opt[bool] `json:"requireBusinessJustification,omitzero"`
	paramObj
}

func (r AccessPolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AccessPolicyUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AccessPolicyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessPolicyUpdateResponseEnvelope struct {
	// The updated access policy.
	Data AccessPolicy `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessPolicyUpdateResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AccessPolicyUpdateResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AccessPolicyListParams struct {
	// The ID of the team.
	TeamID param.Opt[string] `query:"teamId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AccessPolicyListParams]'s query parameters as `url.Values`.
func (r AccessPolicyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccessPolicyListResponseEnvelope struct {
	// The list of access policies.
	Data []AccessPolicy `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AccessPolicyListResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AccessPolicyListResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
