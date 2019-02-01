package edvmodel

import "time"

//EsAssessment holds the assessment in elasticsearch
type EsAssessment struct {
	ARN       *string `json:"arn,omitempty"`
	ClientARN *string `json:"client_arn,omitempty"`
	Code      *string `json:"code,omitempty"`

	AssessmentType *string  `json:"assessment_type,omitempty"`
	Title          *string  `json:"title,omitempty"`
	Description    *string  `json:"description,omitempty"`
	Attachments    []string `json:"attachments,omitempty"`
	Weight         *float64 `json:"weight,omitempty"` //out of 100

	ReleaseDate  *time.Time `json:"release_date,omitempty"`
	DeadlineDate *time.Time `json:"deadline_date,omitempty"`

	Classes []string `json:"classes,omitempty"`
	Course  *string  `json:"course,omitempty"`

	Principal *string    `json:"principal,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	UpdatedBy *string    `json:"updated_by,omitempty"`
}
