package notification

//Payload of a message
type Payload struct {
	Title      *string           `json:"title,omitempty"`
	Message    *string           `json:"message,omitempty"`
	Data       map[string]string `json:"data"`
	GroupingID *string           `json:"grouping_id"`
}
