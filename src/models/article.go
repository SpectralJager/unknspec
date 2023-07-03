package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Title     string             `json:"title" bson:"title"`
	Abstract  string             `json:"abstract" bson:"abstract"`
	Body      string             `json:"body" bson:"body"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"upadted_at" bson:"updated_at"`
	IsDraft   bool               `json:"is_draft" bson:"is_draft"`
	Author    string             `json:"author" bson:"author"`
}
