package service

import (
	"github.com/innvtseeds/wdic-server/internal/config"
	userRepoDTO "github.com/innvtseeds/wdic-server/internal/dto/repository/user"
	loginServiceDto "github.com/innvtseeds/wdic-server/internal/dto/service/auth"
	"github.com/innvtseeds/wdic-server/internal/repository"
	customLogger "github.com/innvtseeds/wdic-server/library/logger"
	"golang.org/x/crypto/bcrypt"
)

var myLogger = customLogger.NewLogger()

func Login(payload *loginServiceDto.Login_ServiceRequestBody) (*string, error) {

	// Verify the payload
	if payload.Email == "" || payload.Password == "" {
		myLogger.Error("SERVICE :: PAYLOAD VARIABLES MISSING :: ", payload)
	}

	// get the user based on email
	userIdentifier := userRepoDTO.GetUserArgsStruct{
		Email: payload.Email,
	}
	myLogger.Info("SERVICE :: USER IDENTIFIER", userIdentifier)
	user, err := repository.NewUserRepository(&config.DbConnection).Get(&userIdentifier)
	if err != nil {
		myLogger.Error("SERVICE :: ENTITY NOT FOUND :: ", err)
		return nil, err
	}

	//verify the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))

	if err != nil {
		myLogger.Error("SERVICE :: PASSWORD INCORRECT :: ", err)
		return nil, err
	}

	// return the user is password verified
	return &user.Email, nil

}
