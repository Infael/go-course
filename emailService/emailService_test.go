package emailService

import "testing"

func TestSendEmail(t *testing.T) {

	const sender = "test@test.test"
	const recipient = "otherTest@test.test"
	const subject = "Test"
	const invalidEmail = "invalidEmail"

	const errorExpectedTrue = "Expected SendEmail to return true, got false"
	const errorExpectedFalse = "Expected SendEmail to return false, got true"

	emailService := &EmailService{}

	// Test case 1: Sending email via SMTP
	email := Email{
		sender:    sender,
		recipient: recipient,
		subject:   subject,
		message:   "This is a test email sent via SMTP",
	}
	protocol := SMTP
	result := emailService.SendEmail(email, protocol)
	if !result {
		t.Errorf(errorExpectedTrue)
	}

	// Test case 2: Sending email via IMAP
	email = Email{
		sender:    sender,
		recipient: recipient,
		subject:   subject,
		message:   "This is a test email sent via IMAP",
	}
	protocol = IMAP
	result = emailService.SendEmail(email, protocol)
	if !result {
		t.Errorf(errorExpectedTrue)
	}

	// Test case 3: Sending email via POP3
	email = Email{
		sender:    sender,
		recipient: recipient,
		subject:   subject,
		message:   "This is a test email sent via POP3",
	}
	protocol = POP3
	result = emailService.SendEmail(email, protocol)
	if !result {
		t.Errorf(errorExpectedTrue)
	}

	// Test case 4: Sending email with invalid protocol
	email = Email{
		sender:    "test@test.test",
		recipient: "test@example.com",
		subject:   "Test Email",
		message:   "This is a test email with SMTP protocol",
	}
	protocol = 999 // Invalid protocol
	result = emailService.SendEmail(email, protocol)
	if !result {
		t.Errorf(errorExpectedTrue)
	}

	// Test case 5: Sending email with invalid recipient
	email = Email{
		sender:    sender,
		recipient: invalidEmail,
		subject:   subject,
		message:   "This is a test email with invalid recipient",
	}

	protocol = SMTP
	result = emailService.SendEmail(email, protocol)

	if result {
		t.Errorf(errorExpectedFalse)
	}

	// Test case 6: Sending email with invalid sender
	email = Email{
		sender:    invalidEmail,
		recipient: recipient,
		subject:   subject,
		message:   "This is a test email with invalid recipient",
	}

	protocol = SMTP
	result = emailService.SendEmail(email, protocol)

	if result {
		t.Errorf(errorExpectedFalse)
	}
}
