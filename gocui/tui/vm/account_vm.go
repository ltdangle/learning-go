package vm

import (
	"learngocui/model"
	"learngocui/tui/events"
)

// AccountVM email account view model
type AccountVM struct {
	account       *model.EmailAccount
	selectedEmail *model.Email
	events        events.IEvent
}

func NewAccountVM(events events.IEvent, account *model.EmailAccount) *AccountVM {
	vm := &AccountVM{
		account: account,
		events:  events,
	}
	vm.SelectEmail(0)
	return vm
}

func (self *AccountVM) GetAccount() *model.EmailAccount {
	return self.account
}

func (self *AccountVM) SelectEmail(index int) *model.Email {
	if index >= len(self.account.Emails) {
		return nil
	}

	self.selectedEmail = &self.account.Emails[index]
	self.events.Fire(EMAIL_SELECTED, map[string]any{"selectedEmail": self.selectedEmail})
	return self.selectedEmail
}
func (self *AccountVM) GetSelectedEmail() *model.Email {
	if self == nil {
		return nil
	}
	return self.selectedEmail
}

func (self *AccountVM) GetEmailsAsList() []string {
	var emails []string

	if self == nil {
		return emails
	}

	for _, email := range self.account.Emails {
		emails = append(emails, email.Subject)
	}
	return emails
}
