package tui

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"strconv"
)

func createAccountsView(gui *gocui.Gui, startX, startY, endX, endY int) error {

	if accountsV, err := gui.SetView(ACCOUNTS_VIEW, startX, startY, endX, endY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		accountsV.Title = strconv.Itoa(startX) + " - " + strconv.Itoa(endX) + " Accounts"

		accountsV.Autoscroll = true
		accountsV.Highlight = true
		accountsV.SelBgColor = gocui.ColorGreen
		accountsV.SelFgColor = gocui.ColorBlack

		for _, v := range Accounts {
			fmt.Fprintln(accountsV, v)
		}

		if _, err = setCurrentViewOnTop(gui, ACCOUNTS_VIEW); err != nil {
			return err
		}

		if err := gui.SetKeybinding(ACCOUNTS_VIEW, gocui.KeyArrowDown, gocui.ModNone, cursorDownAccounts); err != nil {
			return err
		}

		if err := gui.SetKeybinding(ACCOUNTS_VIEW, gocui.KeyArrowUp, gocui.ModNone, cursorUpAccounts); err != nil {
			return err
		}
	}
	return nil
}

func cursorDownAccounts(g *gocui.Gui, v *gocui.View) error {
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

		logAccountsView(g, "Selected item: "+strconv.Itoa(selectedItem)+" = "+Accounts[selectedItem])
		selectedText, _ := v.Line(selectedItem)
		logAccountsView(g, "Selected text: "+selectedText)
		if err := populateEmailsView(g, selectedItem); err != nil {
			return err
		}
	}
	return nil
}

func cursorUpAccounts(g *gocui.Gui, v *gocui.View) error {
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

		logAccountsView(g, "Selected: "+strconv.Itoa(selectedItem)+" = "+Accounts[selectedItem])
	}
	return nil
}

func logAccountsView(g *gocui.Gui, msg string) {
	showLog(g, msg)
}
