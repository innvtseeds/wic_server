package dto

import (
	"github.com/innvtseeds/wdic-server/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetUser_RequestDTO struct {
	UserId primitive.ObjectID
}

type GetUser_ResponseDTO struct {
	User *model.User
}

type CreateUser_RequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
