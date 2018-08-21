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

func (m *mailer) ProcessLetter(bundlepath string) ([]types.Letter, error) {
	// add logic here
	return nil, nil
}

func (m *mailer) GetDeliveryCount(state string) (int, error) {
	// add logic here
	return 0, nil
}

func (m *mailer) UndeliverablePost() []types.Letter {
	return nil
}
