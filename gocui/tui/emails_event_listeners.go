package tui

import (
	"github.com/gookit/event"
	"learngocui/tui/logger"
	"learngocui/tui/vm"
)

func EventListeners(tui *Tui, viewModel *vm.ViewModel, logger logger.ILogger) {
	event.On(vm.ACCOUNT_SELECTED, event.ListenerFunc(func(e event.Event) error {
		logger.Log("vm.ACCOUNT_SELECTED: update emails with ")
		logger.Log(viewModel.GetSelectedtAccount().GetEmailsAsList())
		tui.emails.populate()
		return nil
	}), event.Normal)

	event.On(vm.EMAIL_SELECTED, event.ListenerFunc(func(e event.Event) error {
		logger.Log("vm.EMAIL_SELECTED: update emails with ")
		logger.Log(viewModel.GetSelectedtAccount().GetSelectedEmail())
		tui.preview.populate()
		return nil
	}), event.Normal)
}
