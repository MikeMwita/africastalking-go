package interfaces

import "github.com/MikeMwita/at/domain/entities"

type SMSService interface {
	// SendSMS sends an SMS message to the recipients and returns the response data
	SendSMS(smsSender *entities.SMSSender) (*entities.SMSMessageData, error)
}
