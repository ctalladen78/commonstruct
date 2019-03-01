package notification

//Payload of a message
type Payload struct {
	Title      *string           `json:"title,omitempty"`
	Message    *string           `json:"message,omitempty"`
	Data       map[string]string `json:"data"`
	GroupingID *string           `json:"grouping_id"`
	ID         *string           `json:"id"`
}

//ConvertToNotification converts payload to a notification object
func (p Payload) ConvertToNotification() Platform {
	return Platform{
		Default: p.Title,
		GCM: &GCM{
			Notification: &GCMNotification{
				Title: p.Title,
				Text:  p.Message,
				Tag:   p.ID,
			},
			Config: &AndroidConfig{
				CollapseKey: p.GroupingID,
			},
			Data: p.Data,
		},
		APNS: &APNS{
			Notification: &APNSNotification{
				Alert: &APNSAlert{
					Title: p.Title,
					Body:  p.Message,
				},
				ThreadID: p.GroupingID,
			},
			Data: p.Data,
		},
	}
}
