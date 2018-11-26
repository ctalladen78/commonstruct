package iteauth

import (
	"context"
	"net/http"
)

func GetAccessToken() {
	req, _ := http.NewRequest("GET", "https://mota.cf", nil)
	req = req.WithContext(context.Background())

	resp, _ := http.DefaultClient.Do(req)
}
