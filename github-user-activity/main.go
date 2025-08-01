package main

import (
	"github-activity/commands"
	"github-activity/service"
	"log"
	"os"
)

func main() {
	args := os.Args
	githubService := service.NewGithubService()
	if err := commands.ManageCommands(args, githubService); err != nil {
		log.Fatal(err.Error())
	}
}
