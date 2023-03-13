package vm

import (
	"learngocui/tui/events"
)

const (
	ACCOUNT_SELECTED = "account_selected"
	EMAIL_SELECTED   = "email_selected"
)

type ViewModel struct {
	accounts        []*AccountVM
	selectedAccount *AccountVM
	events          events.IEvent
}

func NewVM(events events.IEvent, accounts []*AccountVM) *ViewModel {
	self := &ViewModel{
		accounts: accounts,
		events:   events,
	}

	// select account and email by default
	self.selectedAccount = self.accounts[0]
	self.accounts[0].selectedEmail = self.accounts[0].SelectEmail(0)

	return self
}

func (self *ViewModel) AddAccount(account *AccountVM) {
	self.accounts = append(self.accounts, account)
}

func (self *ViewModel) GetSelectedtAccount() *AccountVM {
	return self.selectedAccount
}

func (self *ViewModel) SelectAccount(index int) *AccountVM {
	self.selectedAccount = self.accounts[index]
	self.events.Fire(ACCOUNT_SELECTED, map[string]any{"selectedAccount": self.selectedAccount})
	return self.selectedAccount
}

func (self *ViewModel) GetAccountNames() []string {
	accountNames := []string{}
	for _, accVM := range self.accounts {
		accountNames = append(accountNames, accVM.GetAccount().ShortName)
	}
	return accountNames
}
