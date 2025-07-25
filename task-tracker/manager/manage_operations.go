package manager

import (
	"errors"
	"fmt"
	"task-tracker/service"
	"task-tracker/util"
)

var ADD = "add"
var UPDATE = "update"
var DELETE = "delete"
var MARK_IN_PROGRESS = "mark-in-progress"
var MARK_DONE = "mark-done"
var LIST = "list"

func ManageOperations(args []string) {
	taskService := service.NewTaskService()
	if len(args) == 2 {
		switch args[0] {
		case ADD:
			manageAdd(args, taskService)
		case UPDATE:
			println("has puesto update")
		case DELETE:
			println("has puesto delete")
		case MARK_IN_PROGRESS:
			println("has puesto MARK_IN_PROGRESS")
		case MARK_DONE:
			println("has puesto MARK_DONE")
		case LIST:
			println("has puesto LIST. Me falta solo enviar list")
		default:
			fmt.Println("Don't support the operation ", args[1], ". Only support add, update, delete, mark-in-progress, mark-done and list")
		}
	}
}

func manageAdd(args []string, taskService service.TaskService) int {
	if len(args) == 1 {
		util.LogError(errors.New("description not provided"))
	}
	id := taskService.Add(args[1])
	fmt.Printf("the task with description %s has been successfully saved. The id is :%d\n", args[1], id)
	return id
}
