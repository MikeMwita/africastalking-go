package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MikeMwita/at/domain/entities"
	"github.com/MikeMwita/at/domain/interfaces"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	ErrInvalidInput   = errors.New("invalid input")
	ErrMessageNotSent = errors.New("message not sent")
	ErrDatabaseError  = errors.New("database error")
)

// SMSServiceImpl is the implementation of the SMSService interface
type SMSServiceImpl struct{}

// NewSMSService creates a new SMSService instance
func NewSMSService() interfaces.SMSService {
	return &SMSServiceImpl{}
}

// SendSMS sends an SMS message to the recipients and returns the response data
func (s *SMSServiceImpl) SendSMS(smsSender *entities.SMSSender) (*entities.SMSMessageData, error) {
	// Validate the input
	if smsSender == nil || smsSender.APIKey == "" || smsSender.APIUser == "" || len(smsSender.Recipients) == 0 || smsSender.Message == "" {
		return nil, ErrInvalidInput
	}

	// Create the request URL
	url := "https://api.africastalking.com/version1/messaging"

	// Create the request headers
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/x-www-form-urlencoded",
		"apiKey":       smsSender.APIKey,
	}

	// Create the request body
	body := url.Values{}
	body.Set("username", smsSender.APIUser)
	body.Set("to", strings.Join(smsSender.Recipients, ","))
	body.Set("message", smsSender.Message)
	if smsSender.Sender != "" {
		body.Set("from", smsSender.Sender)
	}

	// Create the HTTP request
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(body.Encode()))
	if err != nil {
		return nil, err
	}

	// Set the request headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// Create the HTTP client
	client := &http.Client{}

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Close the response body
	defer resp.Body.Close()

	// Read the response body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Check the response status code
	if resp.StatusCode != http.StatusCreated {
		return nil, ErrMessageNotSent
	}

	// Parse the response data
	var responseData map[string]map[string]interface{}
	err = json.Unmarshal(data, &responseData)
	if err != nil {
		return nil, err
	}

	// Extract the SMS message data
	smsMessageData, ok := responseData["SMSMessageData"]
	if !ok {
		return nil, ErrMessageNotSent
	}

	// Extract the message
	message, ok := smsMessageData["Message"].(string)
	if !ok {
		return nil, ErrMessageNotSent
	}

	// Extract the recipients
	recipientsData, ok := smsMessageData["Recipients"].([]interface{})
	if !ok {
		return nil, ErrMessageNotSent
	}

	// Convert the recipients data to entities
	recipients := make([]*entities.Recipient, len(recipientsData))
	for i, r := range recipientsData {
		recipientData, ok := r.(map[string]interface{})
		if !ok {
			return nil, ErrMessageNotSent
		}

		// Extract the recipient fields
		cost, ok := recipientData["cost"].(string)
		if !ok {
			return nil, ErrMessageNotSent
		}
		messageID, ok := recipientData["messageId"].(string)
		if !ok {
			return nil, ErrMessageNotSent
		}
		messagePart, ok := recipientData["messageParts"].(float64)
		if !ok {

		}
		number, ok := recipientData["number"].(string)
		if !ok {
			return nil, ErrMessageNotSent
		}
		status, ok := recipientData["status"].(string)
		if !ok {
			return nil, ErrMessageNotSent
		}
		statusCode, ok := recipientData["statusCode"].(float64)
		if !ok {
			return nil, ErrMessageNotSent
		}

		// Create the recipient entity
		recipient := &entities.Recipient{
			Key:         fmt.Sprintf("%s-%d", messageID, int(messagePart)),
			Cost:        cost,
			SMSKey:      smsSender.Key,
			MessageID:   messageID,
			MessagePart: int(messagePart),
			Number:      number,
			Status:      status,
			StatusCode:  int(statusCode),
		}

		// Add the recipient to the list
		recipients[i] = recipient
	}

	// Create the SMS message data entity
	smsMessageData := &entities.SMSMessageData{
		Message:    message,
		Cost:       message[strings.LastIndex(message, " ")+1:],
		Recipients: recipients,
	}

	// Return the SMS message data
	return smsMessageData, nil
}
