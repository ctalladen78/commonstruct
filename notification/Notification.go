package notification

import "time"

//Notification holds notification message
type Notification struct {
	TargetUsers []string   `json:"target_users"`
	DeviceType  *string    `json:"device_type"`
	Message     *string    `json:"message"`
	SendTime    *time.Time `json:"send_time"`
}
