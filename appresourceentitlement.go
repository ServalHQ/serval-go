// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package serval

import (
	"context"
	"encoding/json"
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
	// Data from the external system as a JSON string (optional).
	ExternalData string `json:"externalData,nullable"`
	// The external ID of the entitlement in the external system (optional).
	ExternalID string `json:"externalId,nullable"`
	// The name of the entitlement.
	Name string `json:"name"`
	// Provisioning configuration. **Exactly one method should be set.**
	ProvisioningMethod AppResourceEntitlementProvisioningMethodUnion `json:"provisioningMethod"`
	// Whether requests are enabled for the entitlement.
	RequestsEnabled bool `json:"requestsEnabled"`
	// The ID of the resource that the entitlement belongs to.
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
func (r AppResourceEntitlement) RawJSON() string { return r.JSON.raw }
func (r *AppResourceEntitlement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AppResourceEntitlementProvisioningMethodUnion contains all possible properties
// and values from [AppResourceEntitlementProvisioningMethodBuiltinWorkflow],
// [AppResourceEntitlementProvisioningMethodCustomWorkflow],
// [AppResourceEntitlementProvisioningMethodLinkedEntitlements],
// [AppResourceEntitlementProvisioningMethodManual].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AppResourceEntitlementProvisioningMethodUnion struct {
	// This field is from variant
	// [AppResourceEntitlementProvisioningMethodBuiltinWorkflow].
	BuiltinWorkflow any `json:"builtinWorkflow"`
	// This field is from variant
	// [AppResourceEntitlementProvisioningMethodCustomWorkflow].
	CustomWorkflow AppResourceEntitlementProvisioningMethodCustomWorkflowCustomWorkflow `json:"customWorkflow"`
	// This field is from variant
	// [AppResourceEntitlementProvisioningMethodLinkedEntitlements].
	LinkedEntitlements AppResourceEntitlementProvisioningMethodLinkedEntitlementsLinkedEntitlements `json:"linkedEntitlements"`
	// This field is from variant [AppResourceEntitlementProvisioningMethodManual].
	Manual AppResourceEntitlementProvisioningMethodManualManual `json:"manual"`
	JSON   struct {
		BuiltinWorkflow    respjson.Field
		CustomWorkflow     respjson.Field
		LinkedEntitlements respjson.Field
		Manual             respjson.Field
		raw                string
	} `json:"-"`
}

func (u AppResourceEntitlementProvisioningMethodUnion) AsBuiltinWorkflow() (v AppResourceEntitlementProvisioningMethodBuiltinWorkflow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AppResourceEntitlementProvisioningMethodUnion) AsCustomWorkflow() (v AppResourceEntitlementProvisioningMethodCustomWorkflow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AppResourceEntitlementProvisioningMethodUnion) AsLinkedEntitlements() (v AppResourceEntitlementProvisioningMethodLinkedEntitlements) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AppResourceEntitlementProvisioningMethodUnion) AsManual() (v AppResourceEntitlementProvisioningMethodManual) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AppResourceEntitlementProvisioningMethodUnion) RawJSON() string { return u.JSON.raw }

