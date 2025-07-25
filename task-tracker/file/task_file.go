package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"task-tracker/model"
	"task-tracker/util"
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

func openFile() *os.File {
	file, err := os.OpenFile(FILE_NAME, os.O_RDWR, 0666)
	if err != nil {
		util.LogError(errors.New("impossible to open the file to read the last ID"))
	}
	return file
}

func GetTasks() []model.Task {
	var tasks []model.Task
	file := openFile()
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil || len(data) == 0 {
		return nil
	}

	if err := json.Unmarshal(data, &tasks); err != nil {
		util.LogError(errors.New("failed to unmarshal json file into tasks"))
	}

	return tasks
}

func AddTasks(tasks []model.Task) {
	file := openFile()
	defer file.Close()

	tasksMarshaled, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		util.LogError(errors.New("couldn't marshal the tasks"))
	}

	if _, err := file.Write(tasksMarshaled); err != nil {
		util.LogError(errors.New("error while trying write in the file"))
	}

}
