package problem

func NewLock(l *EventLog) Lock {
	return &SimpleLock{}
}

type SimpleLock struct {
}

func (l *SimpleLock) Initialise() {

}

func (l *SimpleLock) EnteredKey(key rune) error {
	return nil
}

func (l *SimpleLock) GetDisplay() string {
	return ""
}

func (l *SimpleLock) Tick() {

}
