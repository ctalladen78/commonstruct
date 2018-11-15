package edviseperm

//ActionConstant holds constant in higher level
type ActionConstant struct {
	Platform PlatformAction `json:"platform"`
	Basic    BasicAction    `json:"general"`
}

//PlatformAction holds action of platform
type PlatformAction struct {
	User   UserAction   `json:"user"`
	School SchoolAction `json:"school"`
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
	},
}
