package edviseperm

//FragmentAction holds additional feature of a resource
type FragmentAction struct {
	Edvise EdviseFragment `json:"edvise"`
}

//EdviseFragment holds the fragment of edvise
type EdviseFragment struct {
	StudentComment       string `json:"student_comment"`
	AssessmentDiscussion string `json:"assessment_discussion"`
}
