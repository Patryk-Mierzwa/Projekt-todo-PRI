package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	ID   primitive.ObjectID `bson:"_id" json:"id"` // `bson:"_id, omitempty"`
	Name string             `json:"name"`
	Done bool               `json:"done"`
}
