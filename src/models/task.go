package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AdminTask struct {
	Id   primitive.ObjectID `bson:"_id"`
	Task string             `bson:"task"`
}
