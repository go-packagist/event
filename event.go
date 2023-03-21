package event

// Event interface(see psr-14:https://github.com/php-fig/event-dispatcher)
type Event interface {
	IsStop() bool
}
