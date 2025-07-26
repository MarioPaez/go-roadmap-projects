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
		if err := createFile(); err != nil {
			util.LogError(fmt.Errorf("couldn't create file: %w", err))
		}
	}
}

func createFile() error {
	file, err := os.Create(FILE_NAME)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func openFile() (*os.File, error) {
	file, err := os.OpenFile(FILE_NAME, os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func GetTasks() []model.Task {
	file, err := openFile()
	if err != nil {
		util.LogError(fmt.Errorf("impossible to open the file: %w", err))
		return nil
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil || len(data) == 0 {
		return nil
	}

	var tasks []model.Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		util.LogError(fmt.Errorf("failed to unmarshal json file into tasks: %w", err))
		return nil
	}
	return tasks
}

func GetTasksFiltered(filter string) []model.Task {
	tasks := GetTasks()
	var filtered []model.Task
	for _, task := range tasks {
		if task.Status == model.TaskStatus(filter) {
			filtered = append(filtered, task)
		}
	}
	return filtered
}

func AddTasks(tasks []model.Task) error {
	file, err := openFile()
	if err != nil {
		return err
	}
	defer file.Close()

	if err := file.Truncate(0); err != nil {
		return err
	}
	if _, err := file.Seek(0, 0); err != nil {
		return err
	}
	tasksMarshaled, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	if _, err := file.Write(tasksMarshaled); err != nil {
		return err
	}
	return nil
}

func UpdateTask(newDescription string, id int) error {
	tasks := GetTasks()
	updated := false

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = newDescription
			updated = true
			break
		}
	}
	if !updated {
		return errors.New("id not found")
	}
	return AddTasks(tasks)
}

func DeleteTask(id int) error {
	tasks := GetTasks()
	var updated []model.Task
	found := false
	for _, task := range tasks {
		if task.ID == id {
			found = true
			continue
		}
		updated = append(updated, task)
	}
	if !found {
		return errors.New("id not found")
	}
	return AddTasks(updated)
}

func ChangeStatus(status model.TaskStatus, id int) error {
	tasks := GetTasks()
	updated := false

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = status
			updated = true
			break
		}
	}
	if !updated {
		return errors.New("id not found")
	}
	return AddTasks(tasks)
}
