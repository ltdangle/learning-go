// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tui

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"learngocui/tui/events"
	"learngocui/tui/logger"
	"learngocui/tui/vm"
	"log"
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

type Tui struct {
	gui      *gocui.Gui
	accounts *accounts
	emails   *emails
	preview  *preview
	bottom   *bottom
	events   *events.EventManager
}

var T = &Tui{}

func Init(e events.IEvent, vm *vm.ViewModel, logger logger.ILogger) {

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()
	T.gui = g

	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorMagenta

	T.accounts = newAccounts(e, vm)
	T.emails = newEmails(e, vm)
	T.preview = newPreview(vm)
	T.bottom = newBottom()

	EventListeners(T, vm, logger)

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

func tuiLog(message string) {
	logV, _ := T.gui.View(BOTTOM_VIEW)
	logV.Clear()
	fmt.Fprintln(logV, message)
}
