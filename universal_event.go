package event

// UniversalEvent universal event
type UniversalEvent struct {
}

// IsStop return true if event is stop
func (e *UniversalEvent) IsStop() bool {
	return false
}

var _ Event = (*UniversalEvent)(nil)
