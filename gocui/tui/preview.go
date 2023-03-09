package tui

import (
	"fmt"
	"github.com/jroimartin/gocui"
	"learngocui/model"
	"strconv"
)

type preview struct {
	view  *gocui.View
	email *model.Email
}

func newPreview(email *model.Email) *preview {
	return &preview{email: email}
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
	self.view.Clear()
	fmt.Fprintln(self.view, "Date: "+self.email.Date)
	fmt.Fprintln(self.view, "From: "+self.email.From)
	fmt.Fprintln(self.view, "To: "+self.email.To)
	fmt.Fprintln(self.view, "Subject: "+self.email.Subject)
	fmt.Fprintln(self.view, "")
	fmt.Fprintln(self.view, self.email.Text)
}
