package repository

import "go.mongodb.org/mongo-driver/mongo"

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	collection := db.Collection("user")
	return &UserRepository{
		collection: collection,
	}
}
