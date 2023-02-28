// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jroimartin/gocui"
)

const (
	ACCOUNTS_VIEW = "accounts view"
	EMAILS_VIEW   = "emails view"
	PREVIEW_VIEW  = "preview view"
	BOTTOM_VIEW   = "bottom view"
)

var (
	viewArr = []string{ACCOUNTS_VIEW, EMAILS_VIEW, PREVIEW_VIEW, BOTTOM_VIEW}
	active  = 0
	Items   = []string{"one", "two", "three", "four", "five"}
)

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func maximizePreview(g *gocui.Gui, v *gocui.View) error {
	bottomV, err := g.View(BOTTOM_VIEW)
	if err != nil {
		return err
	}

	bottomV.Clear()
	//fmt.Fprintln(bottomV, "maximizePreview")

	maxPreviewV, err := g.SetView("MAXIMIZE_PREVIEW", 1, 1, 40, 20)
	maxPreviewV.Wrap = true
	maxPreviewV.Autoscroll = true
	maxPreviewV.Editable = true

	fmt.Fprintln(bottomV, err)

	_, err = setCurrentViewOnTop(g, "MAXIMIZE_PREVIEW")
	//_, err = g.SetViewOnTop("MAXIMIZE_PREVIEW")
	//setCurrentViewOnTop(g, PREVIEW_VIEW)
	return err
}

func nextView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (active + 1) % len(viewArr)
	name := viewArr[nextIndex]

	out, err := g.View(BOTTOM_VIEW)
	if err != nil {
		return err
	}

	out.Clear()
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

	// main section (contains accounts, emails, preview)
	mainViewStartX := 0
	mainViewStartY := 0
	mainViewEndX := winX - 1
	mainViewEndY := winY - 4
	// bottom section
	bottomStartX := 0
	bottomStartY := mainViewEndY + 1
	bottomEndX := winX - 1
	bottomEndY := winY - 1

	// accounts view
	accountsStartX := mainViewStartX
	accountsStartY := mainViewStartY
	accountsEndX := mainViewEndX / 10 * 2
	accountsEndY := mainViewEndY
	// emails list view
	emailsStartX := accountsEndX + 1
	emailsStartY := mainViewStartY
	emailsEndX := mainViewEndX / 10 * 5
	emailsEndY := mainViewEndY
	// email preview view
	previewStartX := emailsEndX + 1
	previewStartY := mainViewStartY
	previewEndX := mainViewEndX
	previewEndY := mainViewEndY

	if accountsV, err := gui.SetView(ACCOUNTS_VIEW, accountsStartX, accountsStartY, accountsEndX, accountsEndY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		accountsV.Title = strconv.Itoa(accountsStartX) + " - " + strconv.Itoa(accountsEndX) + " Accounts"
		//accountsV.Editable = true
		accountsV.Wrap = true

		accountsV.Highlight = true
		accountsV.SelBgColor = gocui.ColorGreen
		accountsV.SelFgColor = gocui.ColorBlack

		for k, v := range Items {
			fmt.Fprintln(accountsV, strconv.Itoa(k)+" - "+v)
		}

		if _, err = setCurrentViewOnTop(gui, ACCOUNTS_VIEW); err != nil {
			return err
		}

		if err := gui.SetKeybinding(ACCOUNTS_VIEW, gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
			return err
		}

		if err := gui.SetKeybinding(ACCOUNTS_VIEW, gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
			return err
		}
	}

	if emailsV, err := gui.SetView(EMAILS_VIEW, emailsStartX, emailsStartY, emailsEndX, emailsEndY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		emailsV.Title = strconv.Itoa(emailsStartX) + " - " + strconv.Itoa(emailsEndX) + " Emails"
		emailsV.Wrap = true
		emailsV.Autoscroll = true
		emailsV.Editable = true
	}
	if previewV, err := gui.SetView(PREVIEW_VIEW, previewStartX, previewStartY, previewEndX, previewEndY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		previewV.Title = strconv.Itoa(previewStartX) + " - " + strconv.Itoa(previewEndX) + " Preview"
		previewV.Editable = true
	}
	if bottomV, err := gui.SetView(BOTTOM_VIEW, bottomStartX, bottomStartY, bottomEndX, bottomEndY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		bottomV.Title = strconv.Itoa(bottomStartX) + " - " + strconv.Itoa(bottomEndX) + " Log"
		bottomV.Editable = true
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

var Gui *gocui.Gui

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

	if err := g.SetKeybinding("", gocui.KeyCtrlZ, gocui.ModNone, maximizePreview); err != nil {
		log.Panicln(err)
	}
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

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		showLog(g, "Cursor down: "+Items[cy])
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

func cursorUp(g *gocui.Gui, v *gocui.View) error {
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

func showLog(g *gocui.Gui, message string) {

	logV, _ := g.View(EMAILS_VIEW)
	logV.Clear()
	fmt.Fprintln(logV, "Cursor up: "+message)
}
