package cogth

import (
	"fmt"
	"testing"
	"time"
)

func TestGetOpenID(test *testing.T) {
	token, err := GetOpenID("arn::itea::1::platform::client::1", time.Hour*1, "iteacloud-platform-api-auth-dev-ap-southeast-1.iteacloud.com", "ap-southeast-1:0712133c-1bab-4796-9c1b-b9b23b64ec62", "ap-southeast-1")
	if err != nil || token == nil {
		test.Fail()
	}
	fmt.Println(*token)
}
