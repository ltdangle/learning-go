package main

import (
	"learngocui/tui"
)

func main() {
	// Add some emails to the slice
	emailAccounts := tui.SeedData()
	tui.Tui(emailAccounts)
}
