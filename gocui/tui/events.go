package tui

import (
	"github.com/gookit/event"
	"github.com/jroimartin/gocui"
)

type IEvent interface {
	Fire(name string, params map[string]any)
}

const (
	UPDATE_EMAILS_VIEW   = "update_emails_view"
	UPDATE_EMAIL_PREVIEW = "update_emails_view"
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
