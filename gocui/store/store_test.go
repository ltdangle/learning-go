package store

import (
	"learngocui/repository"
	"testing"
)

func TestInitialValues(t *testing.T) {
	store := newStore()
	seed := repository.SeedData()
	store.setAccounts(seed)
	AssertEqual(t, store.selectedEmail, &seed[0].Emails[0])

}
func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
