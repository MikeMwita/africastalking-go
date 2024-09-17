package sms

import "net/http"

const (
	DefaultAPIURL = "https://api.africastalking.com/version1/messaging"
	SandboxAPIURL = "https://api.sandbox.africastalking.com/version1/messaging"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Client struct {
	apiURL     string
	apiKey     string
	apiUser    string
	httpClient HTTPClient
}

func NewClient(httpClient HTTPClient, apiKey, apiUser string, useSandbox bool) *Client {
	apiURL := DefaultAPIURL
	if useSandbox {
		apiURL = SandboxAPIURL
	}

	return &Client{
		apiURL:     apiURL,
		apiKey:     apiKey,
		apiUser:    apiUser,
		httpClient: httpClient,
	}
}
