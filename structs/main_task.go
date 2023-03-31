package structs

import (
	"database/sql"
	"time"
)

type MainTask struct {
	TaskId      int64          `json:"taskId"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	DueDate     time.Time      `json:"dueDate"`
	Status      string         `json:"status"`
	ManagerID   sql.NullInt64  `json:"managerID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}
