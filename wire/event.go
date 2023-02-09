package main

import "fmt"

type Event struct {
	Greeter  Greeter // <- adding a Greeter field
	Person   Person
	Geometry geometry
}

func NewEvent(g Greeter, p Person, gm geometry) Event {
	return Event{Greeter: g, Person: p, Geometry: gm}
}
func (this Event) Start() {
	msg := this.Greeter.Greet()
	fmt.Println(msg)
	this.Person.whoami()
	this.measure()
}
func (this Event) measure() {
	fmt.Println(this.Geometry)
	fmt.Println(this.Geometry.area())
	fmt.Println(this.Geometry.perim())
}
