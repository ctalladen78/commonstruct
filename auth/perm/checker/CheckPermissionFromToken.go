package checker

import (
	"os"

	"github.com/kkesley/commonstruct/arn"
)

//CheckPermissionFromToken returns permission expression from token (convenient but no customization)
func CheckPermissionFromToken(request CheckPermissionFromTokenRequest) (*CheckResponse, error) {
	role := ""
	if len(request.Token.RoleARN) > 0 {
		role = request.Token.RoleARN[0]
	}
	return CheckPermission(CheckPermissionRequest{
		Bucket:   os.Getenv("ROLE_OUTPUT_BUCKET"),
		Region:   os.Getenv("ROLE_OUTPUT_REGION"),
		Role:     role,
		ClientID: arn.GetResourceID(request.Token.ClientARN),
		Path:     request.Path,
		Actions:  request.BasicActions,
		Token:    request.Token,
	})
}
