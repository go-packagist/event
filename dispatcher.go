package event

import "sync"

// Dispatcher event dispatcher
type Dispatcher struct {
	listeners map[string][]Listener

	mu sync.Mutex
}

// NewDispatcher create new dispatcher
func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		listeners: make(map[string][]Listener, 0),
	}
}

// Listen add listener to event
func (d *Dispatcher) Listen(name string, listener ...Listener) {
	if _, ok := d.listeners[name]; !ok {
		d.mu.Lock()
		defer d.mu.Unlock()

		d.listeners[name] = make([]Listener, 0, len(listener))
	}

	d.listeners[name] = append(d.listeners[name], listener...)
}

// Dispatch event to all listeners
func (d *Dispatcher) Dispatch(event Event) {
	for _, listener := range d.GetListeners(event.Name()) {
		if event.(StoppableEvent) != nil && event.(StoppableEvent).IsPropagationStopped() {
			return
		}

		listener.Handle(event)
	}
}

// GetListeners return all listeners of event
func (d *Dispatcher) GetListeners(name string) []Listener {
	if listeners, ok := d.listeners[name]; ok {
		return listeners
	}

	return []Listener{}
}

// Flush remove listeners of event
func (d *Dispatcher) Flush(name string) {
	if _, ok := d.listeners[name]; ok {
		delete(d.listeners, name)
	}
}

// FlushAll remove all listeners
func (d *Dispatcher) FlushAll() {
	d.listeners = make(map[string][]Listener, 0)
}
