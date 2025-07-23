package models

import (
	"time"
	"errors" 
)

// Task represents a task in the system.
type Task struct {
	ID          int
	Title       string
	Description string
	DueDate     time.Time
	Status      string
}

func (t *Task) Validate() error {
	if t.Title == "" {return errors.New("the title is empty bozo")}
	if len(t.Title) > 100 {return errors.New("the title is too long bozo")}
	if t.Description == "" {return errors.New("the description is empty bozo")}
	if t.DueDate.IsZero() {return errors.New("the due date is empty bozo")}
	if t.Status == "" || (t.Status != "pending" && t.Status != "in progress" && t.Status != "completed") {return errors.New("invalid status bozo")}
	return nil
}

