package main

import "fmt"

type Event struct {
	Greeter Greeter // <- adding a Greeter field
    Person Person
}

func NewEvent(g Greeter, p Person) Event {
	return Event{Greeter: g}
}
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
    e.Person.whoami()
}
