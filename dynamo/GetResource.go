package dynamo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//GetResource get a single resource
func GetResource(svc *dynamodb.DynamoDB, tableName string, key map[string]*dynamodb.AttributeValue) error {

	input := &dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String(tableName),
	}
	if _, err := svc.GetItem(input); err != nil {
		return err
	}
	return nil
}
