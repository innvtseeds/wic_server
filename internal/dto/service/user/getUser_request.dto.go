package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserCreate_RequestBody struct {
	Email    string
	Password string
}

type GetUserOptions struct {
	Email  string
	UserID primitive.ObjectID
}
