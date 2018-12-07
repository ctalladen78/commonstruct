package checker

//PermissionListQuery build logic of permission query for list operations
func (permission CheckResponse) PermissionListQuery(filterOutput FilterIDsOutput) ([]string, []string) {
	var allows []string
	var denies []string
	if permission.Allow.All { //if permission has allow all
		if len(permission.Deny.Resources) > 0 {
			denies = permission.Deny.Resources
		}

	} else if permission.Deny.All { //if permission has deny all
		if len(permission.Allow.Resources) > 0 {
			allows = permission.Allow.Resources
		}
	} else { //if permission has individual deny and allow
		if len(permission.Deny.Resources) > 0 {
			denies = permission.Deny.Resources
		}
		if len(permission.Allow.Resources) > 0 {
			allows = permission.Allow.Resources
		}
		if len(permission.Deny.Resources) <= 0 && len(permission.Allow.Resources) <= 0 {
			allows = []string{"-1"}
		}
	}
	return allows, denies
}
