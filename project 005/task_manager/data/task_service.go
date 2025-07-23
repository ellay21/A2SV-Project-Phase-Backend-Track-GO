package data

import (
	"context"
	"errors"
	"log"
	"os"
	"task_manager/models"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TaskService struct {
	client     *mongo.Client
	collection *mongo.Collection
}

// use mongodb cloud or local mongodb on you env otherwise it will not work
func NewTaskService() (*TaskService, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		return nil, err
	}

	collection := client.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("TASKS_COLLECTION"))
	return &TaskService{client: client, collection: collection}, nil
}

func (s *TaskService) CTask(t models.Task, ctx context.Context) (*models.Task, error) {
	if err := t.Validate(); err != nil {
		return nil, err
	}
	t.ID = primitive.NewObjectID() // Generate a new ObjectID for the task
	if _, err := s.collection.InsertOne(ctx, t); err != nil {
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
	if err := s.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&task); err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *TaskService) GAllTasks(ctx context.Context) []models.Task {
	cursor, err := s.collection.Find(ctx, bson.M{})
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

	result, err := s.collection.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}

func (s *TaskService) DTask(id string, ctx context.Context) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := s.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}
