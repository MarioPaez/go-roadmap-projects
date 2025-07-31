package utils

import (
	"fmt"
	"github-activity/model"
)

func CommandUnknown(cmd1, cmd2 string) error {
	return fmt.Errorf("command unknown. please see '%s %s' for more information", cmd1, cmd2)
}

func PrettyPrint(events []model.Event) {
	for _, event := range events {
		switch event.Type {
		case "PushEvent":
			fmt.Printf("- Pushed %d commit(s) to %s\n", event.Payload.Size, event.Repo.Name)
		case "IssuesEvent":
			if event.Payload.Action == "opened" {
				fmt.Printf("- Opened a new issue in %s\n", event.Repo.Name)
			}
		case "WatchEvent":
			fmt.Printf("- Starred %s\n", event.Repo.Name)
		}
	}
}
