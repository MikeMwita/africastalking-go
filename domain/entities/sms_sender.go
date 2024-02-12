package entities

type SMSSender struct {
	Key        string   `json:"key"`        // the unique identifier of the sender
	APIKey     string   `json:"apiKey"`     // the API key from the AfricasTalking API
	APIUser    string   `json:"apiUser"`    // the username from the AfricasTalking API
	Recipients []string `json:"recipients"` // the list of phone numbers to send the message to
	Message    string   `json:"message"`    // the message content
	Sender     string   `json:"sender"`     // the sender ID
}

type SMSMessageData struct {
	Message    string
	Cost       string
	Recipients []*Recipient
}
