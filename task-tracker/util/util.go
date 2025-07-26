package util

import (
	"fmt"
	"log"
	"task-tracker/model"
)

var States = map[string]struct{}{
	string(model.TASK_STATUS_DONE):        {},
	string(model.TASK_STATUS_IN_PROGRESS): {},
	string(model.TASK_STATUS_TODO):        {},
}

func LogError(err error) {
	log.Fatal(err)
}

func PrintTasksTable(tasks []model.Task) {
	if len(tasks) == 0 {
		fmt.Println("There are not tasks.")
		return
	}
	fmt.Printf("%-4s | %-40s | %-12s | %-20s | %-20s\n", "ID", "Description", "Status", "Created At", "Updated At")
	fmt.Println("-----+------------------------------------------+--------------+----------------------+----------------------")
	for _, t := range tasks {
		fmt.Printf("%-4d | %-40s | %-12s | %-20s | %-20s\n", t.ID, t.Description, t.Status, t.CreatedAt, t.UpdatedAt)
	}
}

func CheckValidStatus(filter string) bool {
	_, ok := States[filter]
	return ok
}
