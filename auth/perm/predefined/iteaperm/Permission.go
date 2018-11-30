package iteaperm

//PermissionConstant hold the permission in the higher level
type PermissionConstant struct {
	Platform PlatformConstant `json:"platform"`
}

//PlatformConstant hold the permission strings in platform
type PlatformConstant struct {
	User    string `json:"user"`
	Role    string `json:"role"`
	Billing string `json:"billing"`
}

//Permission predefined permission
var Permission = PermissionConstant{
	Platform: PlatformConstant{
		User:    "itea::platform::user",
		Role:    "itea::platform::role",
		Billing: "itea::platform::billing",
	},
}
