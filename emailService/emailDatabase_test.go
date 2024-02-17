package emailService

import (
	"testing"
)

func TestSaveEmail(t *testing.T) {
	emailDB := EmailDatabase{}

	// Create a sample email
	email := Email{
		sender:    "test@test.test",
		recipient: "otherTest@test.test",
		subject:   "Test",
		message:   "This is a test",
	}

	// Save the email
	saved := emailDB.SaveEmail(email)

	// Check if the email was saved successfully
	if !saved {
		t.Errorf("Failed to save email")
	}

	// Check if the email is present in the database
	if len(emailDB.mails) != 1 || emailDB.mails[0] != email {
		t.Errorf("Email not saved correctly")
	}
}
