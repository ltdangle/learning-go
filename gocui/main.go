// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jroimartin/gocui"
)

var (
	viewArr = []string{"v1", "v2", "v3", "v4"}
	active  = 0
)

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func nextView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (active + 1) % len(viewArr)
	name := viewArr[nextIndex]

	out, err := g.View("v2")
	if err != nil {
		return err
	}
	fmt.Fprintln(out, "Going from view "+v.Name()+" to "+name)

	if _, err := setCurrentViewOnTop(g, name); err != nil {
		return err
	}

	if nextIndex == 0 || nextIndex == 3 {
		g.Cursor = true
	} else {
		g.Cursor = false
	}

	active = nextIndex
	return nil
}

func layout(gui *gocui.Gui) error {
	// window size
	winX, winY := gui.Size()
	// bottom log view
	logStartX := 0
	logStartY := winY - 3
	logEndX := winX - 1
	logEndY := winY - 1
	// main view 'wrapper' (contains accounts, emails, preview)
	mainViewStartY := 0
	mainViewEndY := logStartY - 1
	// accounts view
	accountsStartX := 0
	accountsEndX := (winX / 10 * 2) - 1
	// emails list view
	emailsStartX := accountsEndX + 1
	emailsEndX := winX / 10 * 5
	// email preview view
	previewStartX := emailsEndX
	previewEndX := winX

	if accountsV, err := gui.SetView("v1", accountsStartX, mainViewStartY, accountsEndX, mainViewEndY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		accountsV.Title = strconv.Itoa(accountsStartX) + " - " + strconv.Itoa(accountsEndX) + " Accounts"
		accountsV.Editable = true
		accountsV.Wrap = true

		if _, err = setCurrentViewOnTop(gui, "v1"); err != nil {
			return err
		}
	}

	if emailsV, err := gui.SetView("v2", emailsStartX, 0, emailsEndX-1, logStartY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		emailsV.Title = strconv.Itoa(emailsStartX) + " - " + strconv.Itoa(emailsEndX) + " Emails"
		emailsV.Wrap = true
		emailsV.Autoscroll = true
		emailsV.Editable = true
	}
	if v3, err := gui.SetView("v3", previewStartX, 0, previewEndX-1, logStartY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v3.Title = strconv.Itoa(previewEndX) + " Preview"
		v3.Editable = true
	}
	if v4, err := gui.SetView("v4", logStartX, logStartY, logEndX, logEndY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v4.Title = strconv.Itoa(winX-1) + " Log"
		v4.Editable = true
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorMagenta

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
