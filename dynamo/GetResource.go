package dynamo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//GetResource get a single resource
func GetResource(svc *dynamodb.DynamoDB, tableName string, key map[string]*dynamodb.AttributeValue) (map[string]*dynamodb.AttributeValue, error) {
	input := &dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String(tableName),
	}
	output, err := svc.GetItem(input)
	if err != nil || output == nil || output.Item == nil {
		return nil, err
	}
	return output.Item, nil
}
