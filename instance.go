package event

var instance *Dispatcher

func GetDispatcher() *Dispatcher {
	if instance == nil {
		instance = NewDispatcher()
	}

	return instance
}

func Listen(name string, listener Listener) {
	GetDispatcher().Listen(name, listener)
}

func Dispatch(event Event) {
	GetDispatcher().Dispatch(event)
}
