package data

import (
	"context"
	"errors"
	"task_manager/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskService struct {
	connect *Database
}

func NewTaskService(connection *Database) *TaskService {
	return &TaskService{
		connect: connection,
	}
}

func (s *TaskService) CTask(t models.Task, ctx *gin.Context) (*models.Task, error) {
	if err := t.Validate(); err != nil {
		return nil, err
	}
	t.ID = primitive.NewObjectID()
	if _, err := s.connect.Collections.Tasks.InsertOne(ctx, t); err != nil {
		return nil, err
	}
	// add the task id to the user's task_ids array
	username, exists := ctx.Get("username")
	if !exists {
		return nil, errors.New("username not found in context")
	}
	userFilter := bson.M{"username": username}
	update := bson.M{
		"$addToSet": bson.M{"task_ids": t.ID},
	}
	if _, err := s.connect.Collections.Users.UpdateOne(ctx, userFilter, update); err != nil {
		return nil, err
	}
	return &t, nil
}

func (s *TaskService) GTask(id string, ctx context.Context) (*models.Task, error) {
	// Convert the string ID to a primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var task models.Task
	if err := s.connect.Collections.Tasks.FindOne(ctx, bson.M{"_id": objectID}).Decode(&task); err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *TaskService) GAllTasks(ctx context.Context) []models.Task {
	cursor, err := s.connect.Collections.Tasks.Find(ctx, bson.M{})
	if err != nil {
		return nil
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	for cursor.Next(ctx) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			continue
		}
		tasks = append(tasks, task)
	}
	return tasks
}
func (s *TaskService) GUAllTasks(ctx *gin.Context) ([]models.Task, error) {
	username, exists := ctx.Get("username")
	if !exists {
		return nil, errors.New("username not found in context")
	}

	// First, find the user to get their task IDs
	var user models.User
	userFilter := bson.M{"username": username}
	err := s.connect.Collections.Users.FindOne(ctx, userFilter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	if len(user.TaskIDs) == 0 {
		return []models.Task{}, nil
	}

	// Find all tasks where ID is in the user's TaskIDs array
	taskFilter := bson.M{"_id": bson.M{"$in": user.TaskIDs}}
	cursor, err := s.connect.Collections.Tasks.Find(ctx, taskFilter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []models.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *TaskService) UTask(t models.Task, id string, ctx context.Context) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	if err := t.Validate(); err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"title":       t.Title,
			"description": t.Description,
			"dueDate":     t.DueDate,
			"status":      t.Status,
		},
	}

	result, err := s.connect.Collections.Tasks.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}

func (s *TaskService) DTask(id string, ctx *gin.Context) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := s.connect.Collections.Tasks.DeleteOne(ctx, bson.M{"_id": objectID})

	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}
	// delete the task id from the user's task_ids array
	username, exists := ctx.Get("username")
	if !exists {
		return errors.New("username not found in context")
	}

	userFilter := bson.M{"username": username}
	update := bson.M{
		"$pull": bson.M{"task_ids": objectID},
	}

	_, err = s.connect.Collections.Users.UpdateOne(ctx, userFilter, update)
	if err != nil {
		return err
	}

	return nil
}
