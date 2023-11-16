package repository

import (
	"context"
	"errors"
	"log"

	userRepoDTO "github.com/innvtseeds/wdic-server/internal/dto/repository/user"
	sharedDTO "github.com/innvtseeds/wdic-server/internal/dto/shared"
	"github.com/innvtseeds/wdic-server/internal/model"
	"github.com/innvtseeds/wdic-server/library/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var myLogger = logger.NewLogger()

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
	myLogger.Info("USER REPOSITORY :: INSERT BODY ::", user)
	result, err := r.collection.InsertOne(context.Background(), user)
	if err != nil {

		myLogger.Error("ERROR IN USER INSERT", err)
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
		myLogger.Error("User update", err)
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

// Get One
func (r *UserRepository) Get(identifier *userRepoDTO.GetUserArgsStruct) (*model.User, error) {
	if identifier.Email == "" && identifier.Id == primitive.NilObjectID {
		return nil, errors.New("missing params :: need either id or email")
	}

	var where bson.M

	if identifier.Id != primitive.NilObjectID {
		where = bson.M{"_id": identifier.Id}
	} else if identifier.Email != "" {
		where = bson.M{"email": identifier.Email}
	}

	var user model.User

	err := r.collection.FindOne(context.Background(), where).Decode(&user)

	if err != nil {
		return nil, errors.New("failed to get User")
	}

	return &user, nil

}

// Get ALl
func (r *UserRepository) GetAll(paginationValues *sharedDTO.PaginationStruct) ([]*model.User, error) {

	findOptions := options.Find()
	findOptions.SetSkip((paginationValues.Page - 1) * paginationValues.PageSize)
	findOptions.SetLimit(paginationValues.PageSize)

	var where bson.M

	if paginationValues.Search != "" {
		where = bson.M{"email": paginationValues.Search}
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
			myLogger.Error("ERROR IN DECODING USER", err)
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
