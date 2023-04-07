package event

// Event interface(see psr-14:https://github.com/php-fig/event-dispatcher)
type Event interface {
	Name() string
	IsStop() bool
}

// Eventable abstracts the event
type Eventable struct{}

// IsStop is a method of Eventable
func (e *Eventable) IsStop() bool {
	return false
}
