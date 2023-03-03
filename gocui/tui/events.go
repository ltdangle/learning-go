package tui

import "github.com/gookit/event"

type IEvent interface {
	Fire(name string, params map[string]any)
}

const (
	UPDATE_EMAILS_VIEW = "update_emails_view"
)

type TuiEventManager struct {
}

// constructor
func createTuiEventManager() *TuiEventManager {
	return &TuiEventManager{}
}

func (self *TuiEventManager) Fire(name string, params map[string]any) {
	event.MustFire(name, params)
}
