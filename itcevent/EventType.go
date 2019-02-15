package itcevent

import (
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/kkesley/random"
)

//EventType holds the information regarding an event
type EventType struct {
	Code      *string                `json:"code"`
	Event     string                 `json:"event"`
	ClientARN string                 `json:"client_arn"`
	Principal string                 `json:"principal"`
	Target    Target                 `json:"target"`
	Detail    map[string]interface{} `json:"detail"`
	CreatedAt time.Time
	TTL       *int `json:"ttl,omitempty"`
}

//Target target of an event
type Target struct {
	ResourceGroup string `json:"resource_group"`
	ResourceARN   string `json:"resource_arn"`
}

//GenerateRandomCode returns a random code for this resource
func (e EventType) GenerateRandomCode() *string {
	return aws.String(strconv.FormatInt(time.Now().UnixNano(), 10) + "-" + random.MakeIDCustomCharSet(6, "ABCDEFGHJKLMNPQRSTUVWXYZ123456789"))
}
