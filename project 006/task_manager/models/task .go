package models

import (
	"time"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
 
)

// Task represents a task in the system.
type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	DueDate     time.Time          `bson:"dueDate" json:"dueDate"`
	Status      string             `bson:"status" json:"status"`
}


func (t *Task) Validate() error {
	if t.Title == "" {return errors.New("the title is empty bozo")}
	if len(t.Title) > 100 {return errors.New("the title is too long bozo")}
	if t.Description == "" {return errors.New("the description is empty bozo")}
	if t.DueDate.IsZero() {return errors.New("the due date is empty bozo")}
	if t.Status == "" || (t.Status != "pending" && t.Status != "in progress" && t.Status != "completed") {return errors.New("invalid status bozo")}
	return nil
}

