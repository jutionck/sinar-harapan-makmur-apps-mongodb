package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Brand struct {
	Id   primitive.ObjectID `bson:"_id"`
	Name string
}
