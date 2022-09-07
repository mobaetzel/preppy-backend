package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Recipe struct {
	Id                 primitive.ObjectID `bson:"_id" json:"id"`
	AuthorId           primitive.ObjectID `json:"authorId"`
	Title              string             `json:"title"`
	Description        string             `json:"description"`
	Instructions       string             `json:"instructions"`
	Servings           uint               `json:"servings"`
	CaloriesPerServing uint               `json:"caloriesPerServing"`
	Created            time.Time          `json:"created"`
	Updated            time.Time          `json:"updated"`
	Ingredients        []Ingredient       `json:"ingredients"`
	Tags               []string           `json:"tags"`
}

type Ingredient struct {
	Name   string  `json:"name"`
	Amount float32 `json:"amount"`
	Unit   string  `json:"unit"`
}
