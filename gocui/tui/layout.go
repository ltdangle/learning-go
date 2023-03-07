package tui

import (
	"github.com/gookit/event"
	"github.com/jroimartin/gocui"
	"learngocui/events"
	"learngocui/repository"
	"learngocui/store"
	"strconv"
)

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

	eventManager := events.CreateTuiEventManager(gui)

	var err error
	var emailsV *emailsV
	var previewV *previewV
	accountRepository := repository.NewSeedAccountRepository(repository.SeedData())

	s := store.NewStore(eventManager)
	seed := repository.SeedData()
	s.SetAccounts(seed)

	if emailsV, err = createEmailsView(eventManager, gui, accountRepository, emailsStartX, emailsStartY, emailsEndX, emailsEndY); err != nil {
		return err
	}
	if _, err = createAccountsView(eventManager, gui, emailsV, accountsStartX, accountsStartY, accountsEndX, accountsEndY); err != nil {
		return err
	}

	if previewV, err = createPreviewView(gui, accountRepository, previewStartX, previewStartY, previewEndX, previewEndY); err != nil {
		return err
	}

	if _, err = createBottomView(gui, bottomStartX, bottomStartY, bottomEndX, bottomEndY); err != nil {
		return err
	}

	// Register event listeners
	event.On(events.UPDATE_EMAILS_VIEW, event.ListenerFunc(func(e event.Event) error {
		selectedItem := e.Data()["selectedItem"].(int)
		tuiLog(gui, "handle event from eventManager: "+events.UPDATE_EMAILS_VIEW+", selectedItem: "+strconv.Itoa(selectedItem))
		_ = emailsV.populate(gui, selectedItem)
		return nil
	}), event.Normal)

	event.On(events.UPDATE_EMAIL_PREVIEW, event.ListenerFunc(func(e event.Event) error {
		selectedItem := e.Data()["selectedItem"].(int)
		tuiLog(gui, "handle event from eventManager: "+events.UPDATE_EMAIL_PREVIEW+", selectedItem: "+strconv.Itoa(selectedItem))
		_ = previewV.populate(gui, selectedItem)
		return nil
	}), event.Normal)

	return nil
}
