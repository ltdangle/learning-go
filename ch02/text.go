package main

import "fmt"

func main() {
	aString := "Hello World! €"
	fmt.Println("First character", string(aString[0]))

	// Runes
	r := '€'
	fmt.Println("As an int32 value", r)
	fmt.Printf("As a string: %s and as a character: %c\n", r, r)

	// Print an existing string as Runes
	for _, v := range aString {
		fmt.Printf("%x ", v)
	}
	fmt.Println()

}
