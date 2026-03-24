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
	"time"

	"github.com/ServalHQ/serval-go/internal/apijson"
	"github.com/ServalHQ/serval-go/internal/apiquery"
	"github.com/ServalHQ/serval-go/internal/requestconfig"
	"github.com/ServalHQ/serval-go/option"
	"github.com/ServalHQ/serval-go/packages/pagination"
	"github.com/ServalHQ/serval-go/packages/param"
	"github.com/ServalHQ/serval-go/packages/respjson"
)

// TicketService contains methods and other services that help with interacting
// with the serval API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTicketService] method instead.
type TicketService struct {
	Options []option.RequestOption
}

// NewTicketService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTicketService(opts ...option.RequestOption) (r TicketService) {
	r = TicketService{}
	r.Options = opts
	return
}

// Create a new ticket.
func (r *TicketService) New(ctx context.Context, body TicketNewParams, opts ...option.RequestOption) (res *Ticket, err error) {
	var env TicketNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := "v2/tickets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// Get a specific ticket by ID.
func (r *TicketService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Ticket, err error) {
	var env TicketGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/tickets/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// Update an existing ticket.
func (r *TicketService) Update(ctx context.Context, id string, body TicketUpdateParams, opts ...option.RequestOption) (res *Ticket, err error) {
	var env TicketUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/tickets/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Data
	return res, nil
}

// List tickets for a team.
func (r *TicketService) List(ctx context.Context, query TicketListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Ticket], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v2/tickets"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List tickets for a team.
func (r *TicketService) ListAutoPaging(ctx context.Context, query TicketListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Ticket] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Ticket represents a Serval ticket.
type Ticket struct {
	// The ID of the Serval ticket.
	ID string `json:"id"`
	// A friendly identifier for the ticket (e.g. "ACM-10").
	FriendlyIdentifier string `json:"friendlyIdentifier"`
	// The ID of the team that the ticket belongs to.
	TeamID string `json:"teamId"`
	// A name or title for the ticket.
	Name string `json:"name"`
	// A description of the ticket.
	Description string `json:"description"`
	// Timestamp indicating when the ticket was created (RFC 3339 format).
	CreatedAt time.Time `json:"createdAt" format:"date-time"`
	// Timestamp indicating when the ticket was completed (RFC 3339 format).
	CompletedAt string `json:"completedAt" api:"nullable"`
	// Timestamp indicating when the ticket was escalated (RFC 3339 format).
	EscalatedAt string `json:"escalatedAt" api:"nullable"`
	// The ID of the user who created the ticket.
	CreatedByUserID string `json:"createdByUserId"`
	// The ID of the user who is assigned to the ticket.
	AssignedToUserID string `json:"assignedToUserId" api:"nullable"`
	// The ID of the user this ticket was created on behalf of (the requester).
	RequesterUserID string `json:"requesterUserId"`
	// The ID of the status option for the ticket.
	StatusID string `json:"statusId"`
	// The escalation level of the ticket.
	//
	// Any of "AI", "HUMAN".
	EscalationLevel TicketEscalationLevel `json:"escalationLevel"`
	// The ID of the priority option for the ticket.
	PriorityID string `json:"priorityId"`
	// Timestamp indicating when the SLA started (RFC 3339 format).
	SLAStartedAt string `json:"slaStartedAt" api:"nullable"`
	// Timestamp indicating when the SLA was breached (RFC 3339 format).
	SLABreachesAt string `json:"slaBreachesAt" api:"nullable"`
	// List of label IDs for the ticket.
	LabelIDs []string `json:"labelIds"`
	// The type of ticket.
	//
	// Any of "TICKET_TYPE_UNSPECIFIED", "TICKET_TYPE_REQUEST", "TICKET_TYPE_TASK",
	// "TICKET_TYPE_INCIDENT".
	Type TicketType `json:"type"`
	// List of relationship IDs for this ticket.
	RelationshipIDs []string `json:"relationshipIds"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		FriendlyIdentifier respjson.Field
		TeamID             respjson.Field
		Name               respjson.Field
		Description        respjson.Field
		CreatedAt          respjson.Field
		CompletedAt        respjson.Field
		EscalatedAt        respjson.Field
		CreatedByUserID    respjson.Field
		AssignedToUserID   respjson.Field
		RequesterUserID    respjson.Field
		StatusID           respjson.Field
		EscalationLevel    respjson.Field
		PriorityID         respjson.Field
		SLAStartedAt       respjson.Field
		SLABreachesAt      respjson.Field
		LabelIDs           respjson.Field
		Type               respjson.Field
		RelationshipIDs    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Ticket) RawJSON() string { return r.JSON.raw }
func (r *Ticket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The escalation level of a ticket.
type TicketEscalationLevel string

const (
	TicketEscalationLevelAI    TicketEscalationLevel = "AI"
	TicketEscalationLevelHuman TicketEscalationLevel = "HUMAN"
)

// The type of a ticket.
type TicketType string

const (
	TicketTypeTicketTypeUnspecified TicketType = "TICKET_TYPE_UNSPECIFIED"
	TicketTypeTicketTypeRequest     TicketType = "TICKET_TYPE_REQUEST"
	TicketTypeTicketTypeTask        TicketType = "TICKET_TYPE_TASK"
	TicketTypeTicketTypeIncident    TicketType = "TICKET_TYPE_INCIDENT"
)

// Which user to target for channel sync.
type ChannelSyncTargetUserType string

const (
	ChannelSyncTargetUserTypeUnspecified     ChannelSyncTargetUserType = "CHANNEL_SYNC_USER_TYPE_UNSPECIFIED"
	ChannelSyncTargetUserTypeTicketAssignee  ChannelSyncTargetUserType = "CHANNEL_SYNC_TARGET_USER_TICKET_ASSIGNEE"
	ChannelSyncTargetUserTypeTicketRequestor ChannelSyncTargetUserType = "CHANNEL_SYNC_TARGET_USER_TICKET_REQUESTOR"
)

// TicketNewParamsChannelSyncTargetUnion represents a channel sync target.
// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TicketNewParamsChannelSyncTargetUnion struct {
	OfEmail   *TicketNewParamsChannelSyncTargetEmail   `json:",omitzero,inline"`
	OfSlackDm *TicketNewParamsChannelSyncTargetSlackDm `json:",omitzero,inline"`
	paramUnion
}

func (u TicketNewParamsChannelSyncTargetUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfEmail, u.OfSlackDm)
}
func (u *TicketNewParamsChannelSyncTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TicketNewParamsChannelSyncTargetUnion) asAny() any {
	if !param.IsOmitted(u.OfEmail) {
		return u.OfEmail
	} else if !param.IsOmitted(u.OfSlackDm) {
		return u.OfSlackDm
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TicketNewParamsChannelSyncTargetUnion) GetEmail() *TicketNewParamsChannelSyncTargetEmailEmail {
	if vt := u.OfEmail; vt != nil {
		return &vt.Email
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u TicketNewParamsChannelSyncTargetUnion) GetSlackDm() *TicketNewParamsChannelSyncTargetSlackDmSlackDm {
	if vt := u.OfSlackDm; vt != nil {
		return &vt.SlackDm
	}
	return nil
}

// Email channel sync target variant.
//
// The property Email is required.
type TicketNewParamsChannelSyncTargetEmail struct {
	Email TicketNewParamsChannelSyncTargetEmailEmail `json:"email,omitzero" api:"required"`
	paramObj
}

func (r TicketNewParamsChannelSyncTargetEmail) MarshalJSON() (data []byte, err error) {
	type shadow TicketNewParamsChannelSyncTargetEmail
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TicketNewParamsChannelSyncTargetEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TicketNewParamsChannelSyncTargetEmailEmail struct {
	TargetUserType param.Opt[ChannelSyncTargetUserType] `json:"targetUserType,omitzero"`
	paramObj
}

func (r TicketNewParamsChannelSyncTargetEmailEmail) MarshalJSON() (data []byte, err error) {
	type shadow TicketNewParamsChannelSyncTargetEmailEmail
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TicketNewParamsChannelSyncTargetEmailEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Slack DM channel sync target variant.
//
// The property SlackDm is required.
type TicketNewParamsChannelSyncTargetSlackDm struct {
	SlackDm TicketNewParamsChannelSyncTargetSlackDmSlackDm `json:"slackDm,omitzero" api:"required"`
	paramObj
}

func (r TicketNewParamsChannelSyncTargetSlackDm) MarshalJSON() (data []byte, err error) {
	type shadow TicketNewParamsChannelSyncTargetSlackDm
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TicketNewParamsChannelSyncTargetSlackDm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TicketNewParamsChannelSyncTargetSlackDmSlackDm struct {
	TargetUserType param.Opt[ChannelSyncTargetUserType] `json:"targetUserType,omitzero"`
	paramObj
}

func (r TicketNewParamsChannelSyncTargetSlackDmSlackDm) MarshalJSON() (data []byte, err error) {
	type shadow TicketNewParamsChannelSyncTargetSlackDmSlackDm
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TicketNewParamsChannelSyncTargetSlackDmSlackDm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChannelSyncTargetUnion contains all possible properties and values from
// [ChannelSyncTargetEmail], [ChannelSyncTargetSlackDm].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChannelSyncTargetUnion struct {
	// This field is from variant [ChannelSyncTargetEmail].
	Email ChannelSyncTargetEmailEmail `json:"email"`
	// This field is from variant [ChannelSyncTargetSlackDm].
	SlackDm ChannelSyncTargetSlackDmSlackDm `json:"slackDm"`
	JSON    struct {
		Email   respjson.Field
		SlackDm respjson.Field
		raw     string
	} `json:"-"`
}

func (u ChannelSyncTargetUnion) AsEmail() (v ChannelSyncTargetEmail) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChannelSyncTargetUnion) AsSlackDm() (v ChannelSyncTargetSlackDm) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChannelSyncTargetUnion) RawJSON() string { return u.JSON.raw }

func (r *ChannelSyncTargetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChannelSyncTargetEmail struct {
	Email ChannelSyncTargetEmailEmail `json:"email" api:"required"`
	JSON  struct {
		Email       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

func (r ChannelSyncTargetEmail) RawJSON() string { return r.JSON.raw }
func (r *ChannelSyncTargetEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChannelSyncTargetEmailEmail struct {
	TargetUserType ChannelSyncTargetUserType `json:"targetUserType"`
	JSON           struct {
		TargetUserType respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

func (r ChannelSyncTargetEmailEmail) RawJSON() string { return r.JSON.raw }
func (r *ChannelSyncTargetEmailEmail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChannelSyncTargetSlackDm struct {
	SlackDm ChannelSyncTargetSlackDmSlackDm `json:"slackDm" api:"required"`
	JSON    struct {
		SlackDm     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

func (r ChannelSyncTargetSlackDm) RawJSON() string { return r.JSON.raw }
func (r *ChannelSyncTargetSlackDm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChannelSyncTargetSlackDmSlackDm struct {
	TargetUserType ChannelSyncTargetUserType `json:"targetUserType"`
	JSON           struct {
		TargetUserType respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

func (r ChannelSyncTargetSlackDmSlackDm) RawJSON() string { return r.JSON.raw }
func (r *ChannelSyncTargetSlackDmSlackDm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TicketNewParams struct {
	TeamID        string            `json:"teamId" api:"required"`
	Name          string            `json:"name" api:"required"`
	Description   param.Opt[string] `json:"description,omitzero"`
	CreatedByUserID param.Opt[string] `json:"createdByUserId,omitzero"`
	AssignedToUserID param.Opt[string] `json:"assignedToUserId,omitzero"`
	ParentTicketID param.Opt[string] `json:"parentTicketId,omitzero"`
	// Additional channel sync targets for the ticket. These specify external channels
	// (e.g. Slack DM, Email) that the ticket should sync to. Note that these are
	// ADDITIONAL sync targets; there may be default sync targets for the team.
	ChannelSyncTargets []TicketNewParamsChannelSyncTargetUnion `json:"channelSyncTargets,omitzero"`
	// Whether AI is active for this ticket. Defaults to false.
	AIActive param.Opt[bool] `json:"aiActive,omitzero"`
	// The type of ticket to create. Defaults to REQUEST.
	//
	// Any of "TICKET_TYPE_UNSPECIFIED", "TICKET_TYPE_REQUEST", "TICKET_TYPE_TASK",
	// "TICKET_TYPE_INCIDENT".
	Type param.Opt[TicketType] `json:"type,omitzero"`
	paramObj
}

func (r TicketNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TicketNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TicketNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TicketNewResponseEnvelope struct {
	Data Ticket `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TicketNewResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *TicketNewResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TicketGetResponseEnvelope struct {
	Data Ticket `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TicketGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *TicketGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TicketUpdateParams struct {
	Name             param.Opt[string]                `json:"name,omitzero"`
	Description      param.Opt[string]                `json:"description,omitzero"`
	StatusID         param.Opt[string]                `json:"statusId,omitzero"`
	PriorityID       param.Opt[string]                `json:"priorityId,omitzero"`
	AssignedToUserID param.Opt[string]                `json:"assignedToUserId,omitzero"`
	LabelIDs         []string                         `json:"labelIds,omitzero"`
	EscalationLevel  param.Opt[TicketEscalationLevel] `json:"escalationLevel,omitzero"`
	RequesterUserID  param.Opt[string]                `json:"requesterUserId,omitzero"`
	paramObj
}

func (r TicketUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TicketUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TicketUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TicketUpdateResponseEnvelope struct {
	Data Ticket `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TicketUpdateResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *TicketUpdateResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TicketListParams struct {
	// The ID of the team.
	TeamID param.Opt[string] `query:"teamId,omitzero" json:"-"`
	// Maximum number of results to return. Default is 1000, maximum is 1000.
	PageSize param.Opt[int64] `query:"pageSize,omitzero" json:"-"`
	// Token for pagination. Leave empty for the first request.
	PageToken param.Opt[string] `query:"pageToken,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TicketListParams]'s query parameters as `url.Values`.
func (r TicketListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
