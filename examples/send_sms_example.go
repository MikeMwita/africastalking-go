package sms_sender

import (
	"fmt"
	"github.com/MikeMwita/africastalking-go/pkg/sms"
	"log"
)

func main() {
	// Example usage
	sender := sms.SmsSender{
		ApiKey:     "your_api_key",
		ApiUser:    "your_api_user",
		Recipients: []string{"+1234567890"},
		Message:    "Hello, world!",
		Sender:     "your_sender",
	}

	response, err := sender.SendSMS()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Response: %+v\n", response)
}
