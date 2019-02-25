package notification

//Payload of a message
type Payload struct {
	Title   string            `json:"title"`
	Message string            `json:"message"`
	Data    map[string]string `json:"data"`
}
