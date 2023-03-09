package events

import (
	"github.com/gookit/event"
)

type IEvent interface {
	Fire(name string, params map[string]any)
}

const (
	UPDATE_EMAILS_VIEW   = "update_emails_view"
	UPDATE_EMAIL_PREVIEW = "update_emails_preview"
)

type EventManager struct {
}

// constructor
func CreateTuiEventManager() *EventManager {
	return &EventManager{}
}

func (self *EventManager) Fire(name string, params map[string]any) {
	event.MustFire(name, params)
}
