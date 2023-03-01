package tui

import (
	"github.com/jroimartin/gocui"
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

	var err error
	//var accountsV *accountsV
	var emailsV *emailsV
	if emailsV, err = createEmailsView(gui, emailsStartX, emailsStartY, emailsEndX, emailsEndY); err != nil {
		return err
	}

	if _, err = createAccountsView(gui, emailsV, accountsStartX, accountsStartY, accountsEndX, accountsEndY); err != nil {
		return err
	}

	if previewV, err := gui.SetView(PREVIEW_VIEW, previewStartX, previewStartY, previewEndX, previewEndY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		previewV.Title = strconv.Itoa(previewStartX) + " - " + strconv.Itoa(previewEndX) + " Preview"
		previewV.Editable = true
		previewV.Autoscroll = true
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
