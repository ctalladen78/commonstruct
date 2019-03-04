package edvmodel

import (
	"time"

	"github.com/kkesley/commonstruct/es/edvmodel"
)

//EsCourse holds the course in elasticsearch
type EsCourse struct {
	ARN             *string `json:"arn,omitempty"`
	ClientARN       *string `json:"client_arn,omitempty"`
	Code            *string `json:"code,omitempty"`
	CustomCode      *string `json:"custom_code,omitempty"`
	Name            *string `json:"name,omitempty"`
	Description     *string `json:"description"`
	DefaultDuration *int    `json:"default_duration"` //in minutes

	Levels []string `json:"levels,omitempty"`

	//Teachers hold the user arn of teachers
	//This can be used for home teachers for this class. This field can also be used if every courses is taught per class basis,
	Teachers              []string                          `json:"teachers,omitempty"`
	Settings              map[string]string                 `json:"settings,omitempty"`
	Principal             *string                           `json:"principal,omitempty"`
	CreatedAt             *time.Time                        `json:"created_at,omitempty"`
	UpdatedAt             *time.Time                        `json:"updated_at,omitempty"`
	UpdatedBy             *string                           `json:"updated_by,omitempty"`
	LevelOverrides        map[string]EsCourseOverride       `json:"level_overrides,omitempty"`
	AssessmentTypeWeights []edvmodel.EsCourseAssessmentType `json:"assessment_type_weights,omitempty"`
}

//EsCourseOverride enables overriding attributes for different levels
type EsCourseOverride struct {
	Name        *string  `json:"name,omitempty"`
	Description *string  `json:"description,omitempty"`
	Teachers    []string `json:"teachers,omitempty"`
}

//EsCourseAssessmentType holds assessment_types for a level
//e.g. Math course has multiple weighting ["QUIZ": 20%, "EXAM": 40%] etc
type EsCourseAssessmentType struct {
	Level           *string                        `json:"level,omitempty"`
	AssessmentTypes []EsCourseAssessmentTypeWeight `json:"assessment_types,omitempty"`
}

//EsCourseAssessmentTypeWeight holds the actual weight
//e.g. "EXAM" => 20%
type EsCourseAssessmentTypeWeight struct {
	AssessmentType *string  `json:"assessment_type,omitempty"`
	Weight         *float64 `json:"weight,omitempty"`
}
