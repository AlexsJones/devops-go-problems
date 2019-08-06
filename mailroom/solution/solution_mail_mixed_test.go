package solution_test

import (
	"testing"

	"github.com/AlexsJones/devops-go-problems/mailroom/solution"
)

func TestMixedContentBundle(t *testing.T) {
	mailroom := solution.NewMailroom()
	letters, err := mailroom.ProcessLetter("../letters/json-mixed-01")
	if err != nil {
		t.Errorf("Unable to open json-mixed-01 due to: %v", err)
		t.FailNow()
	}
	if len(letters) != 10 {
		t.Error("Should of loaded all 10 letters")
		t.FailNow()
	}
	// calculated totals for the basic bundle
	counts := map[string]int{
		"New South Wales":   1,
		"Victoria":          1,
		"Queensland":        0,
		"South Australia":   0,
		"Western Australia": 0,
		"Tasmania":          0,
	}
	for state, total := range counts {
		calculated, err := mailroom.GetDeliveryCount(state)
		if err != nil {
			t.Errorf("Failed to get count for %v", state)
		}
		if total != calculated {
			t.Errorf("Missmatch total for %v, Expected %v Got %v", state, total, calculated)
		}
	}
	if len(mailroom.UndeliverablePost()) != 8 {
		t.Errorf("Undeliverable post fail, expected 8 got %v", len(mailroom.UndeliverablePost()))
	}
}
