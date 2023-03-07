package tui

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"learngocui/events"
	"learngocui/repository"
	"strconv"
)

// emailsV email accounts view
type emailsV struct {
	view         *gocui.View
	accountsRepo repository.IAccountRepository
	event        events.IEvent
}

func createEmailsView(event events.IEvent, gui *gocui.Gui, accountsRepo repository.IAccountRepository, startX, startY, endX, endY int) (*emailsV, error) {
	var err error
	self := &emailsV{}
	self.event = event
	self.accountsRepo = accountsRepo

	if self.view, err = gui.SetView(EMAILS_VIEW, startX, startY, endX, endY); err != nil {
		if err != gocui.ErrUnknownView {
			return nil, err
		}
		self.view.Title = strconv.Itoa(startX) + " - " + strconv.Itoa(endX) + " Emails"
		self.view.Autoscroll = true
		self.view.Highlight = true
		self.view.SelBgColor = gocui.ColorGreen
		self.view.SelFgColor = gocui.ColorBlack

		if err := self.populate(gui, 0); err != nil {
			return nil, err
		}
		if err := gui.SetKeybinding(EMAILS_VIEW, gocui.KeyArrowUp, gocui.ModNone, self.cursorUp); err != nil {
			return nil, err
		}
		if err := gui.SetKeybinding(EMAILS_VIEW, gocui.KeyArrowDown, gocui.ModNone, self.cursorDown); err != nil {
			return nil, err
		}
	}
	return self, nil
}

func (self *emailsV) populate(g *gocui.Gui, emailAccountIndex int) error {
	v, _ := g.View(EMAILS_VIEW)
	v.Clear()
	for _, email := range self.accountsRepo.FindById(emailAccountIndex).Emails {
		fmt.Fprintln(v, email)
	}

	return nil
}

func (self *emailsV) cursorDown(g *gocui.Gui, v *gocui.View) error {
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
		tuiLog(g, "Cursor down: "+Accounts[cy])
		self.event.Fire(events.UPDATE_EMAIL_PREVIEW, map[string]any{"selectedItem": selectedItem})
	}
	return nil
}

func (self *emailsV) cursorUp(g *gocui.Gui, v *gocui.View) error {
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
		self.event.Fire(events.UPDATE_EMAIL_PREVIEW, map[string]any{"selectedItem": selectedItem})
	}
	return nil
}
