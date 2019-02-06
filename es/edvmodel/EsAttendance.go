package edvmodel

import "time"

//EsAttendance holds the attendance in elasticsearch
type EsAttendance struct {
	ARN       *string    `json:"arn,omitempty"`
	ClientARN *string    `json:"client_arn,omitempty"`
	Code      *string    `json:"code,omitempty"`
	Title     *string    `json:"title,omitempty"`
	Date      *time.Time `json:"date,omitempty"`

	Users []AttendanceUser `json:"users,omitempty"`

	Class   *string  `json:"class,omitempty"`   //If it's not tied to class, make this empty
	Targets []string `json:"targets,omitempty"` //e.g. course [COURSE/asodijaosdji] or events [SCHEDULE/asdoasijdoasd]

	Principal *string    `json:"principal,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	UpdatedBy *string    `json:"updated_by,omitempty"`
}

//AttendanceUser hold attendance information for each user
type AttendanceUser struct {
	Username    *string  `json:"username,omitempty"`
	IsAttend    *bool    `json:"is_attend,omitempty"`
	Status      *string  `json:"status,omitempty"`
	Comment     *string  `json:"comment,omitempty"`
	Attachments []string `json:"attachments,omitempty"`
}
