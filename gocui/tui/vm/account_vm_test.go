package vm

import (
	"learngocui/model"
	"testing"
)

// mocks
type MockEvents struct{}

func (self *MockEvents) Fire(name string, params map[string]any) {
}

func setup() (*ViewModel, []model.EmailAccount) {
	events := &MockEvents{}
	store := NewStore(events)
	seed := model.SeedData()
	store.SetAccounts(seed)
	return store, seed
}

func TestInitialValues(t *testing.T) {
	store, seed := setup()
	AssertEqual(t, store.selectedEmail, &seed[0].Emails[0])

}

// test selecting arbitrary account
func TestSelectAccount(t *testing.T) {
	store, seed := setup()
	AssertEqual(t, store.SelectAccount(seed[3].ShortName), store.selectedAccount)
}

func TestSelectEmail(t *testing.T) {
	store, seed := setup()

	store.SelectAccount(seed[3].ShortName)
	store.SelectEmail(3)

	AssertEqual(t, store.selectedEmail, &seed[3].Emails[3])
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
