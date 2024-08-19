// Package: models
// This package contains the models such as task.
package models

// Task is a struct that represents a task object
type Task struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     string    `json:"dueDate"`
	Assignee    string    `json:"assignee"`
	Notes       string    `json:"notes"`
	Status      string    `json:"status"`
}
