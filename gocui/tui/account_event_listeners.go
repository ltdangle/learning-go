package tui

import (
	"github.com/gookit/event"
	"learngocui/tui/vm"
)

func AccountEventListeners(vm *vm.ViewModel) {
	event.On(ACCOUNTS_CURSOR_DOWN_EVENT, event.ListenerFunc(func(e event.Event) error {
		selectedItem := e.Data()["selectedItem"].(int)
		vm.SelectAccount(selectedItem)
		return nil
	}), event.Normal)

	event.On(ACCOUNTS_CURSOR_UP_EVENT, event.ListenerFunc(func(e event.Event) error {
		selectedItem := e.Data()["selectedItem"].(int)
		vm.SelectAccount(selectedItem)
		return nil
	}), event.Normal)
}
