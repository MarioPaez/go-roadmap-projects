package main

import (
	"fmt"
	"os"
	f "task-tracker/file"
	m "task-tracker/manager"
)

func main() {
	fmt.Println("Welcome to your Task Tracker")
	f.CheckFile()
	m.ManageOperations(os.Args[1:])
}
