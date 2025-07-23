package data

import (
	"errors"
	"task_manager/models"
	"time"
)

type TaskService struct {
	tasks map[int]models.Task
}

func NewTaskService() *TaskService {
	return &TaskService{
		tasks: make(map[int]models.Task),
	}
}
func (s *TaskService) Initialize() {
	s.tasks[1] = models.Task{ID: 1, Title: "Sample Task", Description: "This is a sample task", DueDate: time.Now().Add(24 * time.Hour), Status: "pending"}
	s.tasks[2] = models.Task{ID: 2, Title: "Another Task", Description: "This is another task", DueDate: time.Now().Add(48 * time.Hour), Status: "in progress"}
	s.tasks[3] = models.Task{ID: 3, Title: "Completed Task", Description: "This task is completed", DueDate: time.Now().Add(-24 * time.Hour), Status: "completed"}
	s.tasks[4] = models.Task{ID: 4, Title: "Future Task", Description: "This task is scheduled for the future", DueDate: time.Now().Add(72 * time.Hour), Status: "pending"}
	s.tasks[5] = models.Task{ID: 5, Title: "Overdue Task", Description: "This task is overdue", DueDate: time.Now().Add(-48 * time.Hour), Status: "in progress"}
}

func (s *TaskService) CTask(t models.Task) error {
	if err := t.Validate(); err != nil {
		return err
	}
	t.ID = len(s.tasks) + 1
	s.tasks[t.ID] = t
	return nil
}

func (s *TaskService) GTask(id int) (*models.Task, error) {
	if t, exists := s.tasks[id]; exists{
		return &t, nil
	}
	return nil, errors.New("task not found")
}

func (s *TaskService) GAllTasks() []models.Task {
	tasks := make([]models.Task, 0, len(s.tasks))
	for _, t := range s.tasks {
		tasks = append(tasks, t)
	}
	return tasks
}

func (s *TaskService) UTask(t models.Task) error {
	if err := t.Validate(); err != nil {
		return err
	}
	if _, exists := s.tasks[t.ID]; !exists {
		return errors.New("task not found")
	}
	s.tasks[t.ID] = t
	return nil
}

func (s *TaskService) DTask(id int) error {
	if _, exists := s.tasks[id]; exists {
		delete(s.tasks, id)
		return nil
	}
	return errors.New("task not found")
}
