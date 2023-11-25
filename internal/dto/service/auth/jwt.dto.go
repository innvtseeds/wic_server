package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type JWTCreationPayload struct {
	Id    primitive.ObjectID `json:"id"`
	Email string             `json:"email"`
}
