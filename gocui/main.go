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

	accVm1 := vm.NewAccountVM(e, &seed[0])
	accVm2 := vm.NewAccountVM(e, &seed[1])
	accVm3 := vm.NewAccountVM(e, &seed[2])
	accVm4 := vm.NewAccountVM(e, &seed[3])

	viewModel := vm.NewVM(e)
	viewModel.AddAccount(accVm1)
	viewModel.AddAccount(accVm2)
	viewModel.AddAccount(accVm3)
	viewModel.AddAccount(accVm4)

	tui.Init(e, viewModel)
}
