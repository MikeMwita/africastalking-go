package sms

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
)

func (s *SmsSender) SendSMS(ctx context.Context) (SmsSenderResponse, error) {
	form := url.Values{
		"username": {s.Client.apiUser},
		"to":       {strings.Join(s.Recipients, ",")},
		"message":  {s.Message},
		"from":     {s.Sender},
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.Client.apiURL, strings.NewReader(form.Encode()))
	if err != nil {
		return SmsSenderResponse{}, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("apiKey", s.Client.apiKey)

	res, err := s.Client.httpClient.Do(req)
	if err != nil {
		return SmsSenderResponse{}, fmt.Errorf("failed to send SMS: %w", err)
	}
	if res == nil {
		return SmsSenderResponse{}, errors.New("unexpected empty response")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		return SmsSenderResponse{
			ErrorResponse: ErrorResponse{
				HasError: true,
				Message:  "Message not sent",
			},
		}, fmt.Errorf("non-201 response status: %d", res.StatusCode)
	}

	var data map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return SmsSenderResponse{}, fmt.Errorf("failed to decode response: %w", err)
	}

	smsMessageData, ok := data["SMSMessageData"].(map[string]interface{})
	if !ok {
		return SmsSenderResponse{}, errors.New("invalid response format: missing 'SMSMessageData'")
	}

	recipients, err := parseRecipients(smsMessageData["Recipients"])
	if err != nil {
		return SmsSenderResponse{}, fmt.Errorf("failed to parse recipients: %w", err)
	}

	message, ok := smsMessageData["Message"].(string)
	if !ok {
		return SmsSenderResponse{}, errors.New("invalid message format")
	}

	return SmsSenderResponse{
		ErrorResponse: ErrorResponse{HasError: false},
		SmsMessageData: SmsMessageData{
			Message:    message,
			Cost:       extractCostFromMessage(message),
			Recipients: recipients,
		},
	}, nil
}

func parseRecipients(data interface{}) ([]Recipient, error) {
	recipientsData, ok := data.([]interface{})
	if !ok {
		return nil, errors.New("invalid recipient format")
	}

	var recipients []Recipient
	for _, recipient := range recipientsData {
		recipientData, ok := recipient.(map[string]interface{})
		if !ok {
			return nil, errors.New("invalid recipient data")
		}

		recipients = append(recipients, Recipient{
			Key:         uuid.New().String(),
			Cost:        recipientData["cost"].(string),
			SmsKey:      recipientData["sms_key"].(string),
			MessageId:   recipientData["messageId"].(string),
			MessagePart: int(recipientData["messageParts"].(float64)),
			Number:      recipientData["number"].(string),
			Status:      recipientData["status"].(string),
			StatusCode:  fmt.Sprintf("%v", recipientData["statusCode"]),
		})
	}
	return recipients, nil
}

func extractCostFromMessage(message string) string {
	words := strings.Split(message, " ")
	if len(words) > 0 {
		return words[0]
	}
	return ""
}

func (s *SmsSender) RetrySendSMS(ctx context.Context, maxRetries int) (SmsSenderResponse, error) {
	for retry := 0; retry < maxRetries; retry++ {
		response, err := s.SendSMS(ctx)
		if err == nil {
			return response, nil
		}
		time.Sleep(backoff(retry))
	}
	return SmsSenderResponse{}, errors.New("max retries reached")
}

func backoff(retry int) time.Duration {
	delay := time.Duration(1<<uint(retry)) * time.Second
	jitter := time.Duration(rand.Intn(int(delay))) * time.Millisecond
	return delay + jitter
}
