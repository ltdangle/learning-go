package main

import (
	"learngocui/tui"
	"learngocui/tui/events"
	"learngocui/tui/vm"
)

func main() {
	e := events.NewEventManager()
	viewModel := vm.NewVM(e)
	//seed := model.SeedData()

	tui.Init(e, viewModel)
}
