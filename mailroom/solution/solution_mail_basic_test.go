package solution_test

import (
	"os"
	"testing"

	"github.com/SeedJobs/devops-go-problems/mailroom/solution"
	"github.com/SeedJobs/devops-go-problems/mailroom/types"
)

func TestBasicFunctionality(t *testing.T) {
	mailroom := solution.NewMailroom()
	// We have created a basic mailroom with no mail being sent
	// All counts must be set to zero
	for state := range types.Abbreviate {
		count, err := mailroom.GetDeliveryCount(state)
		if err != nil {
			t.Errorf("Failed to get count for key %s", state)
		}
		if count != 0 {
			t.Error("Count should be equal to zero")
		}
	}
	failedmail := mailroom.UndeliverablePost()
	if failedmail == nil {
		t.Errorf("UndeliverablePost should return an array")
	}
	if len(failedmail) != 0 {
		t.Errorf("There shouldn't be any failed delieviers")
	}
}

func TestBadBasicFunctionality(t *testing.T) {
	mailroom := solution.NewMailroom()
	if _, err := mailroom.GetDeliveryCount(""); err == nil {
		t.Error("Should return an error if the string doesn't exist in Map")
	}
	if _, err := mailroom.ProcessLetter("notexist"); err == nil {
		t.Error("Failed to check if file exists")
	}
	// Having to set the bad file permissions during test as it will
	// fail to load to git if it misses r
	badfile, err := os.Open("../letters/bad-file")
	if err != nil {
		t.Error("Unable to load file due to:", err)
	}
	defer badfile.Chmod(0640)
	if err = badfile.Chmod(240); err != nil {
		t.Error("Unable to remove r due to:", err)
	}
	if _, err = mailroom.ProcessLetter("../letters/bad-file"); err == nil {
		t.Error("Mailroom should not be able to load bad-file")
	}
}

func TestLoadingBasicMailLoading(t *testing.T) {
	mailroom := solution.NewMailroom()
	letters, err := mailroom.ProcessLetter("../letters/json-bundle-01")
	if err != nil {
		t.Errorf("Was unable to load bundle and obtained %v", err)
	}
	if len(letters) != 20 {
		t.Error("Failed to load all the letters in the bundle")
	}
	// calculated totals for the basic bundle
	counts := map[string]int{
		"New South Wales":   13,
		"Victoria":          0,
		"Queensland":        1,
		"South Australia":   2,
		"Western Australia": 1,
		"Tasmania":          3,
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
	if len(mailroom.UndeliverablePost()) != 0 {
		t.Error("All mail is deliverable in this bundle")
	}
}
