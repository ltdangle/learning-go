package tui

import (
	"fmt"
	"github.com/jroimartin/gocui"
)

func goToNextView(g *gocui.Gui, v *gocui.View) error {
	nextIndex := (active + 1) % len(viewArr)
	name := viewArr[nextIndex]

	out, err := g.View(BOTTOM_VIEW)
	if err != nil {
		return err
	}

	out.Clear()
	fmt.Fprintln(out, "Going from view "+v.Name()+" to "+name)

	if _, err := setCurrentViewOnTop(g, name); err != nil {
		return err
	}

	active = nextIndex
	return nil
}
