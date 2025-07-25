package main

import (
	"fmt"
	"log"
	"os"
	f "task-tracker/file"
	m "task-tracker/manager"
)

func main() {
	fmt.Println("Welcome to your Task Tracker")
	f.CheckFile()
	if err := m.ManageOperations(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
