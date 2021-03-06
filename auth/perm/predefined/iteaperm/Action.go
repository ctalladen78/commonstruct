package iteaperm

//ActionConstant holds constant in higher level
type ActionConstant struct {
	Platform PlatformAction `json:"platform"`
	Basic    BasicAction    `json:"general"`
}

//PlatformAction holds action of platform
type PlatformAction struct {
	User           UserAction           `json:"user"`
	Role           RoleAction           `json:"role"`
	PasswordPolicy PasswordPolicyAction `json:"password_policy"`
	Billing        BillingAction        `json:"billing"`
	Comment        CommentAction        `json:"comment"`
	ActivityLog    ActivityLogAction    `json:"activity_log"`
}

//UserAction holds action of user
type UserAction struct {
	Read   string `json:"read"`
	Write  string `json:"write"`
	Delete string `json:"delete"`
	Create string `json:"create"`
}

//RoleAction holds action of role
type RoleAction struct {
	Read   string `json:"read"`
	Write  string `json:"write"`
	Delete string `json:"delete"`
	Create string `json:"create"`
}

//PasswordPolicyAction holds action of role
type PasswordPolicyAction struct {
	Read  string `json:"read"`
	Write string `json:"write"`
}

//ActivityLogAction holds action of activity-log
type ActivityLogAction struct {
	Read string `json:"read"`
}

//BillingAction holds action for billings
type BillingAction struct {
	Read string `json:"billing"`
}

//CommentAction holds action for comment
type CommentAction struct {
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
		Fragment: FragmentAction{
			Platform: PlatformFragment{
				PasswordPolicy: "/password-policy",
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
		Role: RoleAction{
			Read:   Permission.Platform.Role + "::" + "read",
			Write:  Permission.Platform.Role + "::" + "write",
			Delete: Permission.Platform.Role + "::" + "delete",
			Create: Permission.Platform.Role + "::" + "create",
		},
		PasswordPolicy: PasswordPolicyAction{
			Read:  Permission.Platform.User + "::" + "read/password-policy",
			Write: Permission.Platform.User + "::" + "write/password-policy",
		},
		Billing: BillingAction{
			Read: Permission.Platform.Billing + "::" + "read",
		},
		Comment: CommentAction{
			Read:  Permission.Platform.Comment + "::" + "read",
			Write: Permission.Platform.Comment + "::" + "write",
		},
		ActivityLog: ActivityLogAction{
			Read: Permission.Platform.ActivityLog + "::" + "read",
		},
	},
}
