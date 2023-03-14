package tui

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"learngocui/tui/events"
	"learngocui/tui/vm"
	"strconv"
)

// emails email accounts view
type emails struct {
	view      *gocui.View
	event     events.IEvent
	viewModel *vm.ViewModel
}

func newEmails(event events.IEvent, viewModel *vm.ViewModel) *emails {
	return &emails{event: event, viewModel: viewModel}

}
func (self *emails) initView(gui *gocui.Gui, startX, startY, endX, endY int) error {
	var err error

	if self.view, err = gui.SetView(EMAILS_VIEW, startX, startY, endX, endY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		self.view.Title = strconv.Itoa(startX) + " - " + strconv.Itoa(endX) + " emails"
		self.view.Autoscroll = true
		self.view.Highlight = true
		self.view.SelBgColor = gocui.ColorGreen
		self.view.SelFgColor = gocui.ColorBlack

		if err := gui.SetKeybinding(EMAILS_VIEW, gocui.KeyArrowUp, gocui.ModNone, self.cursorUp); err != nil {
			return err
		}
		if err := gui.SetKeybinding(EMAILS_VIEW, gocui.KeyArrowDown, gocui.ModNone, self.cursorDown); err != nil {
			return err
		}

		self.populate()
	}

	return nil
}

func (self *emails) populate() {
	self.view.Clear()
	for _, item := range self.viewModel.GetSelectedtAccount().GetEmailsAsList() {
		fmt.Fprintln(self.view, item)
	}
	// update cursor position to selected (item) email
	self.view.SetCursor(0, self.viewModel.GetSelectedtAccount().GetSelectedEmailIndex())
}

func (self *emails) cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		selectedItem := cy + 1

		// we've reached the end of the list
		if selectedItem == len(self.viewModel.GetSelectedtAccount().GetEmailsAsList()) {
			return nil
		}

		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}

		self.viewModel.GetSelectedtAccount().SelectEmail(selectedItem)
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

		self.viewModel.GetSelectedtAccount().SelectEmail(selectedItem)
	}
	return nil
}
