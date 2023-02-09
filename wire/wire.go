
//+build wireinject

package main
import "github.com/google/wire"
    
func InitializeEvent() Event {
    wire.Build(NewPerson, NewEvent, NewGreeter, NewMessage)
    return Event{}
}
