package sms

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"
)

// Mock Doer to simulate HTTP client behavior
type MockDoer struct {
	Responses []*http.Response
	Errors    []error
	CallCount int
}

func (m *MockDoer) Do(req *http.Request) (*http.Response, error) {
	if m.CallCount >= len(m.Responses) {
		return nil, fmt.Errorf("out of mock responses")
	}
	resp := m.Responses[m.CallCount]
	err := m.Errors[m.CallCount]
	m.CallCount++
	return resp, err
}

// Helper function to create a mock response
func newMockResponse(statusCode int, body string) *http.Response {
	return &http.Response{
		StatusCode: statusCode,
		Body:       io.NopCloser(strings.NewReader(body)),
		Header:     make(http.Header),
	}
}

func TestSendSMS(t *testing.T) {
	tests := []struct {
		name            string
		mockResponses   []*http.Response
		mockErrors      []error
		expectedMessage string
		expectedError   bool
	}{
		{
			name: "Success - Message sent",
			mockResponses: []*http.Response{
				newMockResponse(http.StatusCreated, `{
					"SMSMessageData": {
						"Message": "Sent",
						"Recipients": [
							{
								"cost": "KES 0.80",
								"messageId": "XYZ123",
								"number": "+254745617596",
								"status": "Success",
								"statusCode": 200
							}
						]
					}
				}`),
			},
			mockErrors:      []error{nil},
			expectedMessage: "Sent",
			expectedError:   false,
		},
		{
			name: "Failure - Message not sent",
			mockResponses: []*http.Response{
				newMockResponse(http.StatusBadRequest, ""),
			},
			mockErrors:      []error{fmt.Errorf("bad request")},
			expectedMessage: "",
			expectedError:   true,
		},
		{
			name: "Empty recipients list",
			mockResponses: []*http.Response{
				newMockResponse(http.StatusCreated, `{
					"SMSMessageData": {
						"Message": "Sent",
						"Recipients": []
					}
				}`),
			},
			mockErrors:      []error{nil},
			expectedMessage: "Sent",
			expectedError:   false,
		},
		{
			name: "Invalid JSON response format",
			mockResponses: []*http.Response{
				newMockResponse(http.StatusCreated, `{invalid-json}`),
			},
			mockErrors:      []error{nil},
			expectedMessage: "",
			expectedError:   true,
		},
		{
			name: "Invalid recipient data",
			mockResponses: []*http.Response{
				newMockResponse(http.StatusCreated, `{
					"SMSMessageData": {
						"Message": "Sent",
						"Recipients": [{}]
					}
				}`),
			},
			mockErrors:      []error{nil},
			expectedMessage: "",
			expectedError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				apiURL:     DefaultAPIURL,
				apiKey:     "test-api-key",
				apiUser:    "test-user",
				httpClient: &MockDoer{Responses: tt.mockResponses, Errors: tt.mockErrors},
			}

			sender := &SmsSender{
				Client:     client,
				Recipients: []string{"+254745617596"},
				Message:    "Hello, this is a test message",
				Sender:     "TestSender",
			}

			ctx := context.Background()

			response, err := sender.SendSMS(ctx)

			if tt.expectedError && err == nil {
				t.Errorf("expected error but got none")
			}

			if !tt.expectedError && err != nil {
				t.Errorf("did not expect error but got: %v", err)
			}

			if response.SmsMessageData.Message != tt.expectedMessage {
				t.Errorf("expected message: %s, got: %s", tt.expectedMessage, response.SmsMessageData.Message)
			}
		})
	}
}

func TestRetrySendSMS(t *testing.T) {
	tests := []struct {
		name          string
		mockResponses []*http.Response
		mockErrors    []error
		maxRetries    int
		expectedError bool
		expectedMsg   string
	}{
		{
			name: "Success after retry",
			mockResponses: []*http.Response{
				newMockResponse(http.StatusInternalServerError, ""),
				newMockResponse(http.StatusCreated, `{
					"SMSMessageData": {
						"Message": "Sent",
						"Recipients": [
							{
								"cost": "KES 0.80",
								"messageId": "XYZ123",
								"number": "+254745617596",
								"status": "Success",
								"statusCode": 200
							}
						]
					}
				}`),
			},
			mockErrors:    []error{fmt.Errorf("internal error"), nil},
			maxRetries:    3,
			expectedError: false,
			expectedMsg:   "Sent",
		},
		{
			name: "Failure after retries",
			mockResponses: []*http.Response{
				newMockResponse(http.StatusInternalServerError, ""),
				newMockResponse(http.StatusInternalServerError, ""),
				newMockResponse(http.StatusInternalServerError, ""),
			},
			mockErrors:    []error{fmt.Errorf("internal error"), fmt.Errorf("internal error"), fmt.Errorf("internal error")},
			maxRetries:    3,
			expectedError: true,
			expectedMsg:   "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				apiURL:     DefaultAPIURL,
				apiKey:     "test-api-key",
				apiUser:    "test-user",
				httpClient: &MockDoer{Responses: tt.mockResponses, Errors: tt.mockErrors},
			}

			sender := &SmsSender{
				Client:     client,
				Recipients: []string{"+254745617596"},
				Message:    "Hello, this is a test message",
				Sender:     "TestSender",
			}

			ctx := context.Background()

			response, err := sender.RetrySendSMS(ctx, tt.maxRetries)

			if tt.expectedError && err == nil {
				t.Errorf("expected error but got none")
			}

			if !tt.expectedError && err != nil {
				t.Errorf("did not expect error but got: %v", err)
			}

			if response.SmsMessageData.Message != tt.expectedMsg {
				t.Errorf("expected message: %s, got: %s", tt.expectedMsg, response.SmsMessageData.Message)
			}
		})
	}
}
