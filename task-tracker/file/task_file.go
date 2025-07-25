package manager

import (
	"fmt"
	"os"
)

var FILE_NAME = "tasks.json"

func CheckFile() {
	if _, err := os.Stat(FILE_NAME); os.IsNotExist(err) {
		createFile()
	}
}

func createFile() {
	file, err := os.Create(FILE_NAME)
	if err != nil {
		panic(fmt.Errorf("couldn't create file with name %s: %w", FILE_NAME, err))
	}
	defer file.Close()
}
