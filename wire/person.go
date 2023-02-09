package main

import "fmt"

type Person struct {
}

func (this Person) whoami() {
	fmt.Println("I am a Person")
}
func NewPerson() Person {
	return Person{}
}
