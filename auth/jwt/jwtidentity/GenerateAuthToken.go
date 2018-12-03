package jwtidentity

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
)

//GenerateAuthToken generate from map to auth token
func GenerateAuthToken(authContext map[string]interface{}) TokenRequest {
	authToken := TokenRequest{}
	iterateAuthContext([]string{"IsRoot", "UserARN", "RoleARN", "ClientID", "ClientARN", "ClientName", "FirstName", "LastName", "Username", "Groups", "Permissions", "Device", "IsExternal", "Service"}, &authToken, authContext)
	return authToken
}

func iterateAuthContext(fields []string, token *TokenRequest, context map[string]interface{}) {
	for _, field := range fields {
		if context[field] != nil {
			fieldType := reflect.TypeOf(context[field]).String()
			switch field {
			case "IsRoot":
				if fieldType == "string" {
					token.IsRoot, _ = strconv.ParseBool(context[field].(string))
				} else if fieldType == "bool" {
					token.IsRoot = context[field].(bool)
				}

			case "UserARN":
				if fieldType != "string" {
					continue
				}
				token.UserARN = context[field].(string)
			case "RoleARN":
				if fieldType == "string" {
					token.RoleARN = strings.Split(context[field].(string), ",")
				} else if fieldType == "[]interface {}" {
					for _, value := range context[field].([]interface{}) {
						if reflect.TypeOf(value).String() == "string" {
							token.RoleARN = append(token.RoleARN, value.(string))
						}
					}
				}
			case "ClientARN":
				if fieldType != "string" {
					continue
				}
				token.ClientARN = context[field].(string)
			case "ClientName":
				if fieldType != "string" {
					continue
				}
				token.ClientName = context[field].(string)
			case "FirstName":
				if fieldType != "string" {
					continue
				}
				token.FirstName = aws.String(context[field].(string))
			case "LastName":
				if fieldType != "string" {
					continue
				}
				token.LastName = aws.String(context[field].(string))
			case "Username":
				if fieldType != "string" {
					continue
				}
				token.Username = context[field].(string)
			case "Groups":
				if fieldType == "string" {
					token.Groups = strings.Split(context[field].(string), ",")
				} else if fieldType == "[]interface {}" {
					for _, value := range context[field].([]interface{}) {
						if reflect.TypeOf(value).String() == "string" {
							token.Groups = append(token.Groups, value.(string))
						}
					}
				}

			case "Device":
				if fieldType != "string" {
					continue
				}
				token.Device = context[field].(string)

			case "IsExternal":
				if fieldType == "string" {
					token.IsExternal, _ = strconv.ParseBool(context[field].(string))
				} else if fieldType == "bool" {
					token.IsExternal = context[field].(bool)
				}

			case "Service":
				if fieldType != "string" {
					continue
				}
				token.Service = context[field].(string)
			}

		}
	}
}
