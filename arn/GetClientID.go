package arn

import "strings"

//GetClientID get client ID for a resource
func GetClientID(resourceARN string) string {
	sections := strings.Split(resourceARN, "::")
	if len(sections) < 3 {
		return "" //invalid arn
	}
	if sections[2] == "" {
		//must be client
		return GetResourceID(resourceARN)
	}
	return sections[2]
}
