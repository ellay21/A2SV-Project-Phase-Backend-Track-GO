package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Username string               `bson:"username" json:"username" minLength:"3" maxLength:"50" binding:"required"`
	Email    string               `bson:"email" json:"email" maxLength:"100" binding:"required"`
	Password string               `bson:"password" json:"password" minLength:"6" maxLength:"100" binding:"required"`
	Role     string               `bson:"role" json:"role" enum:"admin,user"`
	TaskIDs  []primitive.ObjectID `bson:"task_ids,omitempty" json:"task_ids"` // List of task IDs associated with the user
}
