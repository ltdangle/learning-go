package tui

import (
	"github.com/jroimartin/gocui"
	"strconv"
)

type bottom struct {
	view *gocui.View
}

func newBottom() *bottom {
	return &bottom{}
}

func (self *bottom) initView(gui *gocui.Gui, startX, startY, endX, endY int) error {
	var err error
	if self.view, err = gui.SetView(BOTTOM_VIEW, startX, startY, endX, endY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		self.view.Title = strconv.Itoa(startX) + " - " + strconv.Itoa(endX) + " tuiLog"
		self.view.Editable = true
	}
	return nil
}
