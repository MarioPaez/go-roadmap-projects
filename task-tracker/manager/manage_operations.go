package manager

import (
	"errors"
	"fmt"
	"strconv"
	"task-tracker/model"
	"task-tracker/service"
	"task-tracker/util"
)

const ADD = "add"
const UPDATE = "update"
const DELETE = "delete"
const MARK_IN_PROGRESS = "mark-in-progress"
const MARK_DONE = "mark-done"
const LIST = "list"

func ManageOperations(args []string) error {
	taskService := service.NewTaskService()
	if len(args) > 0 {
		switch args[0] {
		case ADD:
			return manageAdd(args, taskService)
		case UPDATE:
			return manageUpdate(args, taskService)
		case DELETE:
			return manageDelete(args, taskService)
		case MARK_IN_PROGRESS:
			return manageChangeStatus(args, model.TASK_STATUS_IN_PROGRESS, taskService)
		case MARK_DONE:
			return manageChangeStatus(args, model.TASK_STATUS_DONE, taskService)
		case LIST:
			return manageList(args, taskService)
		default:
			fmt.Println("Don't support the operation <", args[0], ">. Only support add, update, delete, mark-in-progress, mark-done and list.")
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

func manageUpdate(args []string, taskService service.TaskService) error {
	if len(args) != 3 {
		return errors.New("update operation must to have {update {id} 'new description'}")
	}

	id, err := strconv.Atoi(args[1])
	if err != nil {
		return errors.New("the id must to be numeric")
	}
	newDescription := args[2]
	if err := taskService.Update(newDescription, id); err != nil {
		return errors.New("does not exist task with ID:" + fmt.Sprint(id))
	}

	fmt.Printf("the task with ID %d has been updated. The new description is %s.\n", id, newDescription)
	return nil
}

func manageDelete(args []string, taskService service.TaskService) error {
	if len(args) != 2 {
		return errors.New("delete operation must to have {delete {id}}")
	}
	id, err := strconv.Atoi(args[1])
	if err != nil {
		return errors.New("the id must to be numeric")
	}
	if err := taskService.Delete(id); err != nil {
		return errors.New("does not exist task with ID:" + fmt.Sprint(id))
	}
	fmt.Printf("the task with ID %d has been deleted.\n", id)
	return nil
}

func manageChangeStatus(args []string, status model.TaskStatus, taskService service.TaskService) error {
	if len(args) == 1 {
		return errors.New("id not provided")
	}
	id, err := strconv.Atoi(args[1])
	if err != nil {
		return errors.New("the id must to be numeric")
	}
	if err := taskService.ChangeStatus(status, id); err != nil {
		return errors.New("does not exist task with ID:" + fmt.Sprint(id))
	}
	fmt.Printf("the task with ID %d has been updated to state %s.\n", id, status)
	return nil
}

func manageList(args []string, taskService service.TaskService) error {
	if len(args) > 2 {
		return errors.New("we only support 'list', 'list done', 'list todo', 'list in-progress'")
	}
	var tasks []model.Task
	if len(args) == 1 { //List all tasks
		tasks = taskService.ListAll()

	} else {
		filter := args[1]
		if !util.CheckValidStatus(filter) {
			return errors.New("we only support 'done', 'todo', 'in-progress'")
		}
		tasks = taskService.ListAllWithFilters(filter)

	}
	util.PrintTasksTable(tasks)
	return nil
}
