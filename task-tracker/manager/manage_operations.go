package manager

import (
	"errors"
	"fmt"
	"task-tracker/service"
)

const ADD = "add"
const UPDATE = "update"
const DELETE = "delete"
const MARK_IN_PROGRESS = "mark-in-progress"
const MARK_DONE = "mark-done"
const LIST = "list"

func ManageOperations(args []string) error {
	taskService := service.NewTaskService()
	if len(args) == 2 {
		switch args[0] {
		case ADD:
			return manageAdd(args, taskService)
		case UPDATE:
			fmt.Println("has puesto update")
		case DELETE:
			fmt.Println("has puesto delete")
		case MARK_IN_PROGRESS:
			fmt.Println("has puesto MARK_IN_PROGRESS")
		case MARK_DONE:
			fmt.Println("has puesto MARK_DONE")
		case LIST:
			fmt.Println("has puesto LIST. Me falta solo enviar list")
		default:
			fmt.Println("Don't support the operation ", args[1], ". Only support add, update, delete, mark-in-progress, mark-done and list")
		}
	}
	return nil
}

func manageAdd(args []string, taskService service.TaskService) error {
	if len(args) == 1 {
		return errors.New("description not provided")
	}
	id := taskService.Add(args[1])
	fmt.Printf("the task with description %s has been successfully saved. The id is :%d\n", args[1], id)
	return nil
}
