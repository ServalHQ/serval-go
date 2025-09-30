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

// AppInstanceService contains methods and other services that help with
// interacting with the serval API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppInstanceService] method instead.
type AppInstanceService struct {
	Options []option.RequestOption
}

// NewAppInstanceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAppInstanceService(opts ...option.RequestOption) (r AppInstanceService) {
	r = AppInstanceService{}
	r.Options = opts
	return
}

// Create a new app instance for a team.
func (r *AppInstanceService) New(ctx context.Context, body AppInstanceNewParams, opts ...option.RequestOption) (res *AppInstance, err error) {
	var env AppInstanceNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "v2/app-instances"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Get a specific app instance by ID.
func (r *AppInstanceService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AppInstance, err error) {
	var env AppInstanceGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/app-instances/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Update an existing app instance.
func (r *AppInstanceService) Update(ctx context.Context, id string, body AppInstanceUpdateParams, opts ...option.RequestOption) (res *AppInstance, err error) {
	var env AppInstanceUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/app-instances/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// List all app instances for a team.
func (r *AppInstanceService) List(ctx context.Context, query AppInstanceListParams, opts ...option.RequestOption) (res *[]AppInstance, err error) {
	var env AppInstanceListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "v2/app-instances"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

type AppInstance struct {
	// The ID of the app instance.
	ID string `json:"id"`
	// Whether access requests are enabled for the app instance.
	AccessRequestsEnabled bool `json:"accessRequestsEnabled"`
	// The default access policy for the app instance.
	DefaultAccessPolicyID string `json:"defaultAccessPolicyId,nullable"`
	// The instance ID of the app instance.
	InstanceID string `json:"instanceId"`
	// The name of the app instance.
	Name string `json:"name"`
	// The service of the app instance.
	Service string `json:"service"`
	// The ID of the Serval team that the app instance belongs to.
	TeamID string `json:"teamId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		AccessRequestsEnabled respjson.Field
		DefaultAccessPolicyID respjson.Field
		InstanceID            respjson.Field
		Name                  respjson.Field
		Service               respjson.Field
		TeamID                respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppInstance) RawJSON() string { return r.JSON.raw }
func (r *AppInstance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppInstanceNewParams struct {
	// The default access policy for the app instance (optional).
	DefaultAccessPolicyID param.Opt[string] `json:"defaultAccessPolicyId,omitzero"`
	// Whether access requests are enabled for the app instance.
	AccessRequestsEnabled param.Opt[bool] `json:"accessRequestsEnabled,omitzero"`
	// The instance ID of the app instance.
	InstanceID param.Opt[string] `json:"instanceId,omitzero"`
	// The name of the app instance.
	Name param.Opt[string] `json:"name,omitzero"`
	// The service of the app instance.
	Service param.Opt[string] `json:"service,omitzero"`
	// The ID of the team.
	TeamID param.Opt[string] `json:"teamId,omitzero"`
	paramObj
}

func (r AppInstanceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AppInstanceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppInstanceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppInstanceNewResponseEnvelope struct {
	// The created app instance.
	Data AppInstance `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppInstanceNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AppInstanceNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppInstanceGetResponseEnvelope struct {
	// The app instance.
	Data AppInstance `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppInstanceGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AppInstanceGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppInstanceUpdateParams struct {
	// The default access policy for the app instance (optional).
	DefaultAccessPolicyID param.Opt[string] `json:"defaultAccessPolicyId,omitzero"`
	// Whether access requests are enabled for the app instance.
	AccessRequestsEnabled param.Opt[bool] `json:"accessRequestsEnabled,omitzero"`
	// The instance ID of the app instance.
	InstanceID param.Opt[string] `json:"instanceId,omitzero"`
	// The name of the app instance.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r AppInstanceUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AppInstanceUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppInstanceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppInstanceUpdateResponseEnvelope struct {
	// The updated app instance.
	Data AppInstance `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppInstanceUpdateResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AppInstanceUpdateResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppInstanceListParams struct {
	// The ID of the team.
	TeamID param.Opt[string] `query:"teamId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AppInstanceListParams]'s query parameters as `url.Values`.
func (r AppInstanceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AppInstanceListResponseEnvelope struct {
	// The list of app instances.
	Data []AppInstance `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppInstanceListResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AppInstanceListResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
