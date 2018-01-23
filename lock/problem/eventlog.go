package problem

// EventLog used to print messages to the
// UI. No added benefit
type EventLog struct {
	buff []string
}

// NewEventLog returns an initialised EventLog
func NewEventLog() *EventLog {
	return &EventLog{
		buff: []string{},
	}
}

// Add will append a new event to the buffer
// after the length has reached 7 it will cycle
// out any old logs.
func (log *EventLog) Add(event string) {
	if len(log.buff) >= 7 {
		log.buff = log.buff[1:]
	}
	log.buff = append(log.buff, event)
}

// GetBuff returns the internal buffer
func (log *EventLog) GetBuff() []string {
	return log.buff
}
