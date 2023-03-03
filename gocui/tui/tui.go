// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tui

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"learngocui/model"
	"log"
)

const (
	ACCOUNTS_VIEW = "accounts view"
	EMAILS_VIEW   = "emails view"
	PREVIEW_VIEW  = "preview view"
	BOTTOM_VIEW   = "bottom view"
)

var (
	viewArr  = []string{ACCOUNTS_VIEW, EMAILS_VIEW, PREVIEW_VIEW, BOTTOM_VIEW}
	active   = 0
	Accounts = []string{"one", "two", "three", "four", "five"}
	Emails   = [][]string{
		{"email1@one.com", "email2@one.com", "email3@one.com", "email4@one.com", "email5@one.com"},
		{"email1@two.com", "email2@two.com", "email3@two.com", "email4@two.com", "email5@two.com"},
		{"email1@three.com", "email2@three.com", "email3@three.com", "email4@three.com", "email5@three.com"},
		{"email1@four.com", "email2@four.com", "email3@four.com", "email4@four.com", "email5@four.com"},
		{"email1@five.com", "email2@five.com", "email3@five.com", "email4@five.com", "email5@five.com"},
	}
)

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

var Gui *gocui.Gui

func Tui(emails []model.EmailAccount) {

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
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, goToNextView); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func showLog(g *gocui.Gui, message string) {

	logV, _ := g.View(PREVIEW_VIEW)
	logV.Clear()
	fmt.Fprintln(logV, message)
}
