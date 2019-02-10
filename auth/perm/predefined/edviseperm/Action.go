package edviseperm

//ActionConstant holds constant in higher level
type ActionConstant struct {
	Platform  PlatformAction  `json:"platform"`
	Education EducationAction `json:"education"`
	Basic     BasicAction     `json:"general"`
}

//PlatformAction holds action of platform
type PlatformAction struct {
	User   UserAction   `json:"user"`
	School SchoolAction `json:"school"`
}

//EducationAction holds action of education
type EducationAction struct {
	Level      LevelAction      `json:"level"`
	Course     CourseAction     `json:"course"`
	Class      ClassAction      `json:"class"`
	Schedule   ScheduleAction   `json:"schedule"`
	Assessment AssessmentAction `json:"assessment"`
	Score      ScoreAction      `json:"score"`
	Attendance AttendanceAction `json:"attendance"`
	Comment    CommentAction    `json:"comment"`
}

//UserAction holds action of user
type UserAction struct {
	Read   string `json:"read"`
	Write  string `json:"write"`
	Delete string `json:"delete"`
	Create string `json:"create"`
}

//SchoolAction holds action of user
type SchoolAction struct {
	Read  string `json:"read"`
	Write string `json:"write"`
}

//LevelAction holds action for levels
type LevelAction struct {
	Read   string `json:"read"`
	Write  string `json:"write"`
	Delete string `json:"delete"`
	Create string `json:"create"`
}

//CourseAction holds action for courses
type CourseAction struct {
	Read   string `json:"read"`
	Write  string `json:"write"`
	Delete string `json:"delete"`
	Create string `json:"create"`
}

//ClassAction holds action for classes
type ClassAction struct {
	Read   string `json:"read"`
	Write  string `json:"write"`
	Delete string `json:"delete"`
	Create string `json:"create"`
}

//ScheduleAction holds action for schedules
type ScheduleAction struct {
	Read   string `json:"read"`
	Write  string `json:"write"`
	Delete string `json:"delete"`
	Create string `json:"create"`
}

//AssessmentAction holds action for assessments
type AssessmentAction struct {
	Read   string `json:"read"`
	Write  string `json:"write"`
	Delete string `json:"delete"`
	Create string `json:"create"`
}

//ScoreAction holds action for scores
type ScoreAction struct {
	Read   string `json:"read"`
	Write  string `json:"write"`
	Delete string `json:"delete"`
	Create string `json:"create"`
}

//AttendanceAction holds action for attendance
type AttendanceAction struct {
	Read   string `json:"read"`
	Write  string `json:"write"`
	Delete string `json:"delete"`
	Create string `json:"create"`
}

//CommentAction holds action for comment
type CommentAction struct {
	ReadStudentComment        string `json:"read:student_comment"`
	WriteStudentComment       string `json:"write:student_comment"`
	ReadAssessmentDiscussion  string `json:"read:assessment_discussion"`
	WriteAssessmentDiscussion string `json:"write:assessment_discussion"`
}

//BasicAction holds basic action
type BasicAction struct {
	Read     string         `json:"read"`
	Write    string         `json:"write"`
	Delete   string         `json:"delete"`
	Create   string         `json:"create"`
	Fragment FragmentAction `json:"fragment"`
}

//Action holds the action variable
var Action = ActionConstant{
	Basic: BasicAction{
		Read:   "read",
		Write:  "write",
		Delete: "delete",
		Create: "create",
		Fragment: FragmentAction{
			Edvise: EdviseFragment{
				StudentComment:       "/student-comment",
				AssessmentDiscussion: "/assessment-discussion",
			},
		},
	},
	Platform: PlatformAction{
		User: UserAction{
			Read:   Permission.Platform.User + "::" + "read",
			Write:  Permission.Platform.User + "::" + "write",
			Delete: Permission.Platform.User + "::" + "delete",
			Create: Permission.Platform.User + "::" + "create",
		},
		School: SchoolAction{
			Read:  Permission.Platform.School + "::" + "read",
			Write: Permission.Platform.School + "::" + "write",
		},
	},
	Education: EducationAction{
		Level: LevelAction{
			Read:   Permission.Education.Level + "::" + "read",
			Write:  Permission.Education.Level + "::" + "write",
			Delete: Permission.Education.Level + "::" + "delete",
			Create: Permission.Education.Level + "::" + "create",
		},
		Course: CourseAction{
			Read:   Permission.Education.Course + "::" + "read",
			Write:  Permission.Education.Course + "::" + "write",
			Delete: Permission.Education.Course + "::" + "delete",
			Create: Permission.Education.Course + "::" + "create",
		},
		Class: ClassAction{
			Read:   Permission.Education.Class + "::" + "read",
			Write:  Permission.Education.Class + "::" + "write",
			Delete: Permission.Education.Class + "::" + "delete",
			Create: Permission.Education.Class + "::" + "create",
		},
		Schedule: ScheduleAction{
			Read:   Permission.Education.Schedule + "::" + "read",
			Write:  Permission.Education.Schedule + "::" + "write",
			Delete: Permission.Education.Schedule + "::" + "delete",
			Create: Permission.Education.Schedule + "::" + "create",
		},
		Assessment: AssessmentAction{
			Read:   Permission.Education.Assessment + "::" + "read",
			Write:  Permission.Education.Assessment + "::" + "write",
			Delete: Permission.Education.Assessment + "::" + "delete",
			Create: Permission.Education.Assessment + "::" + "create",
		},
		Score: ScoreAction{
			Read:   Permission.Education.Score + "::" + "read",
			Write:  Permission.Education.Score + "::" + "write",
			Delete: Permission.Education.Score + "::" + "delete",
			Create: Permission.Education.Score + "::" + "create",
		},
		Attendance: AttendanceAction{
			Read:   Permission.Education.Attendance + "::" + "read",
			Write:  Permission.Education.Attendance + "::" + "write",
			Delete: Permission.Education.Attendance + "::" + "delete",
			Create: Permission.Education.Attendance + "::" + "create",
		},
		Comment: CommentAction{
			ReadStudentComment:        Permission.Education.Comment + "::" + "read/student-comment",
			WriteStudentComment:       Permission.Education.Comment + "::" + "write/student-comment",
			ReadAssessmentDiscussion:  Permission.Education.Comment + "::" + "read/assessment-discussion",
			WriteAssessmentDiscussion: Permission.Education.Comment + "::" + "write/assessment-discussion",
		},
	},
}
