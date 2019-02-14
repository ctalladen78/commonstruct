package itcevent

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

//LogEventSNS log event sns
func LogEventSNS(sess *session.Session, topic string, eventType EventType) error {
	eventStr, err := json.Marshal(eventType)
	if err != nil {
		return err
	}
	return FireSNSWithAttribute(topic, string(eventStr), map[string]string{"event": eventType.Event})
}

//LogEventSNSDefault sends sns without session
func LogEventSNSDefault(region string, topic string, eventType EventType) error {
	config := aws.Config{
		Region: aws.String(region),
	}
	sess := session.Must(session.NewSession(&config))

	return LogEventSNS(sess, topic, eventType)
}
