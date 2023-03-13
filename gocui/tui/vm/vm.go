package vm

import (
	"learngocui/tui/events"
)

const (
	ACCOUNT_SELECTED = "account selected"
	EMAIL_SELECTED   = "email selected"
)

type ViewModel struct {
	accounts        []*accountVM
	selectedAccount *accountVM
	events          events.IEvent
}

func NewVM(events events.IEvent) *ViewModel {
	return &ViewModel{
		accounts: []*accountVM{},
		events:   events,
	}
}

func (self *ViewModel) AddAccount(account *accountVM) {
	self.accounts = append(self.accounts, account)
}

func (self *ViewModel) GetSelectedtAccount() *accountVM {
	return self.selectedAccount
}

func (self *ViewModel) SelectAccount(shortName string) *accountVM {
	for _, accVM := range self.accounts {
		if accVM.GetAccount().ShortName == shortName {
			self.selectedAccount = accVM
			self.events.Fire(ACCOUNT_SELECTED, map[string]any{"selectedAccount": self.selectedAccount})
			return accVM
		}
	}
	return nil
}

func (self *ViewModel) GetAccountNames() []string {
	accountNames := []string{}
	for _, accVM := range self.accounts {
		accountNames = append(accountNames, accVM.GetAccount().ShortName)
	}
	return accountNames
}
