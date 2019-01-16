package itcmodel

import (
	"time"
)

//EsActivity holds the user activity in the portal
type EsActivity struct {
	UserARN    *string    `json:"user_arn,omitempty"`
	ClientARN  *string    `json:"client_arn,omitempty"`
	Session    *string    `json:"session,omitempty"`
	Action     *string    `json:"action,omitempty"`
	TargetType *string    `json:"target_type,omitempty"`
	TargetOld  *string    `json:"target_old,omitempty"`
	TargetNew  *string    `json:"target_new,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
}
