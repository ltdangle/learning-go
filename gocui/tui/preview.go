package tui

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"learngocui/tui/vm"
	"strconv"
)

type preview struct {
	view      *gocui.View
	viewModel *vm.ViewModel
}

func newPreview(viewModel *vm.ViewModel) *preview {
	return &preview{viewModel: viewModel}
}

func (self *preview) initView(gui *gocui.Gui, startX, startY, endX, endY int) error {
	var err error

	if self.view, err = gui.SetView(PREVIEW_VIEW, startX, startY, endX, endY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		self.view.Title = strconv.Itoa(startX) + " - " + strconv.Itoa(endX) + " Preview"
		self.view.Editable = true
		self.view.Autoscroll = true
		self.view.Wrap = true
	}

	self.populate()

	return nil
}
func (self *preview) populate() {
	selectedEmail := self.viewModel.GetSelectedtAccount().GetSelectedEmail()
	if self.view == nil {
		return
	}
	if selectedEmail == nil {
		return
	}

	self.view.Clear()

	fmt.Fprintln(self.view, "Date: "+selectedEmail.Date)
	fmt.Fprintln(self.view, "From: "+selectedEmail.From)
	fmt.Fprintln(self.view, "To: "+selectedEmail.To)
	fmt.Fprintln(self.view, "Subject: "+selectedEmail.Subject)
	fmt.Fprintln(self.view, "")
	fmt.Fprintln(self.view, selectedEmail.Text)
}
