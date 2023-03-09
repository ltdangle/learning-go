package tui

import (
	"github.com/jroimartin/gocui"
	"strconv"
)

type bottom struct {
	view *gocui.View
}

func createBottomView(gui *gocui.Gui, startX, startY, endX, endY int) (*bottom, error) {
	var err error
	self := &bottom{}
	if self.view, err = gui.SetView(BOTTOM_VIEW, startX, startY, endX, endY); err != nil {
		if err != gocui.ErrUnknownView {
			return nil, err
		}
		self.view.Title = strconv.Itoa(startX) + " - " + strconv.Itoa(endX) + " tuiLog"
		self.view.Editable = true
	}
	return self, nil
}
