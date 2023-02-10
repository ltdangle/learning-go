//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func InitializeEvent() Event {
	wire.Build(
		NewPerson,
		NewEvent,
		NewGreeter,
		NewMessage,
		wire.Bind(new(GeometryI), new(circle)),
		NewCircle)
	return Event{}
}
