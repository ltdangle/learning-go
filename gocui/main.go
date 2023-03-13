package main

import (
	"learngocui/model"
	"learngocui/tui"
	"learngocui/tui/events"
	"learngocui/tui/vm"
)

func main() {
	e := events.NewEventManager()
	seed := model.SeedData()

	acc1 := vm.NewAccountVM(e, &seed[0])
	acc2 := vm.NewAccountVM(e, &seed[1])
	acc3 := vm.NewAccountVM(e, &seed[2])
	acc4 := vm.NewAccountVM(e, &seed[3])

	viewModel := vm.NewVM(e)
	viewModel.AddAccount(acc1)
	viewModel.AddAccount(acc2)
	viewModel.AddAccount(acc3)
	viewModel.AddAccount(acc4)

	tui.Init(e, viewModel)
}
