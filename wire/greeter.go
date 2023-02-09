package main

type Greeter struct {
	Message Message // <- adding a Message field
}
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}
func (this Greeter) Greet() Message {
	return this.Message
}
