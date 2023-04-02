package event

// Dispatcher event dispatcher
type Dispatcher struct {
	Listeners map[Event][]Listener
}

// NewDispatcher create new dispatcher
func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		Listeners: make(map[Event][]Listener, 0),
	}
}

// Listen add listener to event
func (d *Dispatcher) Listen(event Event, listener Listener) {
	d.Listeners[event] = append(d.Listeners[event], listener)
}

// Dispatch event to all listeners
func (d *Dispatcher) Dispatch(event Event) {
	for _, listener := range d.GetListeners(event) {
		if event.IsStop() {
			return
		}

		listener.Handle(event)
	}
}

// GetListeners return all listeners of event
func (d *Dispatcher) GetListeners(event Event) []Listener {
	return d.Listeners[event]
}

// Flush remove all listeners of event
func (d *Dispatcher) Flush(event Event) {
	delete(d.Listeners, event)
}

// FlushAll remove all listeners
func (d *Dispatcher) FlushAll() {
	d.Listeners = make(map[Event][]Listener)
}
