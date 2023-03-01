package tui

import (
	"github.com/jroimartin/gocui"
	"strconv"
)

func createPreviewView(gui *gocui.Gui, startX, startY, endX, endY int) error {
	if previewV, err := gui.SetView(PREVIEW_VIEW, startX, startY, endX, endY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		previewV.Title = strconv.Itoa(startX) + " - " + strconv.Itoa(endX) + " Preview"
		previewV.Editable = true
		previewV.Autoscroll = true
	}
	return nil
}
