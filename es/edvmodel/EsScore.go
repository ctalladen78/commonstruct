package edvmodel

import "time"

//EsScore holds the score in elasticsearch
type EsScore struct {
	ARN       *string `json:"arn,omitempty"`
	ClientARN *string `json:"client_arn,omitempty"`
	Code      *string `json:"code,omitempty"`

	Title   *string `json:"title,omitempty"`
	Comment *string `json:"comment"`

	//ScoringDate is when the score is written by the scorer
	ScoringDate *time.Time `json:"scoring_date,omitempty"`

	//ReleaseDate is the date of release of the score so that it can be viewed by the students
	ReleaseDate *time.Time `json:"release_date,omitempty"`

	//StudentUsername is the student username for this score
	//e.g. kkesley/student
	StudentUsername *string `json:"student_username,omitempty"`

	//Targets can be either of the following:
	//1. Assessment scores, e.g. ["ASSESSMENT/OSDJIOIASD-ASDASD"]
	//2. Final scores e.g. ["FINAL"]
	Targets []string `json:"targets,omitempty"`

	//Course will be the course of this score
	//e.g. COURSE/OIWMMW-JASDUSDm
	Course *string `json:"course,omitempty"`

	Weight *string `json:"weight,omitempty"` //if empty or AUTO, weight will fill the remaining weight

	Score       *float64 `json:"score,omitempty"`
	ScoreLetter *string  `json:"score_letter,omitempty"` //optional if exist e.g. A, A+, A-

	AssessmentType *string `json:"assessment_type,omitempty"` //if not from assessment, can use assessment type (e.g. behaviour)

	Principal *string    `json:"principal,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	UpdatedBy *string    `json:"updated_by,omitempty"`
}
