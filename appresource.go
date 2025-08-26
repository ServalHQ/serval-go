// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package serval

import (
	"context"
	"net/http"
	"net/url"

	"github.com/stainless-sdks/serval-go/internal/apijson"
	"github.com/stainless-sdks/serval-go/internal/apiquery"
	"github.com/stainless-sdks/serval-go/internal/requestconfig"
	"github.com/stainless-sdks/serval-go/option"
	"github.com/stainless-sdks/serval-go/packages/param"
	"github.com/stainless-sdks/serval-go/packages/respjson"
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

// List all app resources for an app instance.
func (r *AppResourceService) List(ctx context.Context, query AppResourceListParams, opts ...option.RequestOption) (res *[]AppResource, err error) {
	var env AppResourceListResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "v2/app-resources"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
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

type AppResourceListParams struct {
	// The ID of the app instance.
	AppInstanceID param.Opt[string] `query:"appInstanceId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AppResourceListParams]'s query parameters as `url.Values`.
func (r AppResourceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AppResourceListResponseEnvelope struct {
	// The list of resources.
	Data []AppResource `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceListResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AppResourceListResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
