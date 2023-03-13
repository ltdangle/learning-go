package vm

import (
	"learngocui/model"
	"learngocui/tui/events"
)

// accountVM email account view model
type accountVM struct {
	account       *model.EmailAccount
	selectedEmail *model.Email
	events        events.IEvent
}

func NewAccountVM(events events.IEvent, account *model.EmailAccount) *accountVM {
	return &accountVM{
		account:       account,
		events:        events,
		selectedEmail: nil,
	}
}

func (self *accountVM) GetAccount() *model.EmailAccount {
	return self.account
}

func (self *accountVM) SelectEmail(index int) *model.Email {
	if index >= len(self.account.Emails) {
		return nil
	}

	self.selectedEmail = &self.account.Emails[index]
	self.events.Fire(EMAIL_SELECTED, map[string]any{"selectedEmail": self.selectedEmail})
	return self.selectedEmail
}
func (self *accountVM) GetSelectedEmail() *model.Email {
	if self == nil {
		return nil
	}
	return self.selectedEmail
}

func (self *accountVM) GetEmailsAsList() []string {
	var emails []string

	if self == nil {
		return emails
	}

	for _, email := range self.account.Emails {
		emails = append(emails, email.Subject)
	}
	return emails
}
