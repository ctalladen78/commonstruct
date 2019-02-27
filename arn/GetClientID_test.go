package arn

import (
	"testing"
)

func TestGetClientID(test *testing.T) {
	resourceARN := "arn::itea::1::property::hotel::3::room::4::window::5"
	if GetClientID(resourceARN) != "1" {
		test.Error("client id must be 1")
	}
}
