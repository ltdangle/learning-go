//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func InitializeEvent() Event {
	wire.Build(NewPerson, NewEvent, NewGreeter, NewMessage, NewGeometry)
	return Event{}
}
