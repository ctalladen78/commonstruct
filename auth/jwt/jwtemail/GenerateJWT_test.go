package jwtemail

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerateJWT(test *testing.T) {
	request := TokenRequest{
		Email: "kk@gmail.com",
	}
	jwt, err := GenerateJWT(request, time.Minute*1, "SuperSecretKeyOnlyForThisApplication")
	if err != nil {
		fmt.Println(err)
		test.Fail()
	}
	fmt.Println(jwt)
}

func TestGenerateJWTWithIssuer(test *testing.T) {
	request := TokenRequest{
		Email: "kk@gmail.com",
	}
	jwt, err := GenerateJWTWithIssuer(request, time.Minute*1, "SuperSecretKeyOnlyForThisApplication", "pss")
	if err != nil {
		fmt.Println(err)
		test.Fail()
	}
	fmt.Println(jwt)
}
