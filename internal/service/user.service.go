package service

import (
	"errors"

	"github.com/innvtseeds/wdic-server/internal/config"
	userRepoDTO "github.com/innvtseeds/wdic-server/internal/dto/repository/user"
	dto "github.com/innvtseeds/wdic-server/internal/dto/service/user"
	sharedDto "github.com/innvtseeds/wdic-server/internal/dto/shared"

	"github.com/innvtseeds/wdic-server/internal/model"
	"github.com/innvtseeds/wdic-server/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserService(user *dto.UserCreate_RequestBody) (*model.User, error) {

	if user.Email == "" || user.Password == "" {
		myLogger.Error("SERVICE_ERROR :: ", "PARAMETER_MISSING :: ", user.Email, user.Password)
		return nil, errors.New("parameters missing")
	}

	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		myLogger.Error("SERVICE_ERROR :: Password Hashing Failed")
		return nil, errors.New("password hashing failed")
	}

	var createUserModelDTO model.User = model.User{
		Email:    user.Email,
		Password: string(hashedPassword),
	}

	createdUser, err := repository.NewUserRepository(&config.DbConnection).Create(&createUserModelDTO)

	if err != nil {
		myLogger.Error("SERVICE_ERROR :: ", "REPO_RESPONSE_ERROR :::", err)
		return nil, err
	}

	return createdUser, nil

}

func GetUser(options dto.GetUserOptions) (*model.User, error) {

	if options.UserID == primitive.NilObjectID && options.Email == "" {
		return nil, errors.New("invalid args :: need either email or user id")
	}

	getUserDTO := userRepoDTO.GetUserArgsStruct{
		Email: options.Email,
		Id:    options.UserID,
	}

	searchedUser, err := repository.NewUserRepository(&config.DbConnection).Get(&getUserDTO)

	if err != nil {
		return nil, err
	}

	return searchedUser, nil
}

func GetUsers(pagination *sharedDto.PaginationStruct) ([]*model.User, error) {

	users, err := repository.NewUserRepository(&config.DbConnection).GetAll(pagination)

	if err != nil {
		myLogger.Error("ERROR IN USER SERVICE :: GET ALL USERS :: ", err)
		return nil, err
	}

	return users, nil
}

func DeleteUser(userId *primitive.ObjectID) (*string, error) {

	repo := repository.NewUserRepository(&config.DbConnection)

	dto := userRepoDTO.GetUserArgsStruct{
		Id: *userId,
	}
	userVerify, err := repo.Get(&dto)
	if err != nil {
		return nil, err
	}
	if userVerify == nil {
		return nil, errors.New("user not found")
	}

	response, err := repo.Delete(userId)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func UpdateUser(userId *primitive.ObjectID, data *model.User) (*string, error) {

	repo := repository.NewUserRepository(&config.DbConnection)

	dto := userRepoDTO.GetUserArgsStruct{
		Id: *userId,
	}
	userVerify, err := repo.Get(&dto)
	if err != nil {
		return nil, err
	}
	if userVerify == nil {
		return nil, errors.New("user not found")
	}

	response, err := repo.Update(userId, data)

	if err != nil {
		myLogger.Error("ERROR IN UpdateUser in User Service :: repository response", err)
		return nil, err
	}

	return response, nil
}
