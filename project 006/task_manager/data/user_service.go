package data

import (
	"context"
	"errors"
	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	connect *Database
}

func NewUserService(connection *Database) *UserService {
	return &UserService{
		connect: connection,
	}
}
func (s *UserService) CUser(u models.User, ctx context.Context) (*models.User, error) {
	u.ID = primitive.NewObjectID()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u.Password = string(hashedPassword)
	if _, err := s.connect.Collections.Users.InsertOne(ctx, u); err != nil {
		return nil, err
	}
	return &u, nil
}
func (s *UserService) AuthenticateUser(username, password string) (*models.User, error) {
    var user models.User
    err := s.connect.Collections.Users.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, errors.New("user not found")
        }
        return nil, err
    }

    // Compare the provided password with the hashed password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return nil, errors.New("invalid password")
    }

    return &user, nil
}
func (s *UserService) GUser(id string, ctx context.Context) (*models.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user models.User
	if err := s.connect.Collections.Users.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
func (s *UserService) GAllUsers(ctx context.Context) ([]models.User, error) {
	cursor, err := s.connect.Collections.Users.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []models.User
	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			continue
		}
		users = append(users, user)
	}
	return users, nil
}
func (s *UserService) UUser(id string, u models.User, ctx context.Context) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	update := bson.M{
		"$set": bson.M{
			"username": u.Username,
			"email":    u.Email,
			"password": u.Password,
			"role":     u.Role,
			"task_ids": u.TaskIDs,
		},
	}

	result, err := s.connect.Collections.Users.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}
func (s *UserService) DUser(id string, ctx context.Context) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := s.connect.Collections.Users.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}
