package sms

import (
	"testing"
)

func TestSendSMS(t *testing.T) {
	sender := SmsSender{
		ApiKey:     "",
		ApiUser:    "",
		Recipients: []string{"+254745617596"},
		Message:    "",
		Sender:     "",
	}

	response, err := sender.SendSMS()
	if err != nil {
		t.Errorf("error occurred: %v", err)
	}

	if response.ErrorResponse.HasError {
		t.Errorf("error response received: %s", response.ErrorResponse.Message)
	}

	if len(response.SmsMessageData.Recipients) == 0 {
		t.Errorf("no recipients received in response")
	}

	if response.SmsMessageData.Message == "" {
		t.Errorf("empty message received in response")
	}
}
