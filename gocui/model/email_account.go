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
}

func (ea *EmailAccount) ToArray() map[string]interface{} {
	arr := make(map[string]interface{})
	arr["accountType"] = ea.AccountType
	arr["shortName"] = ea.ShortName
	arr["email"] = ea.Email
	arr["inboxShellCommand"] = ea.InboxShellCommand
	arr["trashFolder"] = ea.TrashFolder
	arr["dsn"] = ea.DeliveryTransport
	return arr
}
