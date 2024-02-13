package adapters

//type PostgresSMSAdapter struct {
//	db *sql.DB
//}
//
//func (p *PostgresSMSAdapter) SaveSMSSender(smsSender *entities.SMSSender) error {
//	// Prepare the SQL statement for inserting an SMS sender
//	stmt, err := p.db.Prepare("INSERT INTO climatequestions (id, title) VALUES ($1, $2)")
//	if err != nil {
//		return err
//	}
//	defer stmt.Close()
//
//	// Execute the SQL statement with the SMS sender fields
//	_, err = stmt.Exec(smsSender.Key, smsSender.Title)
//	if err != nil {
//		return err
//	}
//
//	// Return nil if no error
//	return nil
//}
//
//func (p *PostgresSMSAdapter) SaveRecipients(recipients []*entities.Recipient) error {
//	// Prepare the SQL statement for inserting a recipient
//	stmt, err := p.db.Prepare("INSERT INTO steps (id, label, number, question_id) VALUES ($1, $2, $3, $4)")
//	if err != nil {
//		return err
//	}
//	defer stmt.Close()
//
//	// Loop over the recipients slice and execute the SQL statement for each recipient
//	for _, recipient := range recipients {
//		_, err = stmt.Exec(recipient.Key, recipient.Label, recipient.Number, recipient.QuestionId)
//		if err != nil {
//			return err
//		}
//	}
//
//	// Return nil if no error
//	return nil
//}
//
//func (p *PostgresSMSAdapter) GetSMSSenderByKey(key string) (*entities.SMSSender, error) {
//	// Prepare the SQL statement for querying an SMS sender by key
//	stmt, err := p.db.Prepare("SELECT id, title FROM climatequestions WHERE id = $1")
//	if err != nil {
//		return nil, err
//	}
//	defer stmt.Close()
//
//	// Execute the SQL statement with the key as the argument
//	row := stmt.QueryRow(key)
//
//	// Create an empty SMS sender struct
//	smsSender := &entities.SMSSender{}
//
//	// Scan the row into the SMS sender struct
//	err = row.Scan(&smsSender.Key, &smsSender.Title)
//	if err != nil {
//		return nil, err
//	}
//
//	// Return the SMS sender struct and nil error
//	return smsSender, nil
//}
//
//func (p *PostgresSMSAdapter) GetRecipientsBySMSKey(smsKey string) ([]*entities.Recipient, error) {
//	// Prepare the SQL statement for querying recipients by SMS key
//	stmt, err := p.db.Prepare("SELECT id, label, number, question_id FROM steps WHERE question_id = $1")
//	if err != nil {
//		return nil, err
//	}
//	defer stmt.Close()
//
//	// Execute the SQL statement with the SMS key as the argument
//	rows, err := stmt.Query(smsKey)
//	if err != nil {
//		return nil, err
//	}
//	defer rows.Close()
//
//	// Create an empty slice of pointers to recipient structs
//	recipients := []*entities.Recipient{}
//
//	// Loop over the rows and scan each row into a recipient struct
//	for rows.Next() {
//		// Create an empty recipient struct
//		recipient := &entities.Recipient{}
//
//		// Scan the row into the recipient struct
//		err = rows.Scan(&recipient.Key, &recipient.Label, &recipient.Number, &recipient.QuestionId)
//		if err != nil {
//			return nil, err
//		}
//
//		// Append the recipient struct to the recipients slice
//		recipients = append(recipients, recipient)
//	}
//
//	// Check for any error from iterating over the rows
//	err = rows.Err()
//	if err != nil {
//		return nil, err
//	}
//
//	// Return the recipients slice and nil error
//	return recipients, nil
//}
//
//func NewPostgresSMSAdapter(connectionString string) (interfaces.SMSRepository, error) {
//	db, err := sql.Open("postgres", connectionString)
//	if err != nil {
//		return nil, err
//	}
//	return &PostgresSMSAdapter{db: db}, nil
//}
