package dynamo

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//UpdateResource is a function to update a generic resource in dynamodb
//Used by Class, Course, Assessment, Schedule
func UpdateResource(svc *dynamodb.DynamoDB, tableName string, key map[string]*dynamodb.AttributeValue, attributes map[string]*dynamodb.AttributeValue) error {
	return UpdateRaw(svc, tableName, key, attributes, aws.String("attribute_exists(client_arn) AND attribute_exists(code)"))
}

//UpdateRaw updates a resource in dynamodb with custom condition
func UpdateRaw(svc *dynamodb.DynamoDB, tableName string, key map[string]*dynamodb.AttributeValue, attributes map[string]*dynamodb.AttributeValue, condition *string) error {
	updateStatements := make([]string, 0)
	newAttributes := make(map[string]*dynamodb.AttributeValue)
	newNames := make(map[string]*string)
	for key, val := range attributes {
		updateStatements = append(updateStatements, "#"+key+" = :"+key)
		newAttributes[":"+key] = val
		newNames["#"+key] = aws.String(key)
	}
	input := &dynamodb.UpdateItemInput{
		Key:                       key,
		TableName:                 aws.String(tableName),
		UpdateExpression:          aws.String("set " + strings.Join(updateStatements, ",")),
		ExpressionAttributeNames:  newNames,
		ExpressionAttributeValues: newAttributes,
		ReturnValues:              aws.String("UPDATED_NEW"),
		ConditionExpression:       condition,
	}
	if _, err := svc.UpdateItem(input); err != nil {
		return err
	}
	return nil
}
