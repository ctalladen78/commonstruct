package itcmodel

import (
	"time"
)

//EsTicket holds the tickets
type EsTicket struct {
	ARN       *string `json:"arn,omitempty"`
	ClientARN *string `json:"client_arn,omitempty"`
	Code      *string `json:"code,omitempty"`

	ConversationRef *string `json:"conversation_ref,omitempty"` //group the tickets to a conversation

	SenderARN     *string  `json:"sender_arn,omitempty"`     //sender ARN
	SenderService *string  `json:"sender_service,omitempty"` //sender service
	TargetARNs    []string `json:"target_arns,omitempty"`    //target ARNs

	ServiceScopes []string `json:"service_scopes,omitempty"` //platform / edvise

	InternalType     *string           `json:"internal_type,omitempty"`
	InternalSubtype  *string           `json:"internal_subtype,omitempty"`
	OriginalTicket   *bool             `json:"original_ticket,omitempty"`
	Closed           *bool             `json:"closed,omitempty"`
	ClosedDate       *bool             `json:"resolved_date,omitempty"`
	IsResolved       *bool             `json:"is_resolved,omitempty"`
	UnresolvedReason *string           `json:"unresolved_reason,omitempty"`
	Title            *string           `json:"title,omitempty"`
	Description      *string           `json:"description,omitempty"`
	Data             []TicketDataValue `json:"data,omitempty"`
	CreatedAt        *time.Time        `json:"created_at,omitempty"`
}

//TicketDataValue represent a data key
type TicketDataValue struct {
	Key   *string `json:"key"`
	Value *string `json:"value"`
}
