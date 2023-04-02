package event

import "testing"

func TestDispatch(t *testing.T) {
	event1, listener11, listener12 := &TestEvent{}, &Test1Listener{}, &Test2Listener{}

	Listen(event1, listener11)
	Listen(event1, listener12)

	Dispatch(event1)
}
