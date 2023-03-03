package tui

import (
	"github.com/gookit/event"
	"github.com/jroimartin/gocui"
)

type IEvent interface {
	Fire(name string, params map[string]any)
	On(name string, callback func())
}

const (
	UPDATE_EMAILS_VIEW = "update_emails_view"
)

type eventManager struct {
	gui *gocui.Gui
}

// constructor
func createTuiEventManager(gui *gocui.Gui) *eventManager {
	return &eventManager{gui}
}

func (self *eventManager) Fire(name string, params map[string]any) {
	event.MustFire(name, params)
}
func (self *eventManager) On(name string, callback func()) {
	event.On(name, event.ListenerFunc(func(e event.Event) error {
		callback()
		return nil
	}), event.Normal)
}
