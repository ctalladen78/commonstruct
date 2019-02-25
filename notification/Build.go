package notification

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
)

//BuildNotification convert notification to string representative for each children
func (p *Platform) BuildNotification() error {
	gcmByte, err := json.Marshal(p.GCM)
	if err != nil {
		return err
	}
	apnsByte, err := json.Marshal(p.APNS)
	if err != nil {
		return err
	}
	p.GCMStr = aws.String(string(gcmByte))
	p.APNSStr = aws.String(string(apnsByte))
	p.APNSSandBoxStr = aws.String(string(apnsByte))
	return nil
}

//BuildNotificationString build the string representative of this notification
func (p *Platform) BuildNotificationString() (*string, error) {
	if err := p.BuildNotification(); err != nil {
		return nil, err
	}
	notificationByte, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return aws.String(string(notificationByte)), nil
}
