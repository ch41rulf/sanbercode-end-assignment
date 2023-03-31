package structs

import (
	"database/sql"
	"time"
)

type Subtask struct {
	SubtaskID   int64          `json:"subtask_id"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	Status      string         `json:"status"`
	MainTaskID  sql.NullInt64  `json:"main_task_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}
