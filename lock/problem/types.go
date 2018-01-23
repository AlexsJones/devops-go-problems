package problem

const (
	// One represents the HID code
	// for 1 on the keyboard
	One rune = 89 + iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Zero
)

// Lock ...
type Lock interface {
	// EnteredKey recieves keystrokes from a keypad
	// and should update its state if needed.
	EnteredKey(k rune) error
	// Initialise will reset the lock to original
	// state described
	Initialise()

	GetDisplay() string
	// Tick is called every 1 second
	Tick()
}
