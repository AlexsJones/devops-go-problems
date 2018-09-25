package solution_test

import (
	"testing"

	"github.com/SeedJobs/devops-go-problems/mailroom/solution"
	"github.com/SeedJobs/devops-go-problems/mailroom/types"
)

func TestAbbreviationUsage(t *testing.T) {
	mailroom := solution.NewMailroom()
	letters, err := mailroom.ProcessLetter("../letters/json-abbrv-01")
	if err != nil {
		t.Errorf("Was unable to load bundle and obtained %v", err)
	}
	if len(letters) != 10 {
		t.Error("Failed to load all the letters in the bundle")
	}
	// calculated totals for the basic bundle
	counts := map[string]int{
		"New South Wales":   3,
		"Victoria":          0,
		"Queensland":        0,
		"South Australia":   1,
		"Western Australia": 2,
		"Tasmania":          4,
	}
	for state, total := range counts {
		calculated, err := mailroom.GetDeliveryCount(state)
		if err != nil {
			t.Errorf("Failed to get count for %v", state)
		}
		if total != calculated {
			t.Errorf("Missmatch total for %v, Expected %v Got %v", state, total, calculated)
		}
		calculated, err = mailroom.GetDeliveryCount(types.Expanded[state])
		if err != nil {
			t.Errorf("Failed to get count for %v", state)
		}
		if total != calculated {
			t.Errorf("Missmatch total for %v, Expected %v Got %v", state, total, calculated)
		}
	}
	if len(mailroom.UndeliverablePost()) != 0 {
		t.Error("All mail is deliverable in this bundle")
	}
}
