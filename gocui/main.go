package main

import (
	"learngocui/events"
	"learngocui/model"
	"learngocui/tui"
	"learngocui/vm"
)

func main() {
	e := events.NewEventManager()
	viewModel := vm.NewStore(e)
	seed := model.SeedData()
	viewModel.SetAccounts(seed)

	tui.Init(e, viewModel)
}
