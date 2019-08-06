package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/AlexsJones/devops-go-problems/mailroom/solution"
	"github.com/AlexsJones/devops-go-problems/mailroom/types"
)

const (
	content = `
%v works by loading processing the bundles paths when starting the application.
It will print out the statistics for each state that had deliverable mail,
then show what letters can not be delivered.
`
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage for : %v [bundle...]\n", path.Base(os.Args[0]))
		fmt.Fprintf(os.Stderr, content, path.Base(os.Args[0]))
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		flag.Usage()
		return
	}
	for _, bundle := range flag.Args() {
		mailroom := solution.NewMailroom()
		if _, err := mailroom.ProcessLetter(bundle); err != nil {
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
			fmt.Printf("Can not post: %+v\n", post)
		}
	}
}
