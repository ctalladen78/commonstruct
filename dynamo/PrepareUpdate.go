package dynamo

import (
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

//PrepareUpdateOutput return general information for updating item in dynamodb
type PrepareUpdateOutput struct {
	TableName                 *string
	Key                       map[string]*dynamodb.AttributeValue
	UpdateExpression          *string
	ExpressionAttributeValues map[string]*dynamodb.AttributeValue
	ExpressionAttributeNames  map[string]*string
	ConditionExpression       *string
}

//PrepareUpdateTransaction returns an update statement for transaction
func PrepareUpdateTransaction(tableName string, key interface{}, attributes interface{}, conditionsKey string, conditionsValue interface{}) (*dynamodb.Update, error) {
	output, err := PrepareUpdate(tableName, key, attributes, conditionsKey, conditionsValue)
	if err != nil {
		return nil, err
	}
	return &dynamodb.Update{
		TableName:                 output.TableName,
		Key:                       output.Key,
		UpdateExpression:          output.UpdateExpression,
		ExpressionAttributeValues: output.ExpressionAttributeValues,
		ExpressionAttributeNames:  output.ExpressionAttributeNames,
		ConditionExpression:       output.ConditionExpression,
	}, nil
}

//PrepareUpdateItem returns an update statement for normal update operation
func PrepareUpdateItem(tableName string, key interface{}, attributes interface{}, conditionsKey string, conditionsValue interface{}) (*dynamodb.UpdateItemInput, error) {
	output, err := PrepareUpdate(tableName, key, attributes, conditionsKey, conditionsValue)
	if err != nil {
		return nil, err
	}
	return &dynamodb.UpdateItemInput{
		TableName:                 output.TableName,
		Key:                       output.Key,
		UpdateExpression:          output.UpdateExpression,
		ExpressionAttributeValues: output.ExpressionAttributeValues,
		ExpressionAttributeNames:  output.ExpressionAttributeNames,
		ConditionExpression:       output.ConditionExpression,
	}, nil
}

//PrepareUpdate return a generic response for updating item in dynamodb
func PrepareUpdate(tableName string, key interface{}, attributes interface{}, conditionsKey string, conditionsValue interface{}) (*PrepareUpdateOutput, error) {
	keyAttr, err := dynamodbattribute.MarshalMap(key)
	if err != nil {
		return nil, err
	}
	valAttr, err := dynamodbattribute.MarshalMap(attributes)
	if err != nil {
		return nil, err
	}

	updateStatements := make([]string, 0)
	newAttributes := make(map[string]*dynamodb.AttributeValue)
	newNames := make(map[string]*string)
	name := 0
	for key, val := range valAttr {
		updateStatements = append(updateStatements, "#val"+strconv.Itoa(name)+" = :val"+key)
		newAttributes[":val"+key] = val
		newNames["#val"+strconv.Itoa(name)] = aws.String(key)
		name++
	}

	conditionStatements := make([]string, 0)
	if len(conditionsKey) > 0 {
		conditionStatements = append(conditionStatements, conditionsKey)
	}
	if conditionsValue != nil {
		conAttr, err := dynamodbattribute.MarshalMap(conditionsValue)
		if err != nil {
			return nil, err
		}
		for key, val := range conAttr {
			conditionStatements = append(conditionStatements, "#con"+strconv.Itoa(name)+" = :con"+key)
			newAttributes[":con"+key] = val
			newNames["#con"+strconv.Itoa(name)] = aws.String(key)
			name++
		}
	}
	var conditionExpression *string
	if len(conditionStatements) > 0 {
		conditionExpression = aws.String(strings.Join(conditionStatements, " AND "))
	}

	return &PrepareUpdateOutput{
		TableName:                 aws.String(tableName),
		Key:                       keyAttr,
		UpdateExpression:          aws.String("set " + strings.Join(updateStatements, " ,")),
		ExpressionAttributeValues: newAttributes,
		ExpressionAttributeNames:  newNames,
		ConditionExpression:       conditionExpression,
	}, nil
}
