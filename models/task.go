package models

import "time"

type Task struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate"`
	Assignee    string    `json:"assignee"`
	Notes       string    `json:"notes"`
	Status      string    `json:"status"`
}
