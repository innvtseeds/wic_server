package service

import (
	"errors"

	"github.com/innvtseeds/wdic-server/internal/config"
	userRepoDTO "github.com/innvtseeds/wdic-server/internal/dto/repository/user"
	dto "github.com/innvtseeds/wdic-server/internal/dto/service/user"
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

func GetUser(userId primitive.ObjectID, email string) (*model.User, error) {

	if userId == primitive.NilObjectID && email == "" {
		return nil, errors.New("invalid args :: need either email or user id")
	}

	getUserDTO := userRepoDTO.GetUserArgsStruct{
		Email: email,
		Id:    userId,
	}

	searchedUser, err := repository.NewUserRepository(&config.DbConnection).Get(&getUserDTO)

	if err != nil {
		return nil, err
	}

	return searchedUser, nil
}
