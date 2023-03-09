package tui

import (
	"github.com/jroimartin/gocui"
	"learngocui/events"
	"strconv"
)

// emails email accounts view
type emails struct {
	view  *gocui.View
	event events.IEvent
}

const (
	EMAILS_CURSOR_DOWN_EVENT = "emails_cursor_down"
	EMAILS_CURSOR_UP_EVENT   = "emails_cursor_up"
)

func newEmails(event events.IEvent) *emails {
	return &emails{event: event}

}
func (self *emails) initView(gui *gocui.Gui, startX, startY, endX, endY int) error {
	var err error

	if self.view, err = gui.SetView(EMAILS_VIEW, startX, startY, endX, endY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		self.view.Title = strconv.Itoa(startX) + " - " + strconv.Itoa(endX) + " Emails"
		self.view.Autoscroll = true
		self.view.Highlight = true
		self.view.SelBgColor = gocui.ColorGreen
		self.view.SelFgColor = gocui.ColorBlack

		if err := self.populate(gui, 0); err != nil {
			return err
		}
		if err := gui.SetKeybinding(EMAILS_VIEW, gocui.KeyArrowUp, gocui.ModNone, self.cursorUp); err != nil {
			return err
		}
		if err := gui.SetKeybinding(EMAILS_VIEW, gocui.KeyArrowDown, gocui.ModNone, self.cursorDown); err != nil {
			return err
		}
	}
	return nil
}

func (self *emails) populate(g *gocui.Gui, emailAccountIndex int) error {
	v, _ := g.View(EMAILS_VIEW)
	v.Clear()
	return nil
}

func (self *emails) cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		// TODO 5?....

		if cy+1 == 5 {
			return nil
		}

		selectedItem := cy + 1

		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}
		tuiLog(g, "Cursor down: "+strconv.Itoa(cy))
		self.event.Fire(EMAILS_CURSOR_DOWN_EVENT, map[string]any{"selectedItem": selectedItem})
	}
	return nil
}

func (self *emails) cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()

		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}

		selectedItem := cy - 1
		// check that selectedItem is not out of bounds
		if selectedItem < 0 {
			return nil
		}

		tuiLog(g, "Cursor up: "+strconv.Itoa(cy))
		self.event.Fire(EMAILS_CURSOR_UP_EVENT, map[string]any{"selectedItem": selectedItem})
	}
	return nil
}