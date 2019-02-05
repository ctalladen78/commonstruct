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
	Weight         *string  `json:"weight,omitempty"` //out of 100

	ManualPublished        *bool                         `json:"manual_published,omitempty"`
	ReleaseDate            *time.Time                    `json:"release_date,omitempty"`
	DeadlineDate           *time.Time                    `json:"deadline_date,omitempty"`
	CustomDeadlineDate     *[]AssessmentDeadlineDate     `json:"custom_deadline_date,omitempty"`
	CustomDeadlineSchedule *[]AssessmentDeadlineSchedule `json:"custom_deadline_schedule,omitempty"`

	IsRemedial                 *bool   `json:"is_remedial,omitempty"`
	RemedialOriginalAssessment *string `json:"remedial_original_assessment,omitempty"`

	Classes []string `json:"classes,omitempty"`
	Course  *string  `json:"course,omitempty"`

	Principal *string    `json:"principal,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	UpdatedBy *string    `json:"updated_by,omitempty"`
}

//AssessmentDeadlineDate holds the array of key, value of assessment date
type AssessmentDeadlineDate struct {
	Key  *string    `json:"key,omitempty"`
	Date *time.Time `json:"date,omitempty"`
}

//AssessmentDeadlineSchedule holds the array of key, value of assessment schedule
type AssessmentDeadlineSchedule struct {
	Key      *string `json:"key,omitempty"`
	Schedule *string `json:"schedule,omitempty"`
}
