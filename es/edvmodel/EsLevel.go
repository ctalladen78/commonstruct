package edvmodel

import "time"

//EsLevel holds the class in elasticsearch
type EsLevel struct {
	ARN        *string `json:"arn,omitempty"`
	ClientARN  *string `json:"client_arn,omitempty"`
	Code       *string `json:"code,omitempty"`
	CustomCode *string `json:"custom_code,omitempty"`
	Name       *string `json:"name,omitempty"`

	Courses []string `json:"courses,omitempty"`

	Principal *string    `json:"principal,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	UpdatedBy *string    `json:"updated_by,omitempty"`
}
