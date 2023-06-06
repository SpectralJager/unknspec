package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	Title     string             `json:"title" bson:"title"`
	Abstract  string             `json:"abstract" bson:"abstract"`
	Tags      []string           `json:"tags" bson:"tags"`
	Body      string             `json:"body" bson:"body"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpadtedAt time.Time          `json:"upadted_at" bson:"upadted"`
	IsDraft   bool               `json:"is_draft" bson:"is_draft"`
	Author    string             `json:"author" bson:"author"`
}
