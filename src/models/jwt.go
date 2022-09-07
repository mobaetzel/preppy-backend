package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Jwt struct {
	AccessToken  string             `json:"accessToken"`
	RefreshToken string             `json:"refreshToken"`
	UserId       primitive.ObjectID `json:"userId"`
}
