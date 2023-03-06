package tui

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"learngocui/repository"
	"strconv"
)

type previewV struct {
	view         *gocui.View
	accountsRepo repository.IAccountRepository
}

func createPreviewView(gui *gocui.Gui, accountsRepo repository.IAccountRepository, startX, startY, endX, endY int) (*previewV, error) {
	var err error
	self := &previewV{}
	self.accountsRepo = accountsRepo

	if self.view, err = gui.SetView(PREVIEW_VIEW, startX, startY, endX, endY); err != nil {
		if err != gocui.ErrUnknownView {
			return nil, err
		}
		self.view.Title = strconv.Itoa(startX) + " - " + strconv.Itoa(endX) + " Preview"
		self.view.Editable = true
		self.view.Autoscroll = true
		self.view.Wrap = true
	}
	return self, nil
}
func (self *previewV) populate(g *gocui.Gui, emailIndex int) error {
	v, _ := g.View(PREVIEW_VIEW)
	v.Clear()
	fmt.Fprintln(v, "show email "+strconv.Itoa(emailIndex))
	//for _, email := range self.accountsRepo.FindById(emailAccountIndex).Emails {
	//	fmt.Fprintln(v, email)
	//}

	return nil
}
