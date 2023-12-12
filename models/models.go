package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ErrorLog struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Project   primitive.ObjectID `json:"project,omitempty" bson:"project,omitempty"`
	ErrorType string             `json:"error_type,omitempty" bson:"error_type,omitempty"`
	Message   string             `json:"message,omitempty" bson:"message,omitempty"`
	Date      time.Time          `json:"date" bson:"date"`
}

type Project struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name" bson:"name"`
	Logs   []ErrorLog         `json:"logs" bson:"logs"`
	ApiKey string             `json:"api_key" bson:"api_key"`
}
