package types

// Mailroom is service abstraction of how the mailroom is to process
// letters being sent int.
type Mailroom interface {
  // ProcessLetter accepts the adress of the letter to be sent
  // and determines what state it is going to.
  // If its unable to be delievered, then it should report an error
  ProcessLetter(address []byte) Letter

  // GetDeliveryCount returns the total amount of letters sent to
  // the given state during that "day" of work.
  GetDeliveryCount(state string) (int, error)

  // UndeliverablePost returns all letters that could not
  // be sent to a valid address and needs to be returned.
  UndeliverablePost() []Letter
}
