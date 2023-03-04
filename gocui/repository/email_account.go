package repository

import "learngocui/model"

type IAccountRepository interface {
	FindById(id int) model.EmailAccount
}

type SeedAccountRepository struct {
	store []model.EmailAccount
}

func (self *SeedAccountRepository) FindById(id int) *model.EmailAccount {
	if id >= len(self.store) {
		return nil
	}

	return &self.store[id]
}
