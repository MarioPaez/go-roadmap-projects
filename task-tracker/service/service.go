package service

import (
	"task-tracker/file"
	"task-tracker/model"
	"time"
)

type TaskService interface {
	Add(description string) int
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
		CreatedAt:   time.Now().Format("02-01-2006 15:04:05"),
		UpdatedAt:   time.Now().Format("02-01-2006 15:04:05"),
	}
	tasks = append(tasks, task)
	file.AddTasks(tasks)
	return id
}
