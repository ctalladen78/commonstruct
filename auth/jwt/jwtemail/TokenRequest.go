package jwtemail

//TokenRequest request object for generating JWT
type TokenRequest struct {
	Email        string `json:"Email,omitempty"`
	ClientARN    string `json:"ClientARN,omitempty"`
	Username     string `json:"Username,omitempty"`
	ClientPrefix string `json:"ClientPrefix,omitempty"`
	IsRoot       bool   `json:"IsRoot"`
}
