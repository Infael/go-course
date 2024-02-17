package emailService

type EmailDatabase struct {
	mails []Email
}

func (emailDatabase *EmailDatabase) SaveEmail(email Email) bool {
	emailDatabase.mails = append(emailDatabase.mails, email)
	return true
}
