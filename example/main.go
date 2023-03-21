package main

import "github.com/go-packagist/event"

type Event struct {
	Stop bool
}

func (e *Event) IsStop() bool {
	return e.Stop
}

func (e *Event) Val() string {
	return "event"
}

type Listener1 struct {
}

func (l *Listener1) Handle(event event.Event) {
	println("listener1:" + event.(*Event).Val())

	event.(*Event).Stop = true
}

type Listener2 struct{}

func (l *Listener2) Handle(event event.Event) {
	println("listener2:" + event.(*Event).Val())
}

var _ event.Event = (*Event)(nil)
var _ event.Listener = (*Listener1)(nil)
var _ event.Listener = (*Listener2)(nil)

func main() {
	d := event.NewDispatcher()

	e := &Event{
		Stop: false,
	}

	d.Listen(e, &Listener1{})
	d.Listen(e, &Listener2{})

	d.Dispatch(e) // echo: listener1:event (because listener1 set Stop to true)
}