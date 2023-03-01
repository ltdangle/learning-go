package tui

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"strconv"
)

func createEmailsView(gui *gocui.Gui, startX, startY, endX, endY int) error {
	if emailsV, err := gui.SetView(EMAILS_VIEW, startX, startY, endX, endY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		emailsV.Title = strconv.Itoa(startX) + " - " + strconv.Itoa(endX) + " Emails"

		emailsV.Autoscroll = true
		emailsV.Highlight = true
		emailsV.SelBgColor = gocui.ColorGreen
		emailsV.SelFgColor = gocui.ColorBlack

		if err := populateEmailsView(gui, 0); err != nil {
			return err
		}
		if err := gui.SetKeybinding(EMAILS_VIEW, gocui.KeyArrowUp, gocui.ModNone, cursorUpEmails); err != nil {
			return err
		}
		if err := gui.SetKeybinding(EMAILS_VIEW, gocui.KeyArrowDown, gocui.ModNone, cursorDownEmails); err != nil {
			return err
		}
	}
	return nil
}

func populateEmailsView(g *gocui.Gui, emailAccountIndex int) error {
	v, _ := g.View(EMAILS_VIEW)
	v.Clear()
	for _, email := range Emails[emailAccountIndex] {
		fmt.Fprintln(v, email)
	}

	return nil
}

func cursorDownEmails(g *gocui.Gui, v *gocui.View) error {
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

func cursorUpEmails(g *gocui.Gui, v *gocui.View) error {
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
