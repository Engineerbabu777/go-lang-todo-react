package models;


import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)


// TO CREATE MODELS IN GO LANG WE USE STRUCTS!
type TodoList struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"` 
	Task string `json:"task,omitempty"`
	Status string `json:"status,omitempty"`
}