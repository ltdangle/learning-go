package main

import (
	"learngocui/events"
	"learngocui/repository"
	"learngocui/tui"
	"learngocui/vm"
)

func main() {
	e := events.CreateTuiEventManager()
	vm := vm.NewStore(e)
	seed := repository.SeedData()
	vm.SetAccounts(seed)

	tui.Init(e, vm)
}
