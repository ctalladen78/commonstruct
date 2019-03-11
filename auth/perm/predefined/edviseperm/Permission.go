package edviseperm

//PermissionConstant hold the permission in the higher level
type PermissionConstant struct {
	Platform  PlatformConstant  `json:"platform"`
	Education EducationConstant `json:"education"`
	Talk      TalkConstant      `json:"talk"`
}

//PlatformConstant hold the permission strings in platform
type PlatformConstant struct {
	User   string `json:"user"`
	School string `json:"school"`
}

//EducationConstant hold the permission strings in education
type EducationConstant struct {
	Level          string `json:"level"`
	Course         string `json:"course"`
	Class          string `json:"class"`
	Schedule       string `json:"schedule"`
	Assessment     string `json:"assessment"`
	AssessmentType string `json:"assessment_type"`
	Score          string `json:"score"`
	Attendance     string `json:"attendance"`
	Report         string `json:"report"`
}

//TalkConstant holds the permission strings in talk
type TalkConstant struct {
	Notice string `json:"notice"`
}

//Permission predefined permission
var Permission = PermissionConstant{
	Platform: PlatformConstant{
		User:   "edvise::platform::user",
		School: "edvise::platform::school",
	},
	Education: EducationConstant{
		Level:          "edvise::education::level",
		Course:         "edvise::education::course",
		Class:          "edvise::education::class",
		Schedule:       "edvise::education::schedule",
		Assessment:     "edvise::education::assessment",
		AssessmentType: "edvise::education::assessment-type",
		Score:          "edvise::education::score",
		Attendance:     "edvise::education::attendance",
		Report:         "edvise::education::report",
	},
	Talk: TalkConstant{
		Notice: "edvise::talk::notice",
	},
}
