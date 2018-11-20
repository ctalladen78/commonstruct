package edviseperm

//ActionConstant holds constant in higher level
type ActionConstant struct {
	Platform PlatformAction `json:"platform"`
	Basic    BasicAction    `json:"general"`
}

//PlatformAction holds action of platform
type PlatformAction struct {
	User       UserAction       `json:"user"`
	School     SchoolAction     `json:"school"`
	Course     CourseAction     `json:"course"`
	Class      ClassAction      `json:"class"`
	Schedule   ScheduleAction   `json:"schedule"`
	Assessment AssessmentAction `json:"assessment"`
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

//CourseAction holds action for courses
type CourseAction struct {
	Read   string `json:"read"`
	Write  string `json:"write"`
	Delete string `json:"delete"`
	Create string `json:"create"`
}

//ClassAction holds action for courses
type ClassAction struct {
	Read   string `json:"read"`
	Write  string `json:"write"`
	Delete string `json:"delete"`
	Create string `json:"create"`
}

//ScheduleAction holds action for courses
type ScheduleAction struct {
	Read   string `json:"read"`
	Write  string `json:"write"`
	Delete string `json:"delete"`
	Create string `json:"create"`
}

//AssessmentAction holds action for courses
type AssessmentAction struct {
	Read   string `json:"read"`
	Write  string `json:"write"`
	Delete string `json:"delete"`
	Create string `json:"create"`
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
		Course: CourseAction{
			Read:   Permission.Platform.Course + "::" + "read",
			Write:  Permission.Platform.Course + "::" + "write",
			Delete: Permission.Platform.Course + "::" + "delete",
			Create: Permission.Platform.Course + "::" + "create",
		},
		Class: ClassAction{
			Read:   Permission.Platform.Class + "::" + "read",
			Write:  Permission.Platform.Class + "::" + "write",
			Delete: Permission.Platform.Class + "::" + "delete",
			Create: Permission.Platform.Class + "::" + "create",
		},
		Schedule: ScheduleAction{
			Read:   Permission.Platform.Schedule + "::" + "read",
			Write:  Permission.Platform.Schedule + "::" + "write",
			Delete: Permission.Platform.Schedule + "::" + "delete",
			Create: Permission.Platform.Schedule + "::" + "create",
		},
		Assessment: AssessmentAction{
			Read:   Permission.Platform.Assessment + "::" + "read",
			Write:  Permission.Platform.Assessment + "::" + "write",
			Delete: Permission.Platform.Assessment + "::" + "delete",
			Create: Permission.Platform.Assessment + "::" + "create",
		},
	},
}
