package store

import (
	"learngocui/model"
	"learngocui/tui"
)

const (
	ACCOUNT_SELECTED = "account selected"
	EMAIL_SELECTED   = "email selected"
)

type Store struct {
	accounts        []model.EmailAccount
	selectedAccount *model.EmailAccount
	selectedEmail   *model.Email
	events          tui.IEvent
}

func NewStore(events tui.IEvent) *Store {
	return &Store{events: events}
}

func (self *Store) SetAccounts(accounts []model.EmailAccount) {
	self.accounts = accounts
	// set default values
	self.selectedAccount = &accounts[0]
	self.selectedEmail = &self.selectedAccount.Emails[0]
}

func (self *Store) AddAccount(account model.EmailAccount) {
	self.accounts = append(self.accounts, account)
}

func (self *Store) GetSelectedtAccount() *model.EmailAccount {
	return self.selectedAccount
}

func (self *Store) GetSelectedEmail() *model.Email {
	return self.selectedEmail
}

func (self *Store) SelectAccount(shortName string) *model.EmailAccount {
	for _, acc := range self.accounts {
		if acc.ShortName == shortName {
			self.selectedAccount = &acc
			self.events.Fire(ACCOUNT_SELECTED, map[string]any{"selectedAccount": self.selectedAccount})
			return &acc
		}
	}
	return nil
}

func (self *Store) SelectEmail(index int) *model.Email {
	if index >= len(self.selectedAccount.Emails) {
		return nil
	}

	self.selectedEmail = &self.selectedAccount.Emails[index]
	self.events.Fire(EMAIL_SELECTED, map[string]any{"selectedEmail": self.selectedEmail})
	return self.selectedEmail

}
