package solution_test

import (
	"testing"

	"github.com/AlexsJones/devops-go-problems/mailroom/solution"
)

func TestingLoadingAllBundles(t *testing.T) {
	bundles := []string{
		"../letters/json-bundle-01",
		"../letters/json-bundle-02",
		"../letters/json-bundle-03",
	}
	for _, bundle := range bundles {
		mailroom := solution.NewMailroom()
		if _, err := mailroom.ProcessLetter(bundle); err != nil {
			t.Errorf("Unable to process letters due to: %v", err)
		}
		if len(mailroom.UndeliverablePost()) != 0 {
			t.Errorf("All mail should be deliverable for %v", bundle)
		}
	}
}

func TestAbbreviatedDataBundles(t *testing.T) {
	bundles := []string{
		"../letters/json-abbrv-01",
		"../letters/json-abbrv-02",
		"../letters/json-abbrv-03",
	}
	for _, bundle := range bundles {
		mailroom := solution.NewMailroom()
		if _, err := mailroom.ProcessLetter(bundle); err != nil {
			t.Errorf("Unable to process %v due to: %v", bundle, err)
		}
	}
}

func TestMixedDataBundles(t *testing.T) {
	bundles := []string{
		"../letters/json-mixed-01",
		"../letters/json-mixed-02",
		"../letters/json-mixed-03",
	}
	for _, bundle := range bundles {
		mailroom := solution.NewMailroom()
		if _, err := mailroom.ProcessLetter(bundle); err != nil {
			t.Errorf("Unable to process %v due to: %v", bundle, err)
		}
	}
}
