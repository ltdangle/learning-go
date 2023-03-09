package main

import (
	"learngocui/model"
	"learngocui/tui"
	"learngocui/tui/events"
	"learngocui/tui/vm"
)

func main() {
	e := events.NewEventManager()
	viewModel := vm.NewStore(e)
	seed := model.SeedData()
	viewModel.SetAccounts(seed)

	tui.Init(e, viewModel)
}
