package main

import (
	"github-activity/commands"
	"log"
	"os"
)

func main() {
	args := os.Args
	if err := commands.ManageCommands(args); err != nil {
		log.Fatal(err.Error())
	}
}
