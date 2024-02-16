package interfaces

import "github.com/MikeMwita/at/domain/entities"

type SMSService interface {
	SendSMS(smsSender *entities.SMSSender) (*entities.SMSMessageData, error)
}
