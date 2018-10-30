package perm

import "github.com/kkesley/commonstruct/auth/jwt/jwtidentity"

//ParseActionsRequest holds the request for parsing actions
type ParseActionsRequest struct {
	PolicyStr string `json:"policy_str"`
	Token     jwtidentity.TokenRequest
}
