package edviseperm

//PermissionConstant hold the permission in the higher level
type PermissionConstant struct {
	Platform PlatformConstant `json:"platform"`
}

//PlatformConstant hold the permission strings in platform
type PlatformConstant struct {
	User   string `json:"user"`
	School string `json:"school"`
	Course string `json:"course"`
}

//Permission predefined permission
var Permission = PermissionConstant{
	Platform: PlatformConstant{
		User:   "edvise::platform::user",
		School: "edvise::platform::school",
		Course: "edvise::platform::course",
	},
}