func (r *AppResourceEntitlementProvisioningMethodUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceEntitlementProvisioningMethodBuiltinWorkflow struct {
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
func (r AppResourceEntitlementProvisioningMethodBuiltinWorkflow) RawJSON() string { return r.JSON.raw }
func (r *AppResourceEntitlementProvisioningMethodBuiltinWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceEntitlementProvisioningMethodCustomWorkflow struct {
	// Provisioning is handled by custom workflows for provision + deprovision.
	CustomWorkflow AppResourceEntitlementProvisioningMethodCustomWorkflowCustomWorkflow `json:"customWorkflow,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomWorkflow respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceEntitlementProvisioningMethodCustomWorkflow) RawJSON() string { return r.JSON.raw }
func (r *AppResourceEntitlementProvisioningMethodCustomWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provisioning is handled by custom workflows for provision + deprovision.
type AppResourceEntitlementProvisioningMethodCustomWorkflowCustomWorkflow struct {
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
func (r AppResourceEntitlementProvisioningMethodCustomWorkflowCustomWorkflow) RawJSON() string {
	return r.JSON.raw
}
func (r *AppResourceEntitlementProvisioningMethodCustomWorkflowCustomWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceEntitlementProvisioningMethodLinkedEntitlements struct {
	// Provisioning depends on prerequisite entitlements being provisioned first.
	LinkedEntitlements AppResourceEntitlementProvisioningMethodLinkedEntitlementsLinkedEntitlements `json:"linkedEntitlements,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LinkedEntitlements respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceEntitlementProvisioningMethodLinkedEntitlements) RawJSON() string {
	return r.JSON.raw
}
func (r *AppResourceEntitlementProvisioningMethodLinkedEntitlements) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provisioning depends on prerequisite entitlements being provisioned first.
type AppResourceEntitlementProvisioningMethodLinkedEntitlementsLinkedEntitlements struct {
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
func (r AppResourceEntitlementProvisioningMethodLinkedEntitlementsLinkedEntitlements) RawJSON() string {
	return r.JSON.raw
}
func (r *AppResourceEntitlementProvisioningMethodLinkedEntitlementsLinkedEntitlements) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceEntitlementProvisioningMethodManual struct {
	// Provisioning is handled manually by assigned users/groups.
	Manual AppResourceEntitlementProvisioningMethodManualManual `json:"manual,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Manual      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceEntitlementProvisioningMethodManual) RawJSON() string { return r.JSON.raw }
func (r *AppResourceEntitlementProvisioningMethodManual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provisioning is handled manually by assigned users/groups.
type AppResourceEntitlementProvisioningMethodManualManual struct {
	// Users and groups that should be assigned/notified for manual provisioning.
	Assignees []AppResourceEntitlementProvisioningMethodManualManualAssignee `json:"assignees"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Assignees   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppResourceEntitlementProvisioningMethodManualManual) RawJSON() string { return r.JSON.raw }
func (r *AppResourceEntitlementProvisioningMethodManualManual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceEntitlementProvisioningMethodManualManualAssignee struct {
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
func (r AppResourceEntitlementProvisioningMethodManualManualAssignee) RawJSON() string {
	return r.JSON.raw
}
func (r *AppResourceEntitlementProvisioningMethodManualManualAssignee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceEntitlementDeleteResponse = any

type AppResourceEntitlementNewParams struct {
	// The default access policy for the entitlement (optional).
	AccessPolicyID param.Opt[string] `json:"accessPolicyId,omitzero"`
	// Data from the external system as a JSON string (optional).
	ExternalData param.Opt[string] `json:"externalData,omitzero"`
	// The external ID of the entitlement in the external system (optional).
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// A description of the entitlement.
	Description param.Opt[string] `json:"description,omitzero"`
	// The name of the entitlement.
	Name param.Opt[string] `json:"name,omitzero"`
	// Whether requests are enabled for the entitlement.
	RequestsEnabled param.Opt[bool] `json:"requestsEnabled,omitzero"`
	// The ID of the resource.
	ResourceID param.Opt[string] `json:"resourceId,omitzero"`
	// Provisioning configuration. Exactly one method should be set.
	ProvisioningMethod AppResourceEntitlementNewParamsProvisioningMethodUnion `json:"provisioningMethod,omitzero"`
	paramObj
}

func (r AppResourceEntitlementNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AppResourceEntitlementNewParamsProvisioningMethodUnion struct {
	OfBuiltinWorkflow    *AppResourceEntitlementNewParamsProvisioningMethodBuiltinWorkflow    `json:",omitzero,inline"`
	OfCustomWorkflow     *AppResourceEntitlementNewParamsProvisioningMethodCustomWorkflow     `json:",omitzero,inline"`
	OfLinkedEntitlements *AppResourceEntitlementNewParamsProvisioningMethodLinkedEntitlements `json:",omitzero,inline"`
	OfManual             *AppResourceEntitlementNewParamsProvisioningMethodManual             `json:",omitzero,inline"`
	paramUnion
}

func (u AppResourceEntitlementNewParamsProvisioningMethodUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBuiltinWorkflow, u.OfCustomWorkflow, u.OfLinkedEntitlements, u.OfManual)
}
func (u *AppResourceEntitlementNewParamsProvisioningMethodUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AppResourceEntitlementNewParamsProvisioningMethodUnion) asAny() any {
	if !param.IsOmitted(u.OfBuiltinWorkflow) {
		return u.OfBuiltinWorkflow
	} else if !param.IsOmitted(u.OfCustomWorkflow) {
		return u.OfCustomWorkflow
	} else if !param.IsOmitted(u.OfLinkedEntitlements) {
		return u.OfLinkedEntitlements
	} else if !param.IsOmitted(u.OfManual) {
		return u.OfManual
	}
	return nil
}

// The property BuiltinWorkflow is required.
type AppResourceEntitlementNewParamsProvisioningMethodBuiltinWorkflow struct {
	// Provisioning is handled by the service's builtin workflow integration.
	BuiltinWorkflow any `json:"builtinWorkflow,omitzero,required"`
	paramObj
}

func (r AppResourceEntitlementNewParamsProvisioningMethodBuiltinWorkflow) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementNewParamsProvisioningMethodBuiltinWorkflow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementNewParamsProvisioningMethodBuiltinWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property CustomWorkflow is required.
type AppResourceEntitlementNewParamsProvisioningMethodCustomWorkflow struct {
	// Provisioning is handled by custom workflows for provision + deprovision.
	CustomWorkflow AppResourceEntitlementNewParamsProvisioningMethodCustomWorkflowCustomWorkflow `json:"customWorkflow,omitzero,required"`
	paramObj
}

func (r AppResourceEntitlementNewParamsProvisioningMethodCustomWorkflow) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementNewParamsProvisioningMethodCustomWorkflow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementNewParamsProvisioningMethodCustomWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provisioning is handled by custom workflows for provision + deprovision.
type AppResourceEntitlementNewParamsProvisioningMethodCustomWorkflowCustomWorkflow struct {
	// The workflow ID to deprovision access.
	DeprovisionWorkflowID param.Opt[string] `json:"deprovisionWorkflowId,omitzero"`
	// The workflow ID to provision access.
	ProvisionWorkflowID param.Opt[string] `json:"provisionWorkflowId,omitzero"`
	paramObj
}

func (r AppResourceEntitlementNewParamsProvisioningMethodCustomWorkflowCustomWorkflow) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementNewParamsProvisioningMethodCustomWorkflowCustomWorkflow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementNewParamsProvisioningMethodCustomWorkflowCustomWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property LinkedEntitlements is required.
type AppResourceEntitlementNewParamsProvisioningMethodLinkedEntitlements struct {
	// Provisioning depends on prerequisite entitlements being provisioned first.
	LinkedEntitlements AppResourceEntitlementNewParamsProvisioningMethodLinkedEntitlementsLinkedEntitlements `json:"linkedEntitlements,omitzero,required"`
	paramObj
}

func (r AppResourceEntitlementNewParamsProvisioningMethodLinkedEntitlements) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementNewParamsProvisioningMethodLinkedEntitlements
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementNewParamsProvisioningMethodLinkedEntitlements) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provisioning depends on prerequisite entitlements being provisioned first.
type AppResourceEntitlementNewParamsProvisioningMethodLinkedEntitlementsLinkedEntitlements struct {
	// The IDs of prerequisite entitlements.
	LinkedEntitlementIDs []string `json:"linkedEntitlementIds,omitzero"`
	paramObj
}

func (r AppResourceEntitlementNewParamsProvisioningMethodLinkedEntitlementsLinkedEntitlements) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementNewParamsProvisioningMethodLinkedEntitlementsLinkedEntitlements
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementNewParamsProvisioningMethodLinkedEntitlementsLinkedEntitlements) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Manual is required.
type AppResourceEntitlementNewParamsProvisioningMethodManual struct {
	// Provisioning is handled manually by assigned users/groups.
	Manual AppResourceEntitlementNewParamsProvisioningMethodManualManual `json:"manual,omitzero,required"`
	paramObj
}

func (r AppResourceEntitlementNewParamsProvisioningMethodManual) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementNewParamsProvisioningMethodManual
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementNewParamsProvisioningMethodManual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provisioning is handled manually by assigned users/groups.
type AppResourceEntitlementNewParamsProvisioningMethodManualManual struct {
	// Users and groups that should be assigned/notified for manual provisioning.
	Assignees []AppResourceEntitlementNewParamsProvisioningMethodManualManualAssignee `json:"assignees,omitzero"`
	paramObj
}

func (r AppResourceEntitlementNewParamsProvisioningMethodManualManual) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementNewParamsProvisioningMethodManualManual
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementNewParamsProvisioningMethodManualManual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceEntitlementNewParamsProvisioningMethodManualManualAssignee struct {
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

func (r AppResourceEntitlementNewParamsProvisioningMethodManualManualAssignee) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementNewParamsProvisioningMethodManualManualAssignee
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementNewParamsProvisioningMethodManualManualAssignee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AppResourceEntitlementNewParamsProvisioningMethodManualManualAssignee](
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
	// Data from the external system as a JSON string (optional).
	ExternalData param.Opt[string] `json:"externalData,omitzero"`
	// The external ID of the entitlement in the external system (optional).
	ExternalID param.Opt[string] `json:"externalId,omitzero"`
	// A description of the entitlement.
	Description param.Opt[string] `json:"description,omitzero"`
	// The name of the entitlement.
	Name param.Opt[string] `json:"name,omitzero"`
	// Whether requests are enabled for the entitlement.
	RequestsEnabled param.Opt[bool] `json:"requestsEnabled,omitzero"`
	// Provisioning configuration. Exactly one method should be set.
	ProvisioningMethod AppResourceEntitlementUpdateParamsProvisioningMethodUnion `json:"provisioningMethod,omitzero"`
	paramObj
}

func (r AppResourceEntitlementUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AppResourceEntitlementUpdateParamsProvisioningMethodUnion struct {
	OfBuiltinWorkflow    *AppResourceEntitlementUpdateParamsProvisioningMethodBuiltinWorkflow    `json:",omitzero,inline"`
	OfCustomWorkflow     *AppResourceEntitlementUpdateParamsProvisioningMethodCustomWorkflow     `json:",omitzero,inline"`
	OfLinkedEntitlements *AppResourceEntitlementUpdateParamsProvisioningMethodLinkedEntitlements `json:",omitzero,inline"`
	OfManual             *AppResourceEntitlementUpdateParamsProvisioningMethodManual             `json:",omitzero,inline"`
	paramUnion
}

func (u AppResourceEntitlementUpdateParamsProvisioningMethodUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBuiltinWorkflow, u.OfCustomWorkflow, u.OfLinkedEntitlements, u.OfManual)
}
func (u *AppResourceEntitlementUpdateParamsProvisioningMethodUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AppResourceEntitlementUpdateParamsProvisioningMethodUnion) asAny() any {
	if !param.IsOmitted(u.OfBuiltinWorkflow) {
		return u.OfBuiltinWorkflow
	} else if !param.IsOmitted(u.OfCustomWorkflow) {
		return u.OfCustomWorkflow
	} else if !param.IsOmitted(u.OfLinkedEntitlements) {
		return u.OfLinkedEntitlements
	} else if !param.IsOmitted(u.OfManual) {
		return u.OfManual
	}
	return nil
}

// The property BuiltinWorkflow is required.
type AppResourceEntitlementUpdateParamsProvisioningMethodBuiltinWorkflow struct {
	// Provisioning is handled by the service's builtin workflow integration.
	BuiltinWorkflow any `json:"builtinWorkflow,omitzero,required"`
	paramObj
}

func (r AppResourceEntitlementUpdateParamsProvisioningMethodBuiltinWorkflow) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementUpdateParamsProvisioningMethodBuiltinWorkflow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementUpdateParamsProvisioningMethodBuiltinWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property CustomWorkflow is required.
type AppResourceEntitlementUpdateParamsProvisioningMethodCustomWorkflow struct {
	// Provisioning is handled by custom workflows for provision + deprovision.
	CustomWorkflow AppResourceEntitlementUpdateParamsProvisioningMethodCustomWorkflowCustomWorkflow `json:"customWorkflow,omitzero,required"`
	paramObj
}

func (r AppResourceEntitlementUpdateParamsProvisioningMethodCustomWorkflow) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementUpdateParamsProvisioningMethodCustomWorkflow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementUpdateParamsProvisioningMethodCustomWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provisioning is handled by custom workflows for provision + deprovision.
type AppResourceEntitlementUpdateParamsProvisioningMethodCustomWorkflowCustomWorkflow struct {
	// The workflow ID to deprovision access.
	DeprovisionWorkflowID param.Opt[string] `json:"deprovisionWorkflowId,omitzero"`
	// The workflow ID to provision access.
	ProvisionWorkflowID param.Opt[string] `json:"provisionWorkflowId,omitzero"`
	paramObj
}

func (r AppResourceEntitlementUpdateParamsProvisioningMethodCustomWorkflowCustomWorkflow) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementUpdateParamsProvisioningMethodCustomWorkflowCustomWorkflow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementUpdateParamsProvisioningMethodCustomWorkflowCustomWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property LinkedEntitlements is required.
type AppResourceEntitlementUpdateParamsProvisioningMethodLinkedEntitlements struct {
	// Provisioning depends on prerequisite entitlements being provisioned first.
	LinkedEntitlements AppResourceEntitlementUpdateParamsProvisioningMethodLinkedEntitlementsLinkedEntitlements `json:"linkedEntitlements,omitzero,required"`
	paramObj
}

func (r AppResourceEntitlementUpdateParamsProvisioningMethodLinkedEntitlements) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementUpdateParamsProvisioningMethodLinkedEntitlements
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementUpdateParamsProvisioningMethodLinkedEntitlements) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provisioning depends on prerequisite entitlements being provisioned first.
type AppResourceEntitlementUpdateParamsProvisioningMethodLinkedEntitlementsLinkedEntitlements struct {
	// The IDs of prerequisite entitlements.
	LinkedEntitlementIDs []string `json:"linkedEntitlementIds,omitzero"`
	paramObj
}

func (r AppResourceEntitlementUpdateParamsProvisioningMethodLinkedEntitlementsLinkedEntitlements) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementUpdateParamsProvisioningMethodLinkedEntitlementsLinkedEntitlements
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementUpdateParamsProvisioningMethodLinkedEntitlementsLinkedEntitlements) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Manual is required.
type AppResourceEntitlementUpdateParamsProvisioningMethodManual struct {
	// Provisioning is handled manually by assigned users/groups.
	Manual AppResourceEntitlementUpdateParamsProvisioningMethodManualManual `json:"manual,omitzero,required"`
	paramObj
}

func (r AppResourceEntitlementUpdateParamsProvisioningMethodManual) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementUpdateParamsProvisioningMethodManual
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementUpdateParamsProvisioningMethodManual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Provisioning is handled manually by assigned users/groups.
type AppResourceEntitlementUpdateParamsProvisioningMethodManualManual struct {
	// Users and groups that should be assigned/notified for manual provisioning.
	Assignees []AppResourceEntitlementUpdateParamsProvisioningMethodManualManualAssignee `json:"assignees,omitzero"`
	paramObj
}

func (r AppResourceEntitlementUpdateParamsProvisioningMethodManualManual) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementUpdateParamsProvisioningMethodManualManual
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementUpdateParamsProvisioningMethodManualManual) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AppResourceEntitlementUpdateParamsProvisioningMethodManualManualAssignee struct {
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

func (r AppResourceEntitlementUpdateParamsProvisioningMethodManualManualAssignee) MarshalJSON() (data []byte, err error) {
	type shadow AppResourceEntitlementUpdateParamsProvisioningMethodManualManualAssignee
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppResourceEntitlementUpdateParamsProvisioningMethodManualManualAssignee) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AppResourceEntitlementUpdateParamsProvisioningMethodManualManualAssignee](
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
