package edvmodel

import (
	"time"
)

//EsReport holds the report in elasticsearch
type EsReport struct {
	ARN       *string `json:"arn,omitempty"`
	ClientARN *string `json:"client_arn,omitempty"`
	Code      *string `json:"code,omitempty"`

	ReportBatch *string `json:"report_batch,omitempty"`

	Title   *string `json:"title,omitempty"`
	Student *string `json:"student,omitempty"`
	Class   *string `json:"class,omitempty"`

	ReportLocation *string  `json:"report_location,omitempty"`
	ExcludedScores []string `json:"excluded_scores,omitempty"`
	Scores         []string `json:"scores,omitempty"`

	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate   *time.Time `json:"end_date,omitempty"`

	ReleaseDate *time.Time `json:"release_date,omitempty"`

	Principal            *string                       `json:"principal,omitempty"`
	CreatedAt            *time.Time                    `json:"created_at,omitempty"`
	UpdatedAt            *time.Time                    `json:"updated_at,omitempty"`
	UpdatedBy            *string                       `json:"updated_by,omitempty"`
	AssessmentTypeScores []EsReportAssessmentTypeScore `json:"assessment_type_scores,omitempty"`
	CommentSections      []EsReportCommentSection      `json:"comment_sections,omitempty"`
}

//EsReportAssessmentTypeScore hold the assessment type accumulated scores (unweighted)
type EsReportAssessmentTypeScore struct {
	Course          *string                  `json:"course"`
	AssessmentType  *string                  `json:"assessment_type"`
	Weight          *float64                 `json:"weight"`
	UnweightedScore *float64                 `json:"unweighted_score"`
	ScoreCount      *int64                   `json:"score_count"`
	OverrideWeights []EsReportOverrideWeight `json:"override_weights"`
}

//EsReportOverrideWeight holds overriden weights for a specific score
type EsReportOverrideWeight struct {
	Key             *string  `json:"key"`
	Weight          *float64 `json:"weight"`
	UnweightedScore *float64 `json:"unweighted_score"`
}

//EsReportCommentSection holds the comment for a section in the report
type EsReportCommentSection struct {
	Section *string `json:"section"`
	Comment *string `json:"comment"`
}
