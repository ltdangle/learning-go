package store

import "learngocui/model"

type Store struct {
	accounts        []model.EmailAccount
	selectedAccount *model.EmailAccount
	selectedEmail   *model.Email
}

func newStore() *Store {
	return &Store{}
}

func (self *Store) setAccounts(accounts []model.EmailAccount) {
	self.accounts = accounts
	// set default values
	self.selectedAccount = &accounts[0]
	self.selectedEmail = &self.selectedAccount.Emails[0]
}

func (self *Store) addAccount(account model.EmailAccount) {
	self.accounts = append(self.accounts, account)
}

func (self *Store) getSelectedtAccount() *model.EmailAccount {
	return self.selectedAccount
}

func (self *Store) getSelectedEmail() *model.Email {
	return self.selectedEmail
}

func (self *Store) selectAccount(shortName string) *model.EmailAccount {
	for _, acc := range self.accounts {
		if acc.ShortName == shortName {
			self.selectedAccount = &acc
			return &acc
		}
	}
	return nil
}

func (self *Store) selectEmail(index int) *model.Email {
	if index >= len(self.selectedAccount.Emails) {
		return nil
	}

	self.selectedEmail = &self.selectedAccount.Emails[index]
	return self.selectedEmail

}
