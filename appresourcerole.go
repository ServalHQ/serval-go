// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package serval

import (
	"context"
	"encoding/json"
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
	ProvisioningMethod AppResourceRoleProvisioningMethodUnion `json:"provisioningMethod"`
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

// AppResourceRoleProvisioningMethodUnion contains all possible properties and
// values from [AppResourceRoleProvisioningMethodBuiltinWorkflow],
// [AppResourceRoleProvisioningMethodCustomWorkflow],
// [AppResourceRoleProvisioningMethodLinkedRoles],
// [AppResourceRoleProvisioningMethodManual].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AppResourceRoleProvisioningMethodUnion struct {
	// This field is from variant [AppResourceRoleProvisioningMethodBuiltinWorkflow].
	BuiltinWorkflow any `json:"builtinWorkflow"`
	// This field is from variant [AppResourceRoleProvisioningMethodCustomWorkflow].
	CustomWorkflow AppResourceRoleProvisioningMethodCustomWorkflowCustomWorkflow `json:"customWorkflow"`
	// This field is from variant [AppResourceRoleProvisioningMethodLinkedRoles].
	LinkedRoles AppResourceRoleProvisioningMethodLinkedRolesLinkedRoles `json:"linkedRoles"`
	// This field is from variant [AppResourceRoleProvisioningMethodManual].
	Manual AppResourceRoleProvisioningMethodManualManual `json:"manual"`
	JSON   struct {
		BuiltinWorkflow respjson.Field
		CustomWorkflow  respjson.Field
		LinkedRoles     respjson.Field
		Manual          respjson.Field
		raw             string
	} `json:"-"`
}

func (u AppResourceRoleProvisioningMethodUnion) AsBuiltinWorkflow() (v AppResourceRoleProvisioningMethodBuiltinWorkflow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AppResourceRoleProvisioningMethodUnion) AsCustomWorkflow() (v AppResourceRoleProvisioningMethodCustomWorkflow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AppResourceRoleProvisioningMethodUnion) AsLinkedRoles() (v AppResourceRoleProvisioningMethodLinkedRoles) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AppResourceRoleProvisioningMethodUnion) AsManual() (v AppResourceRoleProvisioningMethodManual) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AppResourceRoleProvisioningMethodUnion) RawJSON() string { return u.JSON.raw }

func (r *AppResourceRoleProvisioningMethodUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceRoleProvisioningMethodBuiltinWorkflow struct {
	// Provisioning is handled by the service's builtin workflow integration.
	BuiltinWorkflow any `json:"builtinWorkflow,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BuiltinWorkflow respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceRoleProvisioningMethodBuiltinWorkflow) RawJSON() string { return r.JSON.raw }
func (r *AppResourceRoleProvisioningMethodBuiltinWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceRoleProvisioningMethodCustomWorkflow struct {
	// Provisioning is handled by custom workflows for provision + deprovision.
	CustomWorkflow AppResourceRoleProvisioningMethodCustomWorkflowCustomWorkflow `json:"customWorkflow,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomWorkflow respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceRoleProvisioningMethodCustomWorkflow) RawJSON() string { return r.JSON.raw }
func (r *AppResourceRoleProvisioningMethodCustomWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provisioning is handled by custom workflows for provision + deprovision.
type AppResourceRoleProvisioningMethodCustomWorkflowCustomWorkflow struct {
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
func (r AppResourceRoleProvisioningMethodCustomWorkflowCustomWorkflow) RawJSON() string {
	return r.JSON.raw
}
func (r *AppResourceRoleProvisioningMethodCustomWorkflowCustomWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceRoleProvisioningMethodLinkedRoles struct {
	// Provisioning depends on prerequisite roles being provisioned first.
	LinkedRoles AppResourceRoleProvisioningMethodLinkedRolesLinkedRoles `json:"linkedRoles,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LinkedRoles respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceRoleProvisioningMethodLinkedRoles) RawJSON() string { return r.JSON.raw }
func (r *AppResourceRoleProvisioningMethodLinkedRoles) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provisioning depends on prerequisite roles being provisioned first.
type AppResourceRoleProvisioningMethodLinkedRolesLinkedRoles struct {
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
func (r AppResourceRoleProvisioningMethodLinkedRolesLinkedRoles) RawJSON() string { return r.JSON.raw }
func (r *AppResourceRoleProvisioningMethodLinkedRolesLinkedRoles) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceRoleProvisioningMethodManual struct {
	// Provisioning is handled manually by assigned users/groups.
	Manual AppResourceRoleProvisioningMethodManualManual `json:"manual,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Manual      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceRoleProvisioningMethodManual) RawJSON() string { return r.JSON.raw }
func (r *AppResourceRoleProvisioningMethodManual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provisioning is handled manually by assigned users/groups.
type AppResourceRoleProvisioningMethodManualManual struct {
	// Users and groups that should be assigned/notified for manual provisioning.
	Assignees []AppResourceRoleProvisioningMethodManualManualAssignee `json:"assignees"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Assignees   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceRoleProvisioningMethodManualManual) RawJSON() string { return r.JSON.raw }
func (r *AppResourceRoleProvisioningMethodManualManual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceRoleProvisioningMethodManualManualAssignee struct {
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
func (r AppResourceRoleProvisioningMethodManualManualAssignee) RawJSON() string { return r.JSON.raw }
func (r *AppResourceRoleProvisioningMethodManualManualAssignee) UnmarshalJSON(data []byte) error {
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
	ProvisioningMethod AppResourceRoleNewParamsProvisioningMethodUnion `json:"provisioningMethod,omitzero"`
	paramObj
}

func (r AppResourceRoleNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AppResourceRoleNewParamsProvisioningMethodUnion struct {
	OfBuiltinWorkflow *AppResourceRoleNewParamsProvisioningMethodBuiltinWorkflow `json:",omitzero,inline"`
	OfCustomWorkflow  *AppResourceRoleNewParamsProvisioningMethodCustomWorkflow  `json:",omitzero,inline"`
	OfLinkedRoles     *AppResourceRoleNewParamsProvisioningMethodLinkedRoles     `json:",omitzero,inline"`
	OfManual          *AppResourceRoleNewParamsProvisioningMethodManual          `json:",omitzero,inline"`
	paramUnion
}

func (u AppResourceRoleNewParamsProvisioningMethodUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBuiltinWorkflow, u.OfCustomWorkflow, u.OfLinkedRoles, u.OfManual)
}
func (u *AppResourceRoleNewParamsProvisioningMethodUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AppResourceRoleNewParamsProvisioningMethodUnion) asAny() any {
	if !param.IsOmitted(u.OfBuiltinWorkflow) {
		return u.OfBuiltinWorkflow
	} else if !param.IsOmitted(u.OfCustomWorkflow) {
		return u.OfCustomWorkflow
	} else if !param.IsOmitted(u.OfLinkedRoles) {
		return u.OfLinkedRoles
	} else if !param.IsOmitted(u.OfManual) {
		return u.OfManual
	}
	return nil
}

// The property BuiltinWorkflow is required.
type AppResourceRoleNewParamsProvisioningMethodBuiltinWorkflow struct {
	// Provisioning is handled by the service's builtin workflow integration.
	BuiltinWorkflow any `json:"builtinWorkflow,omitzero,required"`
	paramObj
}

func (r AppResourceRoleNewParamsProvisioningMethodBuiltinWorkflow) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleNewParamsProvisioningMethodBuiltinWorkflow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleNewParamsProvisioningMethodBuiltinWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property CustomWorkflow is required.
type AppResourceRoleNewParamsProvisioningMethodCustomWorkflow struct {
	// Provisioning is handled by custom workflows for provision + deprovision.
	CustomWorkflow AppResourceRoleNewParamsProvisioningMethodCustomWorkflowCustomWorkflow `json:"customWorkflow,omitzero,required"`
	paramObj
}

func (r AppResourceRoleNewParamsProvisioningMethodCustomWorkflow) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleNewParamsProvisioningMethodCustomWorkflow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleNewParamsProvisioningMethodCustomWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provisioning is handled by custom workflows for provision + deprovision.
type AppResourceRoleNewParamsProvisioningMethodCustomWorkflowCustomWorkflow struct {
	// The workflow ID to deprovision access.
	DeprovisionWorkflowID param.Opt[string] `json:"deprovisionWorkflowId,omitzero"`
	// The workflow ID to provision access.
	ProvisionWorkflowID param.Opt[string] `json:"provisionWorkflowId,omitzero"`
	paramObj
}

func (r AppResourceRoleNewParamsProvisioningMethodCustomWorkflowCustomWorkflow) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleNewParamsProvisioningMethodCustomWorkflowCustomWorkflow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleNewParamsProvisioningMethodCustomWorkflowCustomWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property LinkedRoles is required.
type AppResourceRoleNewParamsProvisioningMethodLinkedRoles struct {
	// Provisioning depends on prerequisite roles being provisioned first.
	LinkedRoles AppResourceRoleNewParamsProvisioningMethodLinkedRolesLinkedRoles `json:"linkedRoles,omitzero,required"`
	paramObj
}

func (r AppResourceRoleNewParamsProvisioningMethodLinkedRoles) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleNewParamsProvisioningMethodLinkedRoles
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleNewParamsProvisioningMethodLinkedRoles) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provisioning depends on prerequisite roles being provisioned first.
type AppResourceRoleNewParamsProvisioningMethodLinkedRolesLinkedRoles struct {
	// The IDs of prerequisite roles.
	LinkedRoleIDs []string `json:"linkedRoleIds,omitzero"`
	paramObj
}

func (r AppResourceRoleNewParamsProvisioningMethodLinkedRolesLinkedRoles) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleNewParamsProvisioningMethodLinkedRolesLinkedRoles
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleNewParamsProvisioningMethodLinkedRolesLinkedRoles) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Manual is required.
type AppResourceRoleNewParamsProvisioningMethodManual struct {
	// Provisioning is handled manually by assigned users/groups.
	Manual AppResourceRoleNewParamsProvisioningMethodManualManual `json:"manual,omitzero,required"`
	paramObj
}

func (r AppResourceRoleNewParamsProvisioningMethodManual) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleNewParamsProvisioningMethodManual
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleNewParamsProvisioningMethodManual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provisioning is handled manually by assigned users/groups.
type AppResourceRoleNewParamsProvisioningMethodManualManual struct {
	// Users and groups that should be assigned/notified for manual provisioning.
	Assignees []AppResourceRoleNewParamsProvisioningMethodManualManualAssignee `json:"assignees,omitzero"`
	paramObj
}

func (r AppResourceRoleNewParamsProvisioningMethodManualManual) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleNewParamsProvisioningMethodManualManual
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleNewParamsProvisioningMethodManualManual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceRoleNewParamsProvisioningMethodManualManualAssignee struct {
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

func (r AppResourceRoleNewParamsProvisioningMethodManualManualAssignee) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleNewParamsProvisioningMethodManualManualAssignee
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleNewParamsProvisioningMethodManualManualAssignee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AppResourceRoleNewParamsProvisioningMethodManualManualAssignee](
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
	ProvisioningMethod AppResourceRoleUpdateParamsProvisioningMethodUnion `json:"provisioningMethod,omitzero"`
	paramObj
}

func (r AppResourceRoleUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AppResourceRoleUpdateParamsProvisioningMethodUnion struct {
	OfBuiltinWorkflow *AppResourceRoleUpdateParamsProvisioningMethodBuiltinWorkflow `json:",omitzero,inline"`
	OfCustomWorkflow  *AppResourceRoleUpdateParamsProvisioningMethodCustomWorkflow  `json:",omitzero,inline"`
	OfLinkedRoles     *AppResourceRoleUpdateParamsProvisioningMethodLinkedRoles     `json:",omitzero,inline"`
	OfManual          *AppResourceRoleUpdateParamsProvisioningMethodManual          `json:",omitzero,inline"`
	paramUnion
}

func (u AppResourceRoleUpdateParamsProvisioningMethodUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBuiltinWorkflow, u.OfCustomWorkflow, u.OfLinkedRoles, u.OfManual)
}
func (u *AppResourceRoleUpdateParamsProvisioningMethodUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AppResourceRoleUpdateParamsProvisioningMethodUnion) asAny() any {
	if !param.IsOmitted(u.OfBuiltinWorkflow) {
		return u.OfBuiltinWorkflow
	} else if !param.IsOmitted(u.OfCustomWorkflow) {
		return u.OfCustomWorkflow
	} else if !param.IsOmitted(u.OfLinkedRoles) {
		return u.OfLinkedRoles
	} else if !param.IsOmitted(u.OfManual) {
		return u.OfManual
	}
	return nil
}

// The property BuiltinWorkflow is required.
type AppResourceRoleUpdateParamsProvisioningMethodBuiltinWorkflow struct {
	// Provisioning is handled by the service's builtin workflow integration.
	BuiltinWorkflow any `json:"builtinWorkflow,omitzero,required"`
	paramObj
}

func (r AppResourceRoleUpdateParamsProvisioningMethodBuiltinWorkflow) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleUpdateParamsProvisioningMethodBuiltinWorkflow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleUpdateParamsProvisioningMethodBuiltinWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property CustomWorkflow is required.
type AppResourceRoleUpdateParamsProvisioningMethodCustomWorkflow struct {
	// Provisioning is handled by custom workflows for provision + deprovision.
	CustomWorkflow AppResourceRoleUpdateParamsProvisioningMethodCustomWorkflowCustomWorkflow `json:"customWorkflow,omitzero,required"`
	paramObj
}

func (r AppResourceRoleUpdateParamsProvisioningMethodCustomWorkflow) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleUpdateParamsProvisioningMethodCustomWorkflow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleUpdateParamsProvisioningMethodCustomWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provisioning is handled by custom workflows for provision + deprovision.
type AppResourceRoleUpdateParamsProvisioningMethodCustomWorkflowCustomWorkflow struct {
	// The workflow ID to deprovision access.
	DeprovisionWorkflowID param.Opt[string] `json:"deprovisionWorkflowId,omitzero"`
	// The workflow ID to provision access.
	ProvisionWorkflowID param.Opt[string] `json:"provisionWorkflowId,omitzero"`
	paramObj
}

func (r AppResourceRoleUpdateParamsProvisioningMethodCustomWorkflowCustomWorkflow) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleUpdateParamsProvisioningMethodCustomWorkflowCustomWorkflow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleUpdateParamsProvisioningMethodCustomWorkflowCustomWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property LinkedRoles is required.
type AppResourceRoleUpdateParamsProvisioningMethodLinkedRoles struct {
	// Provisioning depends on prerequisite roles being provisioned first.
	LinkedRoles AppResourceRoleUpdateParamsProvisioningMethodLinkedRolesLinkedRoles `json:"linkedRoles,omitzero,required"`
	paramObj
}

func (r AppResourceRoleUpdateParamsProvisioningMethodLinkedRoles) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleUpdateParamsProvisioningMethodLinkedRoles
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleUpdateParamsProvisioningMethodLinkedRoles) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provisioning depends on prerequisite roles being provisioned first.
type AppResourceRoleUpdateParamsProvisioningMethodLinkedRolesLinkedRoles struct {
	// The IDs of prerequisite roles.
	LinkedRoleIDs []string `json:"linkedRoleIds,omitzero"`
	paramObj
}

func (r AppResourceRoleUpdateParamsProvisioningMethodLinkedRolesLinkedRoles) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleUpdateParamsProvisioningMethodLinkedRolesLinkedRoles
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleUpdateParamsProvisioningMethodLinkedRolesLinkedRoles) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Manual is required.
type AppResourceRoleUpdateParamsProvisioningMethodManual struct {
	// Provisioning is handled manually by assigned users/groups.
	Manual AppResourceRoleUpdateParamsProvisioningMethodManualManual `json:"manual,omitzero,required"`
	paramObj
}

func (r AppResourceRoleUpdateParamsProvisioningMethodManual) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleUpdateParamsProvisioningMethodManual
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleUpdateParamsProvisioningMethodManual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provisioning is handled manually by assigned users/groups.
type AppResourceRoleUpdateParamsProvisioningMethodManualManual struct {
	// Users and groups that should be assigned/notified for manual provisioning.
	Assignees []AppResourceRoleUpdateParamsProvisioningMethodManualManualAssignee `json:"assignees,omitzero"`
	paramObj
}

func (r AppResourceRoleUpdateParamsProvisioningMethodManualManual) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleUpdateParamsProvisioningMethodManualManual
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleUpdateParamsProvisioningMethodManualManual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceRoleUpdateParamsProvisioningMethodManualManualAssignee struct {
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

func (r AppResourceRoleUpdateParamsProvisioningMethodManualManualAssignee) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceRoleUpdateParamsProvisioningMethodManualManualAssignee
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceRoleUpdateParamsProvisioningMethodManualManualAssignee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AppResourceRoleUpdateParamsProvisioningMethodManualManualAssignee](
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
