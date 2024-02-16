package sms

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"net/url"
	"strings"
)

type SmsSender struct {
	ApiKey     string   `json:"api_key"`
	ApiUser    string   `json:"api_user"`
	Recipients []string `json:"recipients"`
	Message    string   `json:"message"`
	Sender     string   `json:"sender"`
	SmsKey     string   `json:"sms_key"`
}

type Recipient struct {
	Key         string `json:"key"`
	Cost        string `json:"cost"`
	SmsKey      string `json:"sms_key"`
	MessageId   string `json:"message_id"`
	MessagePart int    `json:"message_part"`
	Number      string `json:"number"`
	Status      string `json:"status"`
	StatusCode  string `json:"status_code"`
}

type SmsMessageData struct {
	Message    string      `json:"message"`
	Cost       string      `json:"cost"`
	Recipients []Recipient `json:"recipients"`
}

type ErrorResponse struct {
	HasError bool   `json:"has_error"`
	Message  string `json:"message"`
}

type SmsSenderResponse struct {
	ErrorResponse  ErrorResponse  `json:"error_response"`
	SmsMessageData SmsMessageData `json:"sms_message_data"`
}

// SendSMS sends an SMS using the Africa's Talking API
func (s *SmsSender) SendSMS() (SmsSenderResponse, error) {
	endpoint := "https://api.africastalking.com/version1/messaging"
	parsedURL, err := url.Parse(endpoint)
	if err != nil {
		return SmsSenderResponse{}, err
	}

	body := map[string][]string{
		"username": {s.ApiUser},
		"to":       s.Recipients,
		"message":  {s.Message},
		"from":     {s.Sender},
	}

	form := url.Values{}
	for key, values := range body {
		for _, value := range values {
			form.Add(key, value)
		}
	}

	req, err := http.NewRequest(http.MethodPost, parsedURL.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return SmsSenderResponse{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("apiKey", s.ApiKey)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return SmsSenderResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusCreated {
		recipients := make([]Recipient, 0)

		var data map[string]interface{}
		err = json.NewDecoder(res.Body).Decode(&data)
		if err != nil {
			return SmsSenderResponse{}, err
		}

		smsMessageData := data["SMSMessageData"].(map[string]interface{})
		message := smsMessageData["Message"].(string)
		cost := strings.Split(message, " ")[len(message)-1]
		recipientsData := smsMessageData["Recipients"].([]interface{})

		for _, recipient := range recipientsData {
			recipientData := recipient.(map[string]interface{})

			rct := Recipient{
				Key:         uuid.New().String(),
				Cost:        recipientData["cost"].(string),
				SmsKey:      s.SmsKey,
				MessageId:   recipientData["messageId"].(string),
				MessagePart: int(recipientData["messageParts"].(float64)),
				Number:      recipientData["number"].(string),
				Status:      recipientData["status"].(string),
				StatusCode:  fmt.Sprintf("%v", recipientData["statusCode"]),
			}

			recipients = append(recipients, rct)
		}

		smsSenderResponse := SmsSenderResponse{
			ErrorResponse: ErrorResponse{
				HasError: false,
			},
			SmsMessageData: SmsMessageData{
				Message:    message,
				Cost:       cost,
				Recipients: recipients,
			},
		}

		return smsSenderResponse, nil
	}

	smsSenderResponse := SmsSenderResponse{
		ErrorResponse: ErrorResponse{
			HasError: true,
			Message:  "Message not sent",
		},
	}

	return smsSenderResponse, fmt.Errorf("status code: %d", res.StatusCode)
}
