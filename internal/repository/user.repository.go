package repository

import (
	"context"
	"errors"
	"log"

	"github.com/innvtseeds/wdic-server/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// Create Logic
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

// Update Logic
func (r *UserRepository) Update(userId *primitive.ObjectID, user *model.User) (*string, error) {
	if userId == nil {
		return nil, errors.New("missing param :: need user Id")
	}

	update := bson.D{
		{
			Key: "$set",
			Value: bson.D{
				{Key: "email", Value: user.Email},
				{Key: "password", Value: user.Password},
			},
		},
	}

	_, err := r.collection.UpdateByID(context.Background(), userId, update)
	if err != nil {
		log.Fatal("User update", err)
		return nil, errors.New("Update Query Failed")
	}

	successMessage := "Successfully Updated"
	return &successMessage, nil

}

// Delete Logic
func (r *UserRepository) Delete(userId *primitive.ObjectID) (*string, error) {
	if userId == nil {
		return nil, errors.New("missing params :: UserId Missing")
	}

	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": userId})

	if err != nil {
		return nil, errors.New("Delete User Failed")
	}

	successMessage := "User Deleted Successfully"
	return &successMessage, nil

}

type GetUserArgsStruct struct {
	id    primitive.ObjectID
	email string
}

// Get One
func (r *UserRepository) Get(identifier *GetUserArgsStruct) (*model.User, error) {
	if identifier.email == "" && identifier.id == primitive.NilObjectID {
		return nil, errors.New("missing params :: need either id or email")
	}

	var where bson.M

	if identifier.id != primitive.NilObjectID {
		where = bson.M{"_id": identifier.id}
	} else if identifier.email != "" {
		where = bson.M{"email": identifier.email}
	}

	var user model.User

	err := r.collection.FindOne(context.Background(), where).Decode(&user)

	if err != nil {
		return nil, errors.New("failed to get User")
	}

	return &user, nil

}

type PaginationStruct struct {
	page     int64
	pageSize int64
	search   string
}

// Get ALl
func (r *UserRepository) GetAll(paginationValues *PaginationStruct) ([]*model.User, error) {

	findOptions := options.Find()
	findOptions.SetSkip((paginationValues.page - 1) * paginationValues.pageSize)
	findOptions.SetLimit(paginationValues.pageSize)

	var where bson.M

	if paginationValues.search != "" {
		where = bson.M{"email": paginationValues.search}
	}

	cursor, err := r.collection.Find(context.Background(), where, findOptions)
	if err != nil {
		log.Println("ERROR IN GET ALL USERS", err)
		return nil, err
	}

	defer cursor.Close(context.Background())

	users := []*model.User{}
	for cursor.Next(context.Background()) {
		var user model.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Println("ERROR IN DECODING USER", err)
			continue
		}
		users = append(users, &user)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal("Get All User Failed", err)
		return nil, errors.New("Get All User Failed")
	}

	return users, nil
}