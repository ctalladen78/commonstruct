package notification

//Platform notification
type Platform struct {
	Default *string `json:"default,omitempty"`
	GCM     *GCM    `json:"-"`
	APNS    *APNS   `json:"-"`

	GCMStr         *string `json:"GCM,omitempty"`
	APNSStr        *string `json:"APNS,omitempty"`
	APNSSandBoxStr *string `json:"APNS_SANDBOX,omitempty"`
}

//GCM payload
type GCM struct {
	Notification *GCMNotification  `json:"notification,omitempty"`
	Config       *AndroidConfig    `json:"android,omitempty"`
	Data         map[string]string `json:"data,omitempty"`
}

//AndroidConfig represents the config of GCM for android
type AndroidConfig struct {
	CollapseKey *string `json:"collapse_key"`
}

//GCMNotification notification payload
type GCMNotification struct {
	Title *string `json:"title,omitempty"`
	Text  *string `json:"text,omitempty"`
	Tag   *string `json:"tag,omitempty"`
}

//APNS payload
type APNS struct {
	Notification *APNSNotification `json:"aps,omitempty"`
	Data         map[string]string `json:"data,omitempty"`
}

//APNSNotification aps payload
type APNSNotification struct {
	Alert    *APNSAlert `json:"alert,omitempty"`
	ThreadID *string    `json:"thread-id,omitempty"`
}

//APNSAlert alert payload
type APNSAlert struct {
	Title *string `json:"title,omitempty"`
	Body  *string `json:"body,omitempty"`
}
