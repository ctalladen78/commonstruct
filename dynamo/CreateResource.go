package dynamo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//CreateResource is a function to create a generic resource in dynamodb
//Used by any resource with client_arn and code as the key
func CreateResource(svc *dynamodb.DynamoDB, tableName string, attributes map[string]*dynamodb.AttributeValue) error {
	input := dynamodb.PutItemInput{
		TableName:           aws.String(tableName),
		Item:                attributes,
		ConditionExpression: aws.String("attribute_not_exists(client_arn) AND attribute_not_exists(code)"),
	}

	if _, err := svc.PutItem(&input); err != nil {
		return err
	}
	return nil
}
