package es

import (
	"encoding/base64"
	"encoding/json"

	"github.com/kkesley/commonstruct/stdreq"
)

//NextToken holds the information to create a token for pagination
type NextToken struct {
	SortBy      []stdreq.Sort `json:"sort_key"`
	LatestValue []interface{} `json:"sort_values"`
}

//CreateNextToken create a base64 representation of pagination token
func (token NextToken) CreateNextToken() (string, error) {
	tokenByte, err := json.Marshal(token)
	if err != nil {
		return "", err
	}
	tokenB64 := base64.StdEncoding.EncodeToString(tokenByte)
	return tokenB64, nil
}

//DecodeNextToken decode base64 token to NextToken
func DecodeNextToken(b64 string) (*NextToken, error) {
	tokenByte, err := base64.URLEncoding.DecodeString(b64)
	if err != nil {
		return nil, err
	}
	var token NextToken
	if err := json.Unmarshal(tokenByte, &token); err != nil {
		return nil, err
	}
	return &token, nil
}
