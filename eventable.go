package event

// Eventable abstracts the event
type Eventable struct {
}

// IsStop is a method of Eventable
func (e *Eventable) IsStop() bool {
	return false
}
