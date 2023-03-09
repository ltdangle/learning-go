package vm

import (
	"learngocui/model"
	"learngocui/tui/events"
)

const (
	ACCOUNT_SELECTED = "account selected"
	EMAIL_SELECTED   = "email selected"
)

type ViewModel struct {
	accounts        []model.EmailAccount
	selectedAccount *model.EmailAccount
	selectedEmail   *model.Email
	events          events.IEvent
}

func NewStore(events events.IEvent) *ViewModel {
	return &ViewModel{events: events}
}

func (self *ViewModel) SetAccounts(accounts []model.EmailAccount) {
	self.accounts = accounts
	// set default values
	self.selectedAccount = &accounts[0]
	self.selectedEmail = &self.selectedAccount.Emails[0]
}

func (self *ViewModel) AddAccount(account model.EmailAccount) {
	self.accounts = append(self.accounts, account)
}

func (self *ViewModel) GetSelectedtAccount() *model.EmailAccount {
	return self.selectedAccount
}

func (self *ViewModel) GetSelectedEmail() *model.Email {
	return self.selectedEmail
}

func (self *ViewModel) SelectAccount(shortName string) *model.EmailAccount {
	for _, acc := range self.accounts {
		if acc.ShortName == shortName {
			self.selectedAccount = &acc
			self.events.Fire(ACCOUNT_SELECTED, map[string]any{"selectedAccount": self.selectedAccount})
			return &acc
		}
	}
	return nil
}

func (self *ViewModel) SelectEmail(index int) *model.Email {
	if index >= len(self.selectedAccount.Emails) {
		return nil
	}

	self.selectedEmail = &self.selectedAccount.Emails[index]
	self.events.Fire(EMAIL_SELECTED, map[string]any{"selectedEmail": self.selectedEmail})
	return self.selectedEmail

}

func (self *ViewModel) GetAccountNames() []string {
	accounts := []string{}
	for _, acc := range self.accounts {
		accounts = append(accounts, acc.ShortName)
	}
	return accounts
}
