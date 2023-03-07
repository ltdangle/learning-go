package store

import (
	"learngocui/model"
	"learngocui/repository"
	"testing"
)

// mocks
type MockEvents struct{}

func (self *MockEvents) Fire(name string, params map[string]any) {
}

func setup() (*Store, []model.EmailAccount) {
	events := &MockEvents{}
	store := newStore(events)
	seed := repository.SeedData()
	store.setAccounts(seed)
	return store, seed
}

func TestInitialValues(t *testing.T) {
	store, seed := setup()
	AssertEqual(t, store.selectedEmail, &seed[0].Emails[0])

}

// test selecting arbitrary account
func TestSelectAccount(t *testing.T) {
	store, seed := setup()
	AssertEqual(t, store.selectAccount(seed[3].ShortName), store.selectedAccount)
}

func TestSelectEmail(t *testing.T) {
	store, seed := setup()

	store.selectAccount(seed[3].ShortName)
	store.selectEmail(3)

	AssertEqual(t, store.selectedEmail, &seed[3].Emails[3])
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
