package types

// Mailroom is service abstraction of how the mailroom is to process
// letters being sent int.
type Mailroom interface {
	// ProcessLetter takes the bundlepath and returns all the letters
	// it read in.
	// If it unable to read letter for a given reason, then it is required to
	// return an error.
	ProcessLetter(bundlepath string) ([]Letter, error)

	// GetDeliveryCount returns the total amount of letters sent to
	// the given state during that "day" of work.
	GetDeliveryCount(state string) (int, error)

	// UndeliverablePost returns all letters that could not
	// be sent to a valid address and needs to be returned.
	UndeliverablePost() []Letter
}
