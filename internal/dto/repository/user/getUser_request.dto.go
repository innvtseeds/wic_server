package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type GetUserArgsStruct struct {
	Id    primitive.ObjectID
	Email string
}
