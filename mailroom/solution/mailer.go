package solution

import (

  "github.com/SeedJobs/devops-go-problems/mailroom/types"
)

type mailer struct {
  // Add any additional content here
}

func NewMailroom() types.Mailroom {
  return &mailer{}
}

func (m *mailer) ProcessLetter(address []byte) types.Letter{
  // add logic here
  return &post{}
}

func (m *mailer) GetDeliveryCount(state string) (int, error) {
  // add logic here
  return 0, nil
}

func (m *mailer) UndeliverablePost() []types.Letter {
  return nil
}
