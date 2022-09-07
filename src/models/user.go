package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	Username  string             `json:"username"`
	Password  string             `json:"password"`
	IsAdmin   bool               `json:"isAdmin"`
	IsDeleted bool               `json:"isDeleted"`
}
