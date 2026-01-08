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

// Delete an app resource entitlement.
func (r *AppResourceEntitlementService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *AppResourceEntitlementDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/app-resource-entitlements/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
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
	// Provisioning configuration. **Exactly one method should be set.**
	ProvisioningMethod AppResourceEntitlementProvisioningMethod `json:"provisioningMethod"`
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

// Provisioning configuration. **Exactly one method should be set.**
type AppResourceEntitlementProvisioningMethod struct {
	// **Option: builtin_workflow**
	BuiltinWorkflow any `json:"builtinWorkflow"`
	// **Option: custom_workflow**
	CustomWorkflow AppResourceEntitlementProvisioningMethodCustomWorkflow `json:"customWorkflow"`
	// **Option: linked_entitlements**
	LinkedEntitlements AppResourceEntitlementProvisioningMethodLinkedEntitlements `json:"linkedEntitlements"`
	// **Option: manual**
	Manual AppResourceEntitlementProvisioningMethodManual `json:"manual"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BuiltinWorkflow    respjson.Field
		CustomWorkflow     respjson.Field
		LinkedEntitlements respjson.Field
		Manual             respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceEntitlementProvisioningMethod) RawJSON() string { return r.JSON.raw }
func (r *AppResourceEntitlementProvisioningMethod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Option: custom_workflow**
type AppResourceEntitlementProvisioningMethodCustomWorkflow struct {
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
func (r AppResourceEntitlementProvisioningMethodCustomWorkflow) RawJSON() string { return r.JSON.raw }
func (r *AppResourceEntitlementProvisioningMethodCustomWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Option: linked_entitlements**
type AppResourceEntitlementProvisioningMethodLinkedEntitlements struct {
	// The IDs of prerequisite entitlements.
	LinkedEntitlementIDs []string `json:"linkedEntitlementIds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LinkedEntitlementIDs respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceEntitlementProvisioningMethodLinkedEntitlements) RawJSON() string {
	return r.JSON.raw
}
func (r *AppResourceEntitlementProvisioningMethodLinkedEntitlements) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Option: manual**
type AppResourceEntitlementProvisioningMethodManual struct {
	// Users and groups that should be assigned/notified for manual provisioning.
	Assignees []AppResourceEntitlementProvisioningMethodManualAssignee `json:"assignees"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Assignees   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceEntitlementProvisioningMethodManual) RawJSON() string { return r.JSON.raw }
func (r *AppResourceEntitlementProvisioningMethodManual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceEntitlementProvisioningMethodManualAssignee struct {
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
func (r AppResourceEntitlementProvisioningMethodManualAssignee) RawJSON() string { return r.JSON.raw }
func (r *AppResourceEntitlementProvisioningMethodManualAssignee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceEntitlementDeleteResponse = any

type AppResourceEntitlementNewParams struct {
	// The default access policy for the entitlement (optional).
	AccessPolicyID param.Opt[string] `json:"accessPolicyId,omitzero"`
	// A description of the entitlement.
	Description param.Opt[string] `json:"description,omitzero"`
	// The name of the entitlement.
	Name param.Opt[string] `json:"name,omitzero"`
	// Whether requests are enabled for the entitlement.
	RequestsEnabled param.Opt[bool] `json:"requestsEnabled,omitzero"`
	// The ID of the resource.
	ResourceID param.Opt[string] `json:"resourceId,omitzero"`
	// Provisioning configuration. Exactly one method should be set.
	ProvisioningMethod AppResourceEntitlementNewParamsProvisioningMethod `json:"provisioningMethod,omitzero"`
	paramObj
}

func (r AppResourceEntitlementNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provisioning configuration. Exactly one method should be set.
type AppResourceEntitlementNewParamsProvisioningMethod struct {
	// **Option: builtin_workflow**
	BuiltinWorkflow any `json:"builtinWorkflow,omitzero"`
	// **Option: custom_workflow**
	CustomWorkflow AppResourceEntitlementNewParamsProvisioningMethodCustomWorkflow `json:"customWorkflow,omitzero"`
	// **Option: linked_entitlements**
	LinkedEntitlements AppResourceEntitlementNewParamsProvisioningMethodLinkedEntitlements `json:"linkedEntitlements,omitzero"`
	// **Option: manual**
	Manual AppResourceEntitlementNewParamsProvisioningMethodManual `json:"manual,omitzero"`
	paramObj
}

func (r AppResourceEntitlementNewParamsProvisioningMethod) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementNewParamsProvisioningMethod
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementNewParamsProvisioningMethod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Option: custom_workflow**
type AppResourceEntitlementNewParamsProvisioningMethodCustomWorkflow struct {
	// The workflow ID to deprovision access.
	DeprovisionWorkflowID param.Opt[string] `json:"deprovisionWorkflowId,omitzero"`
	// The workflow ID to provision access.
	ProvisionWorkflowID param.Opt[string] `json:"provisionWorkflowId,omitzero"`
	paramObj
}

func (r AppResourceEntitlementNewParamsProvisioningMethodCustomWorkflow) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementNewParamsProvisioningMethodCustomWorkflow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementNewParamsProvisioningMethodCustomWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Option: linked_entitlements**
type AppResourceEntitlementNewParamsProvisioningMethodLinkedEntitlements struct {
	// The IDs of prerequisite entitlements.
	LinkedEntitlementIDs []string `json:"linkedEntitlementIds,omitzero"`
	paramObj
}

func (r AppResourceEntitlementNewParamsProvisioningMethodLinkedEntitlements) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementNewParamsProvisioningMethodLinkedEntitlements
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementNewParamsProvisioningMethodLinkedEntitlements) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Option: manual**
type AppResourceEntitlementNewParamsProvisioningMethodManual struct {
	// Users and groups that should be assigned/notified for manual provisioning.
	Assignees []AppResourceEntitlementNewParamsProvisioningMethodManualAssignee `json:"assignees,omitzero"`
	paramObj
}

func (r AppResourceEntitlementNewParamsProvisioningMethodManual) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementNewParamsProvisioningMethodManual
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementNewParamsProvisioningMethodManual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceEntitlementNewParamsProvisioningMethodManualAssignee struct {
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

func (r AppResourceEntitlementNewParamsProvisioningMethodManualAssignee) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementNewParamsProvisioningMethodManualAssignee
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementNewParamsProvisioningMethodManualAssignee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AppResourceEntitlementNewParamsProvisioningMethodManualAssignee](
		"assigneeType", "MANUAL_PROVISIONING_ASSIGNEE_TYPE_UNSPECIFIED", "MANUAL_PROVISIONING_ASSIGNEE_TYPE_USER", "MANUAL_PROVISIONING_ASSIGNEE_TYPE_GROUP",
	)
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
	// Whether requests are enabled for the entitlement.
	RequestsEnabled param.Opt[bool] `json:"requestsEnabled,omitzero"`
	// Provisioning configuration. Exactly one method should be set.
	ProvisioningMethod AppResourceEntitlementUpdateParamsProvisioningMethod `json:"provisioningMethod,omitzero"`
	paramObj
}

func (r AppResourceEntitlementUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provisioning configuration. Exactly one method should be set.
type AppResourceEntitlementUpdateParamsProvisioningMethod struct {
	// **Option: builtin_workflow**
	BuiltinWorkflow any `json:"builtinWorkflow,omitzero"`
	// **Option: custom_workflow**
	CustomWorkflow AppResourceEntitlementUpdateParamsProvisioningMethodCustomWorkflow `json:"customWorkflow,omitzero"`
	// **Option: linked_entitlements**
	LinkedEntitlements AppResourceEntitlementUpdateParamsProvisioningMethodLinkedEntitlements `json:"linkedEntitlements,omitzero"`
	// **Option: manual**
	Manual AppResourceEntitlementUpdateParamsProvisioningMethodManual `json:"manual,omitzero"`
	paramObj
}

func (r AppResourceEntitlementUpdateParamsProvisioningMethod) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementUpdateParamsProvisioningMethod
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementUpdateParamsProvisioningMethod) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Option: custom_workflow**
type AppResourceEntitlementUpdateParamsProvisioningMethodCustomWorkflow struct {
	// The workflow ID to deprovision access.
	DeprovisionWorkflowID param.Opt[string] `json:"deprovisionWorkflowId,omitzero"`
	// The workflow ID to provision access.
	ProvisionWorkflowID param.Opt[string] `json:"provisionWorkflowId,omitzero"`
	paramObj
}

func (r AppResourceEntitlementUpdateParamsProvisioningMethodCustomWorkflow) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementUpdateParamsProvisioningMethodCustomWorkflow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementUpdateParamsProvisioningMethodCustomWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Option: linked_entitlements**
type AppResourceEntitlementUpdateParamsProvisioningMethodLinkedEntitlements struct {
	// The IDs of prerequisite entitlements.
	LinkedEntitlementIDs []string `json:"linkedEntitlementIds,omitzero"`
	paramObj
}

func (r AppResourceEntitlementUpdateParamsProvisioningMethodLinkedEntitlements) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementUpdateParamsProvisioningMethodLinkedEntitlements
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementUpdateParamsProvisioningMethodLinkedEntitlements) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **Option: manual**
type AppResourceEntitlementUpdateParamsProvisioningMethodManual struct {
	// Users and groups that should be assigned/notified for manual provisioning.
	Assignees []AppResourceEntitlementUpdateParamsProvisioningMethodManualAssignee `json:"assignees,omitzero"`
	paramObj
}

func (r AppResourceEntitlementUpdateParamsProvisioningMethodManual) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementUpdateParamsProvisioningMethodManual
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementUpdateParamsProvisioningMethodManual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceEntitlementUpdateParamsProvisioningMethodManualAssignee struct {
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

func (r AppResourceEntitlementUpdateParamsProvisioningMethodManualAssignee) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementUpdateParamsProvisioningMethodManualAssignee
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementUpdateParamsProvisioningMethodManualAssignee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AppResourceEntitlementUpdateParamsProvisioningMethodManualAssignee](
		"assigneeType", "MANUAL_PROVISIONING_ASSIGNEE_TYPE_UNSPECIFIED", "MANUAL_PROVISIONING_ASSIGNEE_TYPE_USER", "MANUAL_PROVISIONING_ASSIGNEE_TYPE_GROUP",
	)
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
