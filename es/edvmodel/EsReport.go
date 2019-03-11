package edvmodel

import (
	"time"
)

//EsReport holds the report in elasticsearch
type EsReport struct {
	ARN       *string `json:"arn,omitempty"`
	ClientARN *string `json:"client_arn,omitempty"`
	Code      *string `json:"code,omitempty"`

	Student *string `json:"student,omitempty"`
	Class   *string `json:"class,omitempty"`

	ReportLocation      *string  `json:"report_location,omitempty"`
	ExcludedAssessments []string `json:"excluded_assessments,omitempty"`
	ExcludedScores      []string `json:"excluded_scores,omitempty"`

	Principal             *string                        `json:"principal,omitempty"`
	CreatedAt             *time.Time                     `json:"created_at,omitempty"`
	UpdatedAt             *time.Time                     `json:"updated_at,omitempty"`
	UpdatedBy             *string                        `json:"updated_by,omitempty"`
	AssessmentTypeWeights []EsReportAssessmentTypeWeight `json:"assessment_type_weights,omitempty"`
	CommentSections       []EsReportCommentSection       `json:"comment_sections,omitempty"`
}

//EsReportAssessmentTypeWeight holds the assessment type weight
//e.g. "EXAM" => 20%
type EsReportAssessmentTypeWeight struct {
	AssessmentType *string  `json:"assessment_type"`
	Weight         *float64 `json:"weight"`
}

//EsReportCommentSection holds the comment for a section in the report
type EsReportCommentSection struct {
	Section *string `json:"section"`
	Comment *string `json:"comment"`
}
