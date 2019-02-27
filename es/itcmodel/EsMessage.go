package itcmodel

import "time"

//EsMessage holds the message in elasticsearch
type EsMessage struct {
	ARN       *string `json:"arn,omitempty"`
	ClientARN *string `json:"client_arn,omitempty"`
	Code      *string `json:"code,omitempty"`

	Targets []string `json:"targets,omitempty"` //target user arns

	//the name of the user who sent this message.
	//this is the name 'when' the user sent this message.
	//the actual name may differ as user can change their name.
	//this field will be used in the notification
	From    *string            `json:"from,omitempty"`
	Message *string            `json:"message,omitempty"`
	Data    []MessageDataValue `json:"data,omitempty"`

	ServiceScopes []string `json:"service_scopes,omitempty"` //used to filter the result when listing (e.g. edvise/platform)

	//this is used to identify which part of the application is this resource belongs to
	//e.g. class discussion in edvise would have edvise.class_discussion
	HandleKey *string `json:"handle_key,omitempty"`

	//this is used to identify the ID of the handle key.
	//e.g. the code of the class for class discussion
	HandleID  *string    `json:"handle_id,omitempty"`
	Principal *string    `json:"principal,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

//MessageDataValue represent a data key
type MessageDataValue struct {
	Key   *string `json:"key"`
	Value *string `json:"value"`
}
