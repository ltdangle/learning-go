package tui

import (
	"github.com/gookit/event"
	"strconv"
)

func AccountEventListeners(t *Tui) {
	event.On(ACCOUNTS_CURSOR_DOWN_EVENT, event.ListenerFunc(func(e event.Event) error {
		selectedItem := e.Data()["selectedItem"].(int)
		tuiLog("handle event from eventManager: " + EMAILS_CURSOR_DOWN_EVENT + ", selectedItem: " + strconv.Itoa(selectedItem))
		t.populateEmails([]string{"one", "two", "three"})
		return nil
	}), event.Normal)

	event.On(ACCOUNTS_CURSOR_UP_EVENT, event.ListenerFunc(func(e event.Event) error {
		selectedItem := e.Data()["selectedItem"].(int)
		tuiLog("handle event from eventManager: " + EMAILS_CURSOR_UP_EVENT + ", selectedItem: " + strconv.Itoa(selectedItem))
		t.populateEmails([]string{"three", "two", "one"})
		return nil
	}), event.Normal)
}
