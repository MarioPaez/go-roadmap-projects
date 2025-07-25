package model

type TaskStatus string

var (
	TASK_STATUS_TODO        TaskStatus = "TO DO"
	TASK_STATUS_IN_PROGRESS TaskStatus = "IN PROGRESS"
	TASK_STATUS_DONE        TaskStatus = "DONE"
)

type Task struct {
	ID          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   string     `json:"createdAt"`
	UpdatedAt   string     `json:"updatedAt"`
}
