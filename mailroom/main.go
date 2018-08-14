package main

import (
  "fmt"
  "github.com/SeedJobs/devops-go-problems/mailroom/utils"
  "github.com/SeedJobs/devops-go-problems/mailroom/types"
  "github.com/SeedJobs/devops-go-problems/mailroom/solution"
)

func main() {
    contents, err := utils.PostageReader("letters/bundle-01")
    if err != nil {
      fmt.Println("Unable to read post due to:", err)
      return
    }
    mailroom := solution.NewMailroom()
    for _, content := range contents {
      mailroom.ProcessLetter(content)
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
