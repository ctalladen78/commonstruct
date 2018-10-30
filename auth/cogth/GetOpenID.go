package cogth

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentity"
)

//GetOpenID get a cognito token to access aws resources (appsync, iot, etc)
func GetOpenID(userID string, duration time.Duration, cognitoAuthName string, cognitoIdentityID string, region string) (*string, error) {
	config := aws.Config{
		Region: aws.String(region),
	}
	sess := session.Must(session.NewSession(&config))
	svc := cognitoidentity.New(sess)

	output, err := svc.GetOpenIdTokenForDeveloperIdentity(&cognitoidentity.GetOpenIdTokenForDeveloperIdentityInput{
		IdentityPoolId: aws.String(cognitoIdentityID),
		Logins: map[string]*string{
			cognitoAuthName: aws.String(userID),
		},
		TokenDuration: aws.Int64(int64(duration.Seconds())),
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return output.Token, nil
}
