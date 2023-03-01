package tui

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"strconv"
)

// AccountsV email accounts view
type AccountsV struct {
	view *gocui.View
}

func createAccountsView(gui *gocui.Gui, startX, startY, endX, endY int) (*AccountsV, error) {

	var err error
	self := &AccountsV{}
	if self.view, err = gui.SetView(ACCOUNTS_VIEW, startX, startY, endX, endY); err != nil {
		if err != gocui.ErrUnknownView {
			return nil, err
		}
		self.view.Title = strconv.Itoa(startX) + " - " + strconv.Itoa(endX) + " Accounts"
		self.view.Autoscroll = true
		self.view.Highlight = true
		self.view.SelBgColor = gocui.ColorGreen
		self.view.SelFgColor = gocui.ColorBlack

		for _, v := range Accounts {
			fmt.Fprintln(self.view, v)
		}

		if _, err = setCurrentViewOnTop(gui, ACCOUNTS_VIEW); err != nil {
			return nil, err
		}

		if err = gui.SetKeybinding(ACCOUNTS_VIEW, gocui.KeyArrowDown, gocui.ModNone, self.cursorDownAccounts); err != nil {
			return nil, err
		}

		if err = gui.SetKeybinding(ACCOUNTS_VIEW, gocui.KeyArrowUp, gocui.ModNone, self.cursorUpAccounts); err != nil {
			return nil, err
		}
	}
	return self, nil
}

func (self AccountsV) cursorDownAccounts(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		selectedItem := cy + 1
		if selectedItem == 5 {
			return nil
		}
		if err := v.SetCursor(cx, cy+1); err != nil {
			ox, oy := v.Origin()
			if err := v.SetOrigin(ox, oy+1); err != nil {
				return err
			}
		}

		self.log(g, "Selected item: "+strconv.Itoa(selectedItem)+" = "+Accounts[selectedItem])
		selectedText, _ := v.Line(selectedItem)
		self.log(g, "Selected text: "+selectedText)
		if err := populateEmailsView(g, selectedItem); err != nil {
			return err
		}
	}
	return nil
}

func (self AccountsV) cursorUpAccounts(g *gocui.Gui, v *gocui.View) error {
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

		if err := populateEmailsView(g, selectedItem); err != nil {
			return err
		}

		self.log(g, "Selected: "+strconv.Itoa(selectedItem)+" = "+Accounts[selectedItem])
	}
	return nil
}

func (self AccountsV) log(g *gocui.Gui, msg string) {
	showLog(g, msg)
}
