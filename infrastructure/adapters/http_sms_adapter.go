package adapters

import (
	"encoding/json"
	"github.com/MikeMwita/at/config"
	"github.com/MikeMwita/at/domain/entities"
	"github.com/MikeMwita/at/domain/interfaces"
	"net/http"
	"strings"
)

type HTTPSMSHandler struct {
	smsService interfaces.SMSService
}

func (h *HTTPSMSHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Parse request parameters
	r.ParseForm()
	recipients := strings.Split(r.FormValue("recipient"), ",")
	message := r.FormValue("message")
	sender := r.FormValue("sender")

	// Validate input parameters
	if len(recipients) == 0 || message == "" {
		http.Error(w, "Missing required parameters: recipient, message", http.StatusBadRequest)
		return
	}

	cfg := config.LoadConfig()
	// Create the SMSSender object
	smsSender := &entities.SMSSender{
		APIUser:    cfg.Username,
		APIKey:     cfg.APIKey,
		Recipients: recipients,
		Message:    message,
		Sender:     sender,
	}

	// Call the SendSMS method of the sms.Service instance
	sendResponse, err := h.smsService.SendSMS(smsSender)

	// Handle errors and sendResponse
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return HTTP response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sendResponse)
}

func NewHTTPSMSHandler(smsService interfaces.SMSService) *HTTPSMSHandler {
	return &HTTPSMSHandler{smsService: smsService}
}
