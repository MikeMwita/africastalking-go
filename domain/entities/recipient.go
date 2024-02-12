package entities

// Recipient represents a recipient of an SMS message
type Recipient struct {
	Key         string `json:"key"`         // the unique identifier of the recipient
	Cost        string `json:"cost"`        // the cost of sending the message to the recipient
	SMSKey      string `json:"smsKey"`      // the foreign key to the SMS sender
	MessageID   string `json:"messageId"`   // the message ID from the AfricasTalking API
	MessagePart int    `json:"messagePart"` // the number of message parts
	Number      string `json:"number"`      // the phone number of the recipient
	Status      string `json:"status"`      // the status of the message delivery
	StatusCode  int    `json:"statusCode"`  // the status code from the AfricasTalking API
}
