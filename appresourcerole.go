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

// AppResourceRoleService contains methods and other services that help with
// interacting with the serval API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppResourceRoleService] method instead.
type AppResourceRoleService struct {
	Options []option.RequestOption
}

// NewAppResourceRoleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAppResourceRoleService(opts ...option.RequestOption) (r AppResourceRoleService) {
	r = AppResourceRoleService{}
	r.Options = opts
	return
}

// Create a new app resource role for a resource.
func (r *AppResourceRoleService) New(ctx context.Context, body AppResourceRoleNewParams, opts ...option.RequestOption) (res *AppResourceRole, err error) {
	var env AppResourceRoleNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "v2/app-resource-roles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Get a specific app resource role by ID.
func (r *AppResourceRoleService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *AppResourceRole, err error) {
	var env AppResourceRoleGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/app-resource-roles/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// Update an existing app resource role.
func (r *AppResourceRoleService) Update(ctx context.Context, id string, body AppResourceRoleUpdateParams, opts ...option.RequestOption) (res *AppResourceRole, err error) {
	var env AppResourceRoleUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/app-resource-roles/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Data
	return
}

// List all app resource roles for a team, optionally filtered by resource.
func (r *AppResourceRoleService) List(ctx context.Context, query AppResourceRoleListParams, opts ...option.RequestOption) (res *AppResourceRoleListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/app-resource-roles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete an app resource role.
func (r *AppResourceRoleService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *AppResourceRoleDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/app-resource-roles/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Role represents a requestable role on an app resource.
type AppResourceRole struct {
	// The ID of the role.
	ID string `json:"id"`
	// The default access policy for the role.
	AccessPolicyID string `json:"accessPolicyId,nullable"`
	// A description of the role.
	Description string `json:"description"`
	// Data from the external system as a JSON string (optional).
	ExternalData string `json:"externalData,nullable"`
	// The external ID of the role in the external system (optional).
	ExternalID string `json:"externalId,nullable"`
	// The name of the role.
	Name string `json:"name"`
	// Provisioning configuration. **Exactly one method should be set.**
	ProvisioningMethod AppResourceRoleProvisioningMethod `json:"provisioningMethod"`
	// Whether requests are enabled for the role.
	RequestsEnabled bool `json:"requestsEnabled"`
	// The ID of the resource that the role belongs to.
	ResourceID string `json:"resourceId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AccessPolicyID     respjson.Field
		Description        respjson.Field
		ExternalData       respjson.Field
		ExternalID         respjson.Field
		Name               respjson.Field
		ProvisioningMethod respjson.Field
		RequestsEnabled    respjson.Field
		ResourceID         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceRole) RawJSON() string { return r.JSON.raw }
func (r *AppResourceRole) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provisioning configuration. **Exactly one method should be set.**
type AppResourceRoleProvisioningMethod struct {
	BuiltinWorkflow any                                             `json:"builtinWorkflow"`
	CustomWorkflow  AppResourceRoleProvisioningMethodCustomWorkflow `json:"customWorkflow"`
	LinkedRoles     AppResourceRoleProvisioningMethodLinkedRoles    `json:"linkedRoles"`
	Manual          AppResourceRoleProvisioningMethodManual         `json:"manual"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BuiltinWorkflow respjson.Field
		CustomWorkflow  respjson.Field
		LinkedRoles     respjson.Field
		Manual          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceRoleProvisioningMethod) RawJSON() string { return r.JSON.raw }
func (r *AppResourceRoleProvisioningMethod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceRoleProvisioningMethodCustomWorkflow struct {
	// The workflow ID to deprovision access.
	DeprovisionWorkflowID string `json:"deprovisionWorkflowId"`
	// The workflow ID to provision access.
	ProvisionWorkflowID string `json:"provisionWorkflowId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeprovisionWorkflowID respjson.Field
		ProvisionWorkflowID   respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceRoleProvisioningMethodCustomWorkflow) RawJSON() string { return r.JSON.raw }
func (r *AppResourceRoleProvisioningMethodCustomWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceRoleProvisioningMethodLinkedRoles struct {
	// The IDs of prerequisite roles.
	LinkedRoleIDs []string `json:"linkedRoleIds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LinkedRoleIDs respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceRoleProvisioningMethodLinkedRoles) RawJSON() string { return r.JSON.raw }
func (r *AppResourceRoleProvisioningMethodLinkedRoles) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceRoleProvisioningMethodManual struct {
	// Users and groups that should be assigned/notified for manual provisioning.
	Assignees []AppResourceRoleProvisioningMethodManualAssignee `json:"assignees"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Assignees   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceRoleProvisioningMethodManual) RawJSON() string { return r.JSON.raw }
func (r *AppResourceRoleProvisioningMethodManual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceRoleProvisioningMethodManualAssignee struct {
	// The ID of the user or group.
	AssigneeID string `json:"assigneeId"`
	// The type of assignee.
	//
	// Any of "MANUAL_PROVISIONING_ASSIGNEE_TYPE_UNSPECIFIED",
	// "MANUAL_PROVISIONING_ASSIGNEE_TYPE_USER",
	// "MANUAL_PROVISIONING_ASSIGNEE_TYPE_GROUP".
	AssigneeType string `json:"assigneeType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssigneeID   respjson.Field
		AssigneeType respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceRoleProvisioningMethodManualAssignee) RawJSON() string { return r.JSON.raw }
func (r *AppResourceRoleProvisioningMethodManualAssignee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceRoleListResponse struct {
	// The list of roles.
	Data []AppResourceRole `json:"data"`
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
func (r AppResourceRoleListResponse) RawJSON() string { return r.JSON.raw }
func (r *AppResourceRoleListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceRoleDeleteResponse = any

type AppResourceRoleNewParams struct {
	// The default access policy for the role (optional).
	AccessPolicyID param.Opt[string] `json:"accessPolicyId,omitzero"`
	// Data from the external system as a JSON string (optional).
	ExternalData param.Opt[string] `json:"externalData,omitzero"`
	// The external ID of the role in the external system (optional).
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// A description of the role.
	Description param.Opt[string] `json:"description,omitzero"`
	// The name of the role.
	Name param.Opt[string] `json:"name,omitzero"`
	// Whether requests are enabled for the role.
	RequestsEnabled param.Opt[bool] `json:"requestsEnabled,omitzero"`
	// The ID of the resource.
	ResourceID param.Opt[string] `json:"resourceId,omitzero"`
	// Provisioning configuration. Exactly one method should be set.
	ProvisioningMethod AppResourceRoleNewParamsProvisioningMethod `json:"provisioningMethod,omitzero"`
	paramObj
}

func (r AppResourceRoleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provisioning configuration. Exactly one method should be set.
type AppResourceRoleNewParamsProvisioningMethod struct {
	BuiltinWorkflow any                                                      `json:"builtinWorkflow,omitzero"`
	CustomWorkflow  AppResourceRoleNewParamsProvisioningMethodCustomWorkflow `json:"customWorkflow,omitzero"`
	LinkedRoles     AppResourceRoleNewParamsProvisioningMethodLinkedRoles    `json:"linkedRoles,omitzero"`
	Manual          AppResourceRoleNewParamsProvisioningMethodManual         `json:"manual,omitzero"`
	paramObj
}

func (r AppResourceRoleNewParamsProvisioningMethod) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleNewParamsProvisioningMethod
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleNewParamsProvisioningMethod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceRoleNewParamsProvisioningMethodCustomWorkflow struct {
	// The workflow ID to deprovision access.
	DeprovisionWorkflowID param.Opt[string] `json:"deprovisionWorkflowId,omitzero"`
	// The workflow ID to provision access.
	ProvisionWorkflowID param.Opt[string] `json:"provisionWorkflowId,omitzero"`
	paramObj
}

func (r AppResourceRoleNewParamsProvisioningMethodCustomWorkflow) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleNewParamsProvisioningMethodCustomWorkflow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleNewParamsProvisioningMethodCustomWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceRoleNewParamsProvisioningMethodLinkedRoles struct {
	// The IDs of prerequisite roles.
	LinkedRoleIDs []string `json:"linkedRoleIds,omitzero"`
	paramObj
}

func (r AppResourceRoleNewParamsProvisioningMethodLinkedRoles) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleNewParamsProvisioningMethodLinkedRoles
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleNewParamsProvisioningMethodLinkedRoles) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceRoleNewParamsProvisioningMethodManual struct {
	// Users and groups that should be assigned/notified for manual provisioning.
	Assignees []AppResourceRoleNewParamsProvisioningMethodManualAssignee `json:"assignees,omitzero"`
	paramObj
}

func (r AppResourceRoleNewParamsProvisioningMethodManual) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleNewParamsProvisioningMethodManual
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleNewParamsProvisioningMethodManual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceRoleNewParamsProvisioningMethodManualAssignee struct {
	// The ID of the user or group.
	AssigneeID param.Opt[string] `json:"assigneeId,omitzero"`
	// The type of assignee.
	//
	// Any of "MANUAL_PROVISIONING_ASSIGNEE_TYPE_UNSPECIFIED",
	// "MANUAL_PROVISIONING_ASSIGNEE_TYPE_USER",
	// "MANUAL_PROVISIONING_ASSIGNEE_TYPE_GROUP".
	AssigneeType string `json:"assigneeType,omitzero"`
	paramObj
}

func (r AppResourceRoleNewParamsProvisioningMethodManualAssignee) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleNewParamsProvisioningMethodManualAssignee
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleNewParamsProvisioningMethodManualAssignee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AppResourceRoleNewParamsProvisioningMethodManualAssignee](
		"assigneeType", "MANUAL_PROVISIONING_ASSIGNEE_TYPE_UNSPECIFIED", "MANUAL_PROVISIONING_ASSIGNEE_TYPE_USER", "MANUAL_PROVISIONING_ASSIGNEE_TYPE_GROUP",
	)
}

type AppResourceRoleNewResponseEnvelope struct {
	// The created role.
	Data AppResourceRole `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceRoleNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AppResourceRoleNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceRoleGetResponseEnvelope struct {
	// The role.
	Data AppResourceRole `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceRoleGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AppResourceRoleGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceRoleUpdateParams struct {
	// The default access policy for the role (optional).
	AccessPolicyID param.Opt[string] `json:"accessPolicyId,omitzero"`
	// Data from the external system as a JSON string (optional).
	ExternalData param.Opt[string] `json:"externalData,omitzero"`
	// The external ID of the role in the external system (optional).
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// A description of the role.
	Description param.Opt[string] `json:"description,omitzero"`
	// The name of the role.
	Name param.Opt[string] `json:"name,omitzero"`
	// Whether requests are enabled for the role.
	RequestsEnabled param.Opt[bool] `json:"requestsEnabled,omitzero"`
	// Provisioning configuration. Exactly one method should be set.
	ProvisioningMethod AppResourceRoleUpdateParamsProvisioningMethod `json:"provisioningMethod,omitzero"`
	paramObj
}

func (r AppResourceRoleUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provisioning configuration. Exactly one method should be set.
type AppResourceRoleUpdateParamsProvisioningMethod struct {
	BuiltinWorkflow any                                                         `json:"builtinWorkflow,omitzero"`
	CustomWorkflow  AppResourceRoleUpdateParamsProvisioningMethodCustomWorkflow `json:"customWorkflow,omitzero"`
	LinkedRoles     AppResourceRoleUpdateParamsProvisioningMethodLinkedRoles    `json:"linkedRoles,omitzero"`
	Manual          AppResourceRoleUpdateParamsProvisioningMethodManual         `json:"manual,omitzero"`
	paramObj
}

func (r AppResourceRoleUpdateParamsProvisioningMethod) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleUpdateParamsProvisioningMethod
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleUpdateParamsProvisioningMethod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceRoleUpdateParamsProvisioningMethodCustomWorkflow struct {
	// The workflow ID to deprovision access.
	DeprovisionWorkflowID param.Opt[string] `json:"deprovisionWorkflowId,omitzero"`
	// The workflow ID to provision access.
	ProvisionWorkflowID param.Opt[string] `json:"provisionWorkflowId,omitzero"`
	paramObj
}

func (r AppResourceRoleUpdateParamsProvisioningMethodCustomWorkflow) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleUpdateParamsProvisioningMethodCustomWorkflow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleUpdateParamsProvisioningMethodCustomWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceRoleUpdateParamsProvisioningMethodLinkedRoles struct {
	// The IDs of prerequisite roles.
	LinkedRoleIDs []string `json:"linkedRoleIds,omitzero"`
	paramObj
}

func (r AppResourceRoleUpdateParamsProvisioningMethodLinkedRoles) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleUpdateParamsProvisioningMethodLinkedRoles
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleUpdateParamsProvisioningMethodLinkedRoles) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceRoleUpdateParamsProvisioningMethodManual struct {
	// Users and groups that should be assigned/notified for manual provisioning.
	Assignees []AppResourceRoleUpdateParamsProvisioningMethodManualAssignee `json:"assignees,omitzero"`
	paramObj
}

func (r AppResourceRoleUpdateParamsProvisioningMethodManual) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleUpdateParamsProvisioningMethodManual
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleUpdateParamsProvisioningMethodManual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceRoleUpdateParamsProvisioningMethodManualAssignee struct {
	// The ID of the user or group.
	AssigneeID param.Opt[string] `json:"assigneeId,omitzero"`
	// The type of assignee.
	//
	// Any of "MANUAL_PROVISIONING_ASSIGNEE_TYPE_UNSPECIFIED",
	// "MANUAL_PROVISIONING_ASSIGNEE_TYPE_USER",
	// "MANUAL_PROVISIONING_ASSIGNEE_TYPE_GROUP".
	AssigneeType string `json:"assigneeType,omitzero"`
	paramObj
}

func (r AppResourceRoleUpdateParamsProvisioningMethodManualAssignee) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleUpdateParamsProvisioningMethodManualAssignee
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleUpdateParamsProvisioningMethodManualAssignee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AppResourceRoleUpdateParamsProvisioningMethodManualAssignee](
		"assigneeType", "MANUAL_PROVISIONING_ASSIGNEE_TYPE_UNSPECIFIED", "MANUAL_PROVISIONING_ASSIGNEE_TYPE_USER", "MANUAL_PROVISIONING_ASSIGNEE_TYPE_GROUP",
	)
}

type AppResourceRoleUpdateResponseEnvelope struct {
	// The updated role.
	Data AppResourceRole `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceRoleUpdateResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *AppResourceRoleUpdateResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceRoleListParams struct {
	// Filter by app instance ID. At least one of team_id, app_instance_id, or
	// resource_id must be provided.
	AppInstanceID param.Opt[string] `query:"appInstanceId,omitzero" json:"-"`
	// Maximum number of results to return. Default is 5000, maximum is 5000.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Token for pagination. Leave empty for the first request.
	PageToken param.Opt[string] `query:"pageToken,omitzero" json:"-"`
	// Filter by resource ID. At least one of team_id, app_instance_id, or resource_id
	// must be provided.
	ResourceID param.Opt[string] `query:"resourceId,omitzero" json:"-"`
	// Filter by team ID. At least one of team_id, app_instance_id, or resource_id must
	// be provided.
	TeamID param.Opt[string] `query:"teamId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AppResourceRoleListParams]'s query parameters as
// `url.Values`.
func (r AppResourceRoleListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
