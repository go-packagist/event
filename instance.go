package event

var instance *Dispatcher

func GetDispatcher() *Dispatcher {
	if instance == nil {
		instance = NewDispatcher()
	}

	return instance
}

func Listen(event Event, listener Listener) {
	GetDispatcher().Listen(event, listener)
}

func Dispatch(event Event) {
	GetDispatcher().Dispatch(event)
}
