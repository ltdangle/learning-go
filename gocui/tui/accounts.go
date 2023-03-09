package tui

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"learngocui/events"
	"strconv"
)

// accounts email accounts view
type accounts struct {
	// items to be displayed
	items []string
	view  *gocui.View
	event events.IEvent
}

const (
	ACCOUNTS_CURSOR_DOWN_EVENT = "accounts_cursor_down"
	ACCOUNTS_CURSOR_UP_EVENT   = "accounts_cursor_up"
)

func newAccountsV(event events.IEvent, items []string) *accounts {
	return &accounts{
		items: items,
		event: event,
	}
}

func (self *accounts) initView(gui *gocui.Gui, startX, startY, endX, endY int) error {

	var err error
	if self.view, err = gui.SetView(ACCOUNTS_VIEW, startX, startY, endX, endY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		self.view.Title = strconv.Itoa(startX) + " - " + strconv.Itoa(endX) + " Accounts"
		self.view.Autoscroll = true
		self.view.Highlight = true
		self.view.SelBgColor = gocui.ColorGreen
		self.view.SelFgColor = gocui.ColorBlack
		
		if _, err = setCurrentViewOnTop(gui, ACCOUNTS_VIEW); err != nil {
			return err
		}

		if err = gui.SetKeybinding(ACCOUNTS_VIEW, gocui.KeyArrowDown, gocui.ModNone, self.cursorDown); err != nil {
			return err
		}

		if err = gui.SetKeybinding(ACCOUNTS_VIEW, gocui.KeyArrowUp, gocui.ModNone, self.cursorUp); err != nil {
			return err
		}

		self.populate()
	}
	return nil
}

func (self *accounts) populate() {
	self.view.Clear()
	for _, item := range self.items {
		fmt.Fprintln(self.view, item)
	}
}

func (self *accounts) cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		selectedItem := cy + 1

		// we've reached the end of the list
		if selectedItem == len(self.items) {
			return nil
		}

		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}

		self.log(g, "Selected item: "+strconv.Itoa(selectedItem))
		selectedText, _ := v.Line(selectedItem)
		self.log(g, "Selected text: "+selectedText)

		self.event.Fire(ACCOUNTS_CURSOR_DOWN_EVENT, map[string]any{"selectedItem": selectedItem})
	}
	return nil
}

func (self *accounts) cursorUp(g *gocui.Gui, v *gocui.View) error {
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

		self.event.Fire(ACCOUNTS_CURSOR_UP_EVENT, map[string]any{"selectedItem": selectedItem})
	}
	return nil
}

func (self *accounts) log(g *gocui.Gui, msg string) {
	tuiLog(g, msg)
}
