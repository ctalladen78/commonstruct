package itcmodel

import (
	"time"
)

//EsActivity holds the user activity in the portal
type EsActivity struct {
	ClientARN   *string         `json:"client_arn,omitempty"`
	Session     *string         `json:"session,omitempty"`
	Action      *string         `json:"action,omitempty"`
	TargetType  *string         `json:"target_type,omitempty"`  //CLASS_TO_STUDENT/STUDENT
	TargetScope *string         `json:"target_scope,omitempty"` //EDVISE/PLATFORM
	TargetOld   []ActivityValue `json:"target_old,omitempty"`   //FOR UPDATING/DELETING
	TargetNew   []ActivityValue `json:"target_new,omitempty"`
	Subjects    []ActivityValue `json:"subjects,omitempty"` //OPTIONAL FIELD. IF THE ACTION HAS A DIRECT SUBJECT (E.G. SCORING STUDENTS, SUBJECTS WILL BE STUDENT ARN, TARGET WILL BE THE SCORE)

	Principal      *string    `json:"principal,omitempty"`
	PrincipalScope *string    `json:"principal_scope,omitempty"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
}

//ActivityValue represent a resource
type ActivityValue struct {
	Key   *string `json:"key"`
	Value *string `json:"value"`
}
