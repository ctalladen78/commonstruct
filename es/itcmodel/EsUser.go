package itcmodel

import (
	"time"
)

//EsUser holds the user in elastic search
type EsUser struct {
	ARN                *string    `json:"arn,omitempty"`
	ClientKey          *string    `json:"client_key,omitempty"` //this is for the primary key. Root users will have * as the value. non root user will have their client_arn
	ClientARN          *string    `json:"client_arn,omitempty"` //this is the original client_arn
	Username           *string    `json:"username,omitempty"`
	IsRoot             *bool      `json:"is_root,omitempty"`
	Email              *string    `json:"email"`
	FirstName          *string    `json:"first_name"`
	LastName           *string    `json:"last_name"`
	Phone              *string    `json:"phone"`
	Title              *string    `json:"title"`
	Active             *bool      `json:"active,omitempty"`
	Principal          *string    `json:"principal,omitempty"`
	PasswordNeedChange *bool      `json:"password_need_change,omitempty"`
	EmailVerified      *bool      `json:"email_verified,omitempty"`
	PhoneVerified      *bool      `json:"phone_verified,omitempty"`
	CreatedAt          *time.Time `json:"created_at,omitempty"`
	UpdatedAt          *time.Time `json:"updated_at,omitempty"`
	IsDeleted          *bool      `json:"is_deleted,omitempty"`
	DeletedBy          *string    `json:"deleted_by,omitempty"`
	DeletedAt          *time.Time `json:"deleted_at,omitempty"`

	IsMFAEnabled *bool             `json:"is_mfa_enabled,omitempty"`
	LatestLogin  *time.Time        `json:"latest_login,omitempty"`
	Roles        []string          `json:"roles,omitempty"`
	Attributes   map[string]string `json:"attributes,omitempty"`
}
