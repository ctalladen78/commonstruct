package iteauth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//GetAccessToken make an http request to iteacloud platform API to get access token
func GetAccessToken(request AccessTokenRequest) (*AccessTokenResponse, error) {
	url := "https://dev-api.iteacloud.com/external-client-auth"

	if os.Getenv("STAGE") == "prod" {
		url = "https://api.iteacloud.com/external-client-auth"
	}
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	accessToken := AccessTokenResponse{}
	if err := json.Unmarshal(respBody, &accessToken); err != nil {
		return nil, err
	}
	return &accessToken, nil
}
