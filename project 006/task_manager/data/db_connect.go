package data

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client     *mongo.Client
	DB         *mongo.Database
	Collections struct {
		Users *mongo.Collection
		Tasks *mongo.Collection
	}
}

var (
	dbInstance *Database
)

// InitDB initializes the database connection and ensures indexes
func InitDB() (*Database, error) {
	if dbInstance != nil {
		return dbInstance, nil
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Warning: .env file not found, relying on environment variables")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().
		ApplyURI(os.Getenv("MONGODB_URI")).
		SetMaxPoolSize(100).                // Maximum number of connections
		SetMinPoolSize(10).                 // Minimum number of connections
		SetMaxConnIdleTime(5 * time.Minute) // Maximum idle time for a connection

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Verify the connection
	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	db := client.Database(os.Getenv("DB_NAME"))

	// Initialize collections
	dbInstance = &Database{
		Client: client,
		DB:     db,
	}
	dbInstance.Collections.Users = db.Collection(os.Getenv("USERS_COLLECTION"))
	dbInstance.Collections.Tasks = db.Collection(os.Getenv("TASKS_COLLECTION"))

	// Ensure indexes (only for users collection)
	if err := ensureIndexes(ctx, dbInstance.Collections.Users); err != nil {
		return nil, err
	}

	return dbInstance, nil
}

// ensureIndexes creates required indexes for the users collection
func ensureIndexes(ctx context.Context, usersCollection *mongo.Collection) error {
	// Unique index for username
	usernameIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "username", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	//unique index for email
	emailIndex := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	_, err := usersCollection.Indexes().CreateMany(ctx, []mongo.IndexModel{
		usernameIndex,
		emailIndex,
	})
	return err
}

// CloseDB closes the database connection
func (db *Database) CloseDB() error {
	if db.Client == nil {
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return db.Client.Disconnect(ctx)
}