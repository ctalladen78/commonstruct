package notification

import (
	"encoding/json"
	"time"

	"github.com/aws/aws-sdk-go/aws"
)

//Notification holds notification message
type Notification struct {
	TargetUsers []string   `json:"target_users"`
	DeviceType  *string    `json:"device_type"`
	Payload     Payload    `json:"payload"`
	SendTime    *time.Time `json:"send_time"`
	ExternalID  *string    `json:"external_id"` //this id is used for scheduling. Must be unique accross the system
	Disabled    *bool      `json:"disabled"`    //if not enabled, ignore this notification
}

//Marshal to string
func (n Notification) Marshal() (*string, error) {
	nByte, err := json.Marshal(n)
	if err != nil {
		return nil, err
	}
	return aws.String(string(nByte)), nil
}
