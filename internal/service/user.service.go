package service

import (
	"errors"

	"github.com/innvtseeds/wdic-server/internal/config"
	dto "github.com/innvtseeds/wdic-server/internal/dto/handler/user"
	userRepoDTO "github.com/innvtseeds/wdic-server/internal/dto/repository/user"
	"github.com/innvtseeds/wdic-server/internal/model"
	"github.com/innvtseeds/wdic-server/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserService(user *dto.CreateUserPayload) (*model.User, error) {

	if user.Email == "" || user.Password == "" {
		return nil, errors.New("Parameters Missing")
	}

	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("Password Hashing Failed")
	}

	var createUserModelDTO model.User = model.User{
		Email:    user.Email,
		Password: string(hashedPassword),
	}

	createdUser, err := repository.NewUserRepository(&config.DbConnection).Create(&createUserModelDTO)

	if err != nil {
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
