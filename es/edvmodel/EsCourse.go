package edvmodel

import "time"

//EsCourse holds the course in elasticsearch
type EsCourse struct {
	ARN             *string `json:"arn,omitempty"`
	ClientARN       *string `json:"client_arn,omitempty"`
	Code            *string `json:"code,omitempty"`
	CustomCode      *string `json:"custom_code,omitempty"`
	Name            *string `json:"name,omitempty"`
	Description     *string `json:"description,omitempty"`
	DefaultDuration *int    `json:"default_duration,omitempty"` //in minutes

	Principal *string    `json:"principal,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	UpdatedBy *string    `json:"updated_by,omitempty"`
}
