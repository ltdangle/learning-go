package store

import (
	"learngocui/model"
	"learngocui/repository"
	"testing"
)

func setup() (*Store, []model.EmailAccount) {
	store := newStore()
	seed := repository.SeedData()
	store.setAccounts(seed)
	return store, seed
}
func TestInitialValues(t *testing.T) {
	store, seed := setup()
	AssertEqual(t, store.selectedEmail, &seed[0].Emails[0])

}
func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
