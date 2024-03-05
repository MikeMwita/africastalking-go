package sms

import (
	"testing"
)

func TestSendSMS(t *testing.T) {
	sender := SmsSender{
		ApiKey:     "3432e5e51e098ebc001db7c2544ff23504d9c2609c83ef4e23bdcea6a7cefd85",
		ApiUser:    "rangechem",
		Recipients: []string{"+254745617596"},
		Message:    "Hello Mike!",
		Sender:     "RANGECHEM",
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
