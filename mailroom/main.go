package main

import (
	"fmt"

	"github.com/SeedJobs/devops-go-problems/mailroom/solution"
	"github.com/SeedJobs/devops-go-problems/mailroom/types"
)

func main() {
	mailroom := solution.NewMailroom()
	if _, err := mailroom.ProcessLetter("letters/json-bundle-01"); err != nil {
		fmt.Println("Unable to process letters because:", err)
	}
	for state := range types.Abbreviate {
		count, err := mailroom.GetDeliveryCount(state)
		if err != nil {
			fmt.Println("Unable to show delivery count for ", state)
			fmt.Println("Due to:", err)
			continue
		}
		fmt.Printf("Sending %v to %v\n", count, state)
	}
	for _, post := range mailroom.UndeliverablePost() {
		fmt.Println("Can not post:", post)
	}
}
