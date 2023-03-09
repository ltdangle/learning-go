// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tui

import (
	"fmt"
	"github.com/gookit/event"
	"github.com/jroimartin/gocui"
	"learngocui/events"
	"learngocui/model"
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
	Gui       *gocui.Gui
	AccountsV *accounts
	EmailsV   *emailsV
	PreviewV  *previewV
	BottomV   *bottomV
	Events    *events.EventManager
}

var T = &Tui{}

func Init(emails []model.EmailAccount) {

	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	T.Gui = g
	defer g.Close()

	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorMagenta

	eventManager := events.CreateTuiEventManager(g)
	T.AccountsV = newAccountsV(eventManager, []string{"one", "two", "three"})

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

	s := store.NewStore(T.Events)
	seed := repository.SeedData()
	s.SetAccounts(seed)
	fmt.Print(T)

	eventListeners()

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
func eventListeners() {
	event.On(ACCOUNTS_CURSOR_DOWN_EVENT, event.ListenerFunc(func(e event.Event) error {
		selectedItem := e.Data()["selectedItem"].(int)
		tuiLog(T.Gui, "handle event from eventManager: "+EMAILS_CURSOR_DOWN_EVENT+", selectedItem: "+strconv.Itoa(selectedItem))

		return nil
	}), event.Normal)

	event.On(ACCOUNTS_CURSOR_UP_EVENT, event.ListenerFunc(func(e event.Event) error {
		selectedItem := e.Data()["selectedItem"].(int)
		tuiLog(T.Gui, "handle event from eventManager: "+EMAILS_CURSOR_UP_EVENT+", selectedItem: "+strconv.Itoa(selectedItem))
		return nil
	}), event.Normal)
}
func tuiLog(g *gocui.Gui, message string) {

	logV, _ := g.View(BOTTOM_VIEW)
	logV.Clear()
	fmt.Fprintln(logV, message)
}
