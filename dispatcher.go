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
func (d *Dispatcher) Listen(name string, listener ...Listener) {
	d.Listeners[name] = append(d.Listeners[name], listener...)
}

// Dispatch event to all listeners
func (d *Dispatcher) Dispatch(event Event) {
	for _, listener := range d.GetListeners(event.Name()) {
		if event.IsStop() {
			return
		}

		listener.Handle(event)
	}
}

// GetListeners return all listeners of event
func (d *Dispatcher) GetListeners(name string) []Listener {
	return d.Listeners[name]
}

// Flush remove listeners of event
func (d *Dispatcher) Flush(name string) {
	delete(d.Listeners, name)
}

// FlushAll remove all listeners
func (d *Dispatcher) FlushAll() {
	d.Listeners = make(map[string][]Listener, 0)
}
