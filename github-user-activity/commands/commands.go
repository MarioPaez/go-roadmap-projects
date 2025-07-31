package commands

import (
	"fmt"
	"github-activity/service"
	"github-activity/utils"
)

const (
	GITHUB_ACTIVITY = "github-activity"
	HELP            = "help"
)

func ManageCommands(args []string) error {

	if len(args) != 3 {
		return utils.CommandUnknown(GITHUB_ACTIVITY, HELP)
	}

	cmd := args[1]
	username := args[2]

	switch cmd {
	case GITHUB_ACTIVITY:
		return service.GetGithubActivity(username)
	case HELP:
		fmt.Println("Use 'github-activity <username>'")
	default:
		return utils.CommandUnknown(GITHUB_ACTIVITY, HELP)
	}

	return nil
}
