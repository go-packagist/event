package event

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestEvent struct {
	Stoped bool
}

var _ Event = (*TestEvent)(nil)

func (t *TestEvent) IsStop() bool {
	return t.Stoped
}

func (t *TestEvent) Stop() {
	t.Stoped = true
}

type Test1Listener struct {
	Val  string
	stop bool
}

var _ Listener = (*Test1Listener)(nil)

func (t *Test1Listener) Handle(event Event) {
	t.Val = "Test1 Done"

	if t.stop {
		event.(*TestEvent).Stop()
	}
}

func (t *Test1Listener) Stop() {
	t.stop = true
}

type Test2Listener struct {
	Val string
}

var _ Listener = (*Test2Listener)(nil)

func (t *Test2Listener) Handle(event Event) {
	t.Val = "Test2 Done"
}

func TestDispatcher_Dispatch(t *testing.T) {
	d := NewDispatcher()

	// event finish
	event1, listener11, listener12 := &TestEvent{}, &Test1Listener{}, &Test2Listener{}

	d.Listen(event1, listener11)
	d.Listen(event1, listener12)

	d.Dispatch(event1)

	assert.Equal(t, "Test1 Done", listener11.Val)
	assert.Equal(t, "Test2 Done", listener12.Val)

	// event stop
	event2, listener21, listener22 := &TestEvent{}, &Test1Listener{}, &Test2Listener{}

	d.Listen(event2, listener21)
	d.Listen(event2, listener22)

	listener21.Stop()

	d.Dispatch(event2)

	assert.Equal(t, "Test1 Done", listener21.Val)
	assert.Equal(t, "", listener22.Val)
}

func TestFlush(t *testing.T) {
	d := NewDispatcher()

	event1, event2, listener1, listener2 := &TestEvent{}, &TestEvent{}, &Test1Listener{}, &Test2Listener{}

	d.Listen(event1, listener1)
	d.Listen(event1, listener2)

	d.Listen(event2, listener1)
	d.Listen(event2, listener2)

	assert.Equal(t, 2, len(d.GetListeners(event1)))
	assert.Equal(t, 2, len(d.GetListeners(event2)))

	// Flush
	d.Flush(event1)

	assert.Equal(t, 0, len(d.GetListeners(event1)))
	assert.Equal(t, 2, len(d.GetListeners(event2)))

	// FlushAll
	d.Listen(event1, listener1)
	d.Listen(event1, listener2)

	assert.Equal(t, 2, len(d.GetListeners(event1)))
	assert.Equal(t, 2, len(d.GetListeners(event2)))

	d.FlushAll()

	assert.Equal(t, 0, len(d.GetListeners(event1)))
	assert.Equal(t, 0, len(d.GetListeners(event2)))
}

func TestUniversalEvent(t *testing.T) {

	d := NewDispatcher()

	// u1
	u1, listener1, listener2 := &UniversalEvent{}, &Test1Listener{}, &Test2Listener{}
	assert.False(t, u1.IsStop())

	d.Listen(u1, listener1)
	d.Listen(u1, listener2)
	d.Dispatch(u1)

	assert.Equal(t, "Test1 Done", listener1.Val)
	assert.Equal(t, "Test2 Done", listener2.Val)

	// u2
	u2, listener3 := &UniversalEvent{}, &Test1Listener{}

	d.Listen(u2, listener3)
	d.Dispatch(u2)

	assert.Equal(t, "Test1 Done", listener3.Val)
}
