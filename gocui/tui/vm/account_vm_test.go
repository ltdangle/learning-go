package vm

import (
	"reflect"
	"testing"
)

type mockEvent struct {
}

func (self *mockEvent) Fire(name string, params map[string]any) {}

func TestGetEmailsAsList(t *testing.T) {
	//vm := NewAccountVM(&mockEvent{})
	m := NewVM(&mockEvent{})
	var emails []string
	emails = m.GetSelectedtAccount().GetEmailsAsList()
	if len(emails) > 0 {
		t.Errorf("got %v, want %v", emails, []string{})
	}
	if !isSliceOfStrings(emails) {
		t.Errorf("Emails list is not a slice of strings.")
	}

}

func isSliceOfStrings(value interface{}) bool {
	// Get the reflect type of the value
	valueType := reflect.TypeOf(value)

	// Check if the value is a slice
	if valueType.Kind() != reflect.Slice {
		return false
	}

	// Check if the slice type is []string
	elementType := valueType.Elem()
	return elementType.Kind() == reflect.String
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
