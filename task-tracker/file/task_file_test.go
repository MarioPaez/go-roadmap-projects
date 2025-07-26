package file

import (
	"encoding/json"
	"os"
	"task-tracker/model"
	"testing"
)

func setupTestFile(tasks []model.Task) (func(), error) {
	tmpfile, err := os.CreateTemp("", "tasks_test_*.json")
	if err != nil {
		return nil, err
	}
	// Sobrescribe FILE_NAME para los tests
	FILE_NAME = tmpfile.Name()
	data, _ := json.Marshal(tasks)
	tmpfile.Write(data)
	tmpfile.Close()
	return func() { os.Remove(FILE_NAME) }, nil
}

func TestGetTasksFiltered(t *testing.T) {
	tasks := []model.Task{
		{ID: 1, Description: "A", Status: model.TASK_STATUS_TODO},
		{ID: 2, Description: "B", Status: model.TASK_STATUS_DONE},
	}
	cleanup, err := setupTestFile(tasks)
	if err != nil {
		t.Fatalf("setup failed: %v", err)
	}
	defer cleanup()

	filtered := GetTasksFiltered(string(model.TASK_STATUS_DONE))
	if len(filtered) != 1 || filtered[0].ID != 2 {
		t.Errorf("expected 1 done task with ID 2, got %+v", filtered)
	}
}
