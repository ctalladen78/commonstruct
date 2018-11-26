package iteauth

//AccessTokenResponse holds the response for access token
type AccessTokenResponse struct {
	UserARN           string `json:"user_arn"`
	ClientARN         string `json:"client_arn"`
	ClientName        string `json:"client_name"`
	AccessToken       string `json:"access_token,omitempty" log:"false"`
	AccessTokenExpiry int64  `json:"access_token_expiry"`
}
