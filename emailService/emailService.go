package emailService

import (
	"fmt"
	"net/mail"
)

type EmailService struct {
	dbConnection EmailDatabase
}

func (emailService *EmailService) validateEmail(email Email) bool {
	_, errRecipient := mail.ParseAddress(email.recipient)
	_, errSender := mail.ParseAddress(email.sender)
	return errRecipient == nil && errSender == nil
}

func (emailService *EmailService) SaveEmail(email Email) bool {
	return emailService.dbConnection.SaveEmail(email)
}

func (emailService *EmailService) SendEmail(email Email, protocol int) bool {
	// validate email
	valid := emailService.validateEmail(email)
	if !valid {
		return false
	}

	// Code to send email
	switch protocol {
	case SMTP:
		fmt.Println("Sending email via SMTP")
	case IMAP:
		fmt.Println("Sending email via IMAP")
	case POP3:
		fmt.Println("Sending email via POP3")
	default:
		fmt.Println("Sending email via SMTP")
	}

	// Save email
	emailService.SaveEmail(email)

	return true
}
