package checker

import "github.com/kkesley/commonstruct/auth/perm"

//CheckResponse holds the response of checking resource
type CheckResponse struct {
	Allow perm.ResourceXpression `json:"allow"`
	Deny  perm.ResourceXpression `json:"deny"`
}
