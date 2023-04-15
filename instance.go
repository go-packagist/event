package event

import "sync"

var instance *Dispatcher
var once sync.Once

func GetDispatcher() *Dispatcher {
	if instance == nil {
		once.Do(func() {
			instance = NewDispatcher()
		})
	}

	return instance
}

func Listen(name string, listener Listener) {
	GetDispatcher().Listen(name, listener)
}

func Dispatch(event Event) {
	GetDispatcher().Dispatch(event)
}
