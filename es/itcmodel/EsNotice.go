package itcmodel

import (
	"time"
)

//EsNotice holds the notices
type EsNotice struct {
	ARN       *string `json:"arn,omitempty"`
	ClientARN *string `json:"client_arn,omitempty"`
	Code      *string `json:"code,omitempty"`

	//custom made filters for each application. e.g. user type.
	//The outer array is OR statement.
	//The inner array is AND statement.
	//Example: Notice for PARENT or STUDENT will be [][]string{[]string{"parent"}, []string{"student"}}
	//The filter can be anything depending of the application
	Filters []string `json:"filters,omitempty"`

	Targets     []string          `json:"targets,omitempty"` //target arns. can be any resource
	Title       *string           `json:"title,omitempty"`
	Description *string           `json:"description,omitempty"`
	Data        []NoticeDataValue `json:"data,omitempty"`

	ServiceScopes []string   `json:"service_scopes,omitempty"` //platform / edvise
	Principal     *string    `json:"principal,omitempty"`
	PublishDate   *time.Time `json:"publish_date,omitempty"` //if null, then publish now
	CreatedAt     *time.Time `json:"created_at,omitempty"`
}

//NoticeDataValue represent a data key
type NoticeDataValue struct {
	Key   *string `json:"key"`
	Value *string `json:"value"`
}
