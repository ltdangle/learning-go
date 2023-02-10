// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func InitializeEvent() Event {
	message := NewMessage()
	greeter := NewGreeter(message)
	person := NewPerson()
	mainCircle := NewCircle()
	event := NewEvent(greeter, person, mainCircle)
	return event
}