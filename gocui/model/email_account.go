package model

type EmailAccount struct {
	// 'Regular' maildir account or 'virtual' notmuch account
	AccountType string
	// Account's short name
	ShortName string
	// Account's email address
	Email string
	// Shell command to be executed to get list of emails
	InboxShellCommand string
	// Folder where deleted mail is moved
	TrashFolder string
	// Email delivery dsn
	DeliveryTransport string
	// Email collection
	Emails []Email
}

func (self *EmailAccount) GetEmailsAsList() []string {
	var list []string
	for _, email := range self.Emails {
		list = append(list, email.Subject)
	}
	return list
}
