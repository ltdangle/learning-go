package tui

import (
	"github.com/gookit/event"
	"strconv"
)

func eventListeners() {
	event.On(ACCOUNTS_CURSOR_DOWN_EVENT, event.ListenerFunc(func(e event.Event) error {
		selectedItem := e.Data()["selectedItem"].(int)
		tuiLog("handle event from eventManager: " + EMAILS_CURSOR_DOWN_EVENT + ", selectedItem: " + strconv.Itoa(selectedItem))

		return nil
	}), event.Normal)

	event.On(ACCOUNTS_CURSOR_UP_EVENT, event.ListenerFunc(func(e event.Event) error {
		selectedItem := e.Data()["selectedItem"].(int)
		tuiLog("handle event from eventManager: " + EMAILS_CURSOR_UP_EVENT + ", selectedItem: " + strconv.Itoa(selectedItem))
		return nil
	}), event.Normal)
}
