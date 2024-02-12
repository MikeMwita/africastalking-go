package interfaces

import "github.com/MikeMwita/at/domain/entities"

// SMSRepository is the interface for the secondary port to save and retrieve SMS data
type SMSRepository interface {
	// SaveSMSSender saves an SMS sender entity to the database
	SaveSMSSender(smsSender *entities.SMSSender) error
	// SaveRecipients saves a list of recipient entities to the database
	SaveRecipients(recipients []*entities.Recipient) error
	// GetSMSSenderByKey gets an SMS sender entity by its key from the database
	GetSMSSenderByKey(key string) (*entities.SMSSender, error)
	// GetRecipientsBySMSKey gets a list of recipient entities by their SMS key from the database
	GetRecipientsBySMSKey(smsKey string) ([]*entities.Recipient, error)
}
