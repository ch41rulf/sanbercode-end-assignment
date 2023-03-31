package structs

import "time"

type Assignment struct {
	AssignmentID int64     `json:"assignment_id"`
	MainTaskID   int64     `json:"main_task_id"`
	StaffID      int64     `json:"staff_id"`
	DueDate      time.Time `json:"due_date"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
