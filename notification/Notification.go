package notification

import "time"

//Notification holds notification message
type Notification struct {
	TargetUsers []string   `json:"target_users"`
	DeviceType  *string    `json:"device_type"`
	Payload     Payload    `json:"payload"`
	SendTime    *time.Time `json:"send_time"`
	ExternalID  *string    `json:"external_id"` //this id is used for scheduling. Must be unique accross the system
	Disabled    *bool      `json:"disabled"`    //if not enabled, ignore this notification
}
