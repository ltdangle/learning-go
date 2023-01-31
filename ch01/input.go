package main

import "fmt"

func main() {
	// get user input
	fmt.Printf("Please give your name: ")
	var name string
	fmt.Scanln(&name)
	fmt.Println("Your name is", name)

}
