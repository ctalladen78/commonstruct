package notification

//Platform notification
type Platform struct {
	GCM         *GCM  `json:"GCM,omitempty"`
	APNS        *APNS `json:"APNS,omitempty"`
	APNSSandBox *APNS `json:"APNS_SANDBOX,omitempty"`
}

//GCM payload
type GCM struct {
	Notification *GCMNotification       `json:"notification,omitempty"`
	Data         map[string]interface{} `json:"data,omitempty"`
}

//GCMNotification notification payload
type GCMNotification struct {
	Title *string `json:"title,omitempty"`
	Text  *string `json:"text,omitempty"`
}

//APNS payload
type APNS struct {
	Notification *APNSNotification      `json:"aps,omitempty"`
	Data         map[string]interface{} `json:"data,omitempty"`
}

//APNSNotification aps payload
type APNSNotification struct {
	Alert *APNSAlert `json:"alert,omitempty"`
}

//APNSAlert alert payload
type APNSAlert struct {
	Title *string `json:"title,omitempty"`
	Body  *string `json:"body,omitempty"`
}
