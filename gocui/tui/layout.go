package tui

import (
	"github.com/jroimartin/gocui"
	"learngocui/repository"
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
	accountRepository := repository.NewSeedAccountRepository(repository.SeedData())

	if err = T.AccountsV.initView(gui, accountsStartX, accountsStartY, accountsEndX, accountsEndY); err != nil {
		return err
	}

	if err = T.EmailsV.initView(gui, emailsStartX, emailsStartY, emailsEndX, emailsEndY); err != nil {
		return err
	}

	if T.PreviewV, err = createPreviewView(gui, accountRepository, previewStartX, previewStartY, previewEndX, previewEndY); err != nil {
		return err
	}

	if T.BottomV, err = createBottomView(gui, bottomStartX, bottomStartY, bottomEndX, bottomEndY); err != nil {
		return err
	}

	return nil
}
