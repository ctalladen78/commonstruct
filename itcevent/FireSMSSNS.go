package notifier

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/kkesley/commonstruct/itcmessaging"
)

//FireSMSSNS push sms sns notification
func FireSMSSNS(topic string, sms *itcmessaging.StandardSMSStructure) error {
	if sms == nil {
		return errors.New("sms cannot be nil")
	}
	sms.Method = "sms"
	conBytes, err := json.Marshal(sms)
	if err != nil {
		fmt.Println(err)
		return err
	}
	message := string(conBytes)
	return FireSNS(topic, message)
}
