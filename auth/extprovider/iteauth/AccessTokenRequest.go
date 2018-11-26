package iteauth

//AccessTokenRequest holds the access token request
type AccessTokenRequest struct {
	ClientID      string  `json:"client_id"`
	ClientSecret  string  `json:"client_secret,omitempty" log:"false"`
	UserARN       string  `json:"user_arn"`
	UserFirstName *string `json:"user_first_name"`
	UserLastName  *string `json:"user_last_name"`
	Username      string  `json:"username"`
	ClientARN     string  `json:"client_arn"`
	ClientName    string  `json:"client_name"`
	Device        string  `json:"device"`
}
