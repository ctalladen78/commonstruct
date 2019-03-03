package edvmodel

import "time"

//EsAssessmentType holds the assessment type in elasticsearch
type EsAssessmentType struct {
	ARN       *string `json:"arn,omitempty"`
	ClientARN *string `json:"client_arn,omitempty"`
	Code      *string `json:"code,omitempty"`

	Name           *string `json:"name,omitempty"`
	Description    *string `json:"description,omitempty"`
	Classification *string `json:"classification,omitempty"` //classification of this assessment-type. e.g. exam, quiz, homework

	Principal *string    `json:"principal,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	UpdatedBy *string    `json:"updated_by,omitempty"`
}
