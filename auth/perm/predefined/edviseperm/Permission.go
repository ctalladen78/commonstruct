package edviseperm

//PermissionConstant hold the permission in the higher level
type PermissionConstant struct {
	Platform  PlatformConstant  `json:"platform"`
	Education EducationConstant `json:"education"`
}

//PlatformConstant hold the permission strings in platform
type PlatformConstant struct {
	User   string `json:"user"`
	School string `json:"school"`
}

//EducationConstant hold the permission strings in education
type EducationConstant struct {
	Level      string `json:"level"`
	Course     string `json:"course"`
	Class      string `json:"class"`
	Schedule   string `json:"schedule"`
	Assessment string `json:"assessment"`
	Score      string `json:"score"`
	Attendance string `json:"attendance"`
	Comment    string `json:"comment"`
}

//Permission predefined permission
var Permission = PermissionConstant{
	Platform: PlatformConstant{
		User:   "edvise::platform::user",
		School: "edvise::platform::school",
	},
	Education: EducationConstant{
		Level:      "edvise::education::level",
		Course:     "edvise::education::course",
		Class:      "edvise::education::class",
		Schedule:   "edvise::education::schedule",
		Assessment: "edvise::education::assessment",
		Score:      "edvise::education::score",
		Attendance: "edvise::education::attendance",
		Comment:    "edvise::education::comment",
	},
}
