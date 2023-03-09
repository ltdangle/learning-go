package main

import (
	"learngocui/repository"
	"learngocui/tui"
)

func main() {
	// Add some emails to the slice
	emailAccounts := repository.SeedData()
	tui.Init(emailAccounts)
}
