// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tui

import (
	"fmt"
	"github.com/gookit/event"
	"github.com/jroimartin/gocui"
	"learngocui/events"
	"learngocui/repository"
	"learngocui/store"
	"log"
	"strconv"
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

func Init() {

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	T.gui = g
	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorMagenta

	vm := store.NewStore(T.events)
	seed := repository.SeedData()
	vm.SetAccounts(seed)

	e := events.CreateTuiEventManager()
	T.accounts = newAccountsV(e, vm.GetAccountNames())
	T.emails = newEmails(e, vm.GetSelectedtAccount().GetEmailsAsList())
	T.preview = newPreview(vm.GetSelectedEmail())
	T.bottom = newBottom()

	eventListeners()

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
func eventListeners() {
	event.On(ACCOUNTS_CURSOR_DOWN_EVENT, event.ListenerFunc(func(e event.Event) error {
		selectedItem := e.Data()["selectedItem"].(int)
		tuiLog("handle event from eventManager: " + EMAILS_CURSOR_DOWN_EVENT + ", selectedItem: " + strconv.Itoa(selectedItem))

		return nil
	}), event.Normal)

	event.On(ACCOUNTS_CURSOR_UP_EVENT, event.ListenerFunc(func(e event.Event) error {
		selectedItem := e.Data()["selectedItem"].(int)
		tuiLog("handle event from eventManager: " + EMAILS_CURSOR_UP_EVENT + ", selectedItem: " + strconv.Itoa(selectedItem))
		return nil
	}), event.Normal)
}

func tuiLog(message string) {
	logV, _ := T.gui.View(BOTTOM_VIEW)
	logV.Clear()
	fmt.Fprintln(logV, message)
}
