package tui

import (
	"github.com/jroimartin/gocui"
	"strconv"
)

type previewV struct {
	view *gocui.View
}

func createPreviewView(gui *gocui.Gui, startX, startY, endX, endY int) (*previewV, error) {
	var err error
	self := &previewV{}

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
