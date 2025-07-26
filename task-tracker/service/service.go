package service

import (
	"task-tracker/file"
	"task-tracker/model"
	"time"
)

const DATE_FORMAT = "02-01-2006 15:04:05"

type TaskService interface {
	Add(description string) int
	Update(newDescription string, id int) error
	Delete(id int) error
	ListAll() []model.Task
	ListAllWithFilters(filter string) []model.Task
	ChangeStatus(status model.TaskStatus, id int) error
}

type taskService struct {
}

func NewTaskService() TaskService {
	return &taskService{} //More efficient because we avoid to copy the structure
}

func (t *taskService) Add(description string) int {
	tasks := file.GetTasks()
	id := len(tasks)
	task := model.Task{
		ID:          id,
		Description: description,
		Status:      model.TASK_STATUS_TODO,
		CreatedAt:   time.Now().Format(DATE_FORMAT),
		UpdatedAt:   time.Now().Format(DATE_FORMAT),
	}
	tasks = append(tasks, task)
	file.AddTasks(tasks)
	return id
}

func (t *taskService) Update(newDescription string, id int) error {
	return file.UpdateTask(newDescription, id)
}

func (t *taskService) Delete(id int) error {
	return file.DeleteTask(id)
}

func (t *taskService) ListAll() []model.Task {
	return file.GetTasks()
}

func (t *taskService) ListAllWithFilters(filter string) []model.Task {
	return file.GetTasksFiltered(filter)
}

func (t *taskService) ChangeStatus(status model.TaskStatus, id int) error {
	return file.ChangeStatus(status, id)
}
