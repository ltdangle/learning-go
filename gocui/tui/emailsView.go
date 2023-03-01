package tui

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"strconv"
)

// emailsV email accounts view
type emailsV struct {
	view *gocui.View
}

func createEmailsView(gui *gocui.Gui, startX, startY, endX, endY int) (*emailsV, error) {
	var err error
	self := &emailsV{}

	if self.view, err = gui.SetView(EMAILS_VIEW, startX, startY, endX, endY); err != nil {
		if err != gocui.ErrUnknownView {
			return nil, err
		}
		self.view.Title = strconv.Itoa(startX) + " - " + strconv.Itoa(endX) + " Emails"
		self.view.Autoscroll = true
		self.view.Highlight = true
		self.view.SelBgColor = gocui.ColorGreen
		self.view.SelFgColor = gocui.ColorBlack

		if err := self.populateEmails(gui, 0); err != nil {
			return nil, err
		}
		if err := gui.SetKeybinding(EMAILS_VIEW, gocui.KeyArrowUp, gocui.ModNone, self.cursorUpEmails); err != nil {
			return nil, err
		}
		if err := gui.SetKeybinding(EMAILS_VIEW, gocui.KeyArrowDown, gocui.ModNone, self.cursorDownEmails); err != nil {
			return nil, err
		}
	}
	return self, nil
}

func (self *emailsV) populateEmails(g *gocui.Gui, emailAccountIndex int) error {
	v, _ := g.View(EMAILS_VIEW)
	v.Clear()
	for _, email := range Emails[emailAccountIndex] {
		fmt.Fprintln(v, email)
	}

	return nil
}

func (self *emailsV) cursorDownEmails(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		showLog(g, "Cursor down: "+Accounts[cy])
		if cy+1 == 5 {
			return nil
		}
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}

	}
	return nil
}

func (self *emailsV) cursorUpEmails(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		ox, oy := v.Origin()
		cx, cy := v.Cursor()
		if err := v.SetCursor(cx, cy-1); err != nil && oy > 0 {
			if err := v.SetOrigin(ox, oy-1); err != nil {
				return err
			}
		}

		showLog(g, "Cursor down: "+strconv.Itoa(cy))
	}
	return nil
}
