package tui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func maximizePreview(g *gocui.Gui, v *gocui.View) error {
	bottomV, err := g.View(BOTTOM_VIEW)
	if err != nil {
		return err
	}

	bottomV.Clear()
	//fmt.Fprintln(bottom, "maximizePreview")

	maxPreviewV, err := g.SetView("MAXIMIZE_PREVIEW", 1, 1, 40, 20)
	maxPreviewV.Wrap = true
	maxPreviewV.Autoscroll = true
	maxPreviewV.Editable = true

	fmt.Fprintln(bottomV, err)

	_, err = setCurrentViewOnTop(g, "MAXIMIZE_PREVIEW")
	//_, err = gui.SetViewOnTop("MAXIMIZE_PREVIEW")
	//setCurrentViewOnTop(gui, PREVIEW_VIEW)
	return err
}
