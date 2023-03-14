package tui

import (
	"github.com/gookit/event"
	"learngocui/tui/logger"
	"learngocui/tui/vm"
)

func EmailEventListeners(t *Tui, viewModel *vm.ViewModel, logger logger.ILogger) {
	event.On(vm.ACCOUNT_SELECTED, event.ListenerFunc(func(e event.Event) error {
		logger.Log("vm.ACCOUNT_SELECTED: update emails with ")
		logger.Log(viewModel.GetSelectedtAccount().GetEmailsAsList())
		t.emails.populate()
		return nil
	}), event.Normal)
}
