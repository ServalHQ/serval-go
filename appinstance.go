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
func (r *AppInstanceService) List(ctx context.Context, query AppInstanceListParams, opts ...option.RequestOption) (res *AppInstanceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/app-instances"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete an app instance.
func (r *AppInstanceService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *AppInstanceDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/app-instances/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Configuration object.
//
// **Set exactly ONE of:** customServiceId, service
type AppInstance struct {
	// The ID of the app instance.
	ID string `json:"id"`
	// Whether access requests are enabled for the app instance.
	AccessRequestsEnabled bool `json:"accessRequestsEnabled"`
	// **Option: custom_service_id** — The ID of the custom service (for custom apps).
	CustomServiceID string `json:"customServiceId"`
	// The default access policy for the app instance.
	DefaultAccessPolicyID string `json:"defaultAccessPolicyId,nullable"`
	// The external service instance ID (e.g., GitHub org name, Okta domain, AWS
	// account ID).
	ExternalServiceInstanceID string `json:"externalServiceInstanceId"`
	// The name of the app instance.
	Name string `json:"name"`
	// **Option: service** — The service identifier (for built-in services like
	// "github", "okta", "aws").
	Service string `json:"service"`
	// The ID of the Serval team that the app instance belongs to.
	TeamID string `json:"teamId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                        respjson.Field
		AccessRequestsEnabled     respjson.Field
		CustomServiceID           respjson.Field
		DefaultAccessPolicyID     respjson.Field
		ExternalServiceInstanceID respjson.Field
		Name                      respjson.Field
		Service                   respjson.Field
		TeamID                    respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppInstance) RawJSON() string { return r.JSON.raw }
func (r *AppInstance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppInstanceListResponse struct {
	// The list of app instances.
	Data []AppInstance `json:"data"`
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
func (r AppInstanceListResponse) RawJSON() string { return r.JSON.raw }
func (r *AppInstanceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppInstanceDeleteResponse = any

type AppInstanceNewParams struct {
	// The external service instance ID (e.g., GitHub org name, Okta domain, AWS
	// account ID).
	ExternalServiceInstanceID string `json:"externalServiceInstanceId,required"`
	// The name of the app instance.
	Name string `json:"name,required"`
	// The ID of the team.
	TeamID string `json:"teamId,required"`
	// The default access policy for the app instance (optional).
	DefaultAccessPolicyID param.Opt[string] `json:"defaultAccessPolicyId,omitzero"`
	// Whether access requests are enabled for the app instance.
	AccessRequestsEnabled param.Opt[bool] `json:"accessRequestsEnabled,omitzero"`
	// **Option: custom_service_id** — The ID of a custom service to create the app
	// instance for.
	CustomServiceID param.Opt[string] `json:"customServiceId,omitzero"`
	// **Option: service** — The service identifier (for built-in services like
	// "github", "okta", "aws").
	Service param.Opt[string] `json:"service,omitzero"`
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
	// The external service instance ID (e.g., GitHub org name, Okta domain, AWS
	// account ID).
	ExternalServiceInstanceID param.Opt[string] `json:"externalServiceInstanceId,omitzero"`
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
	// Maximum number of results to return. Default is 1000, maximum is 1000.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Token for pagination. Leave empty for the first request.
	PageToken param.Opt[string] `query:"pageToken,omitzero" json:"-"`
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