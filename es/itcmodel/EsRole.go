package itcmodel

import (
	"time"
)

//EsRole holds the role in elastic search
type EsRole struct {
	ARN       *string    `json:"arn,omitempty"`
	ClientKey *string    `json:"client_key,omitempty"` //if it's predefined, use "-", else, use client_arn
	RoleID    *string    `json:"role_id,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Active    *bool      `json:"active,omitempty"`
	Principal *string    `json:"principal,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	IsDeleted *bool      `json:"is_deleted,omitempty"`
}
