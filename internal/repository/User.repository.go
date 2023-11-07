package repository

import (
	"context"
	"errors"
	"log"

	"github.com/innvtseeds/wdic-server/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	collection := db.Collection("user")
	return &UserRepository{
		collection: collection,
	}
}

func (r *UserRepository) Create(user *model.User) (*model.User, error) {
	result, err := r.collection.InsertOne(context.Background(), user)
	if err != nil {
		//* @sreerag: NEED TO CHECK IF LOG.FATAL EXITS THE PROGRAM
		//* IF SO THIS NEEDS TO BE AVOIDED AND INSTEAD JUST HAVE LOGGING BE DONE
		log.Fatal("ERROR IN USER INSERT", err)
		return nil, errors.New("Insert to Collection failed")
	}

	createdUser := &model.User{
		Id:    result.InsertedID.(primitive.ObjectID),
		Email: user.Email,
	}

	return createdUser, nil
}
