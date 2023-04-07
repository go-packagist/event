package event

// Dispatcher event dispatcher
type Dispatcher struct {
	Listeners map[string][]Listener
}

// NewDispatcher create new dispatcher
func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		Listeners: make(map[string][]Listener, 0),
	}
}

// Listen add listener to event
func (d *Dispatcher) Listen(event Event, listener Listener) {
	d.Listeners[event.Name()] = append(d.Listeners[event.Name()], listener)
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
	return d.Listeners[event.Name()]
}

// Flush remove all listeners of event
func (d *Dispatcher) Flush(event Event) {
	delete(d.Listeners, event.Name())
}

// FlushAll remove all listeners
func (d *Dispatcher) FlushAll() {
	d.Listeners = make(map[string][]Listener, 0)
}
