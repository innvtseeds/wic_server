package service

import (
	"github.com/innvtseeds/wdic-server/internal/config"
	userRepoDTO "github.com/innvtseeds/wdic-server/internal/dto/repository/user"
	loginServiceDto "github.com/innvtseeds/wdic-server/internal/dto/service/auth"
	dto "github.com/innvtseeds/wdic-server/internal/dto/service/user"
	"github.com/innvtseeds/wdic-server/internal/repository"
	"github.com/innvtseeds/wdic-server/library/jwt"
	customLogger "github.com/innvtseeds/wdic-server/library/logger"
	"golang.org/x/crypto/bcrypt"
)

var myLogger = customLogger.NewLogger()

func Register(user *dto.UserCreate_RequestBody) (*string, error) {

	response, err := CreateUserService(user)
	if err != nil {
		return nil, err
	}

	tokenPayload := map[string]interface{}{
		"Id":    response.Id,
		"Email": response.Email,
	}

	token, err := jwt.GenerateToken(tokenPayload)
	if err != nil {
		myLogger.Error("SERVICE :: JWT TOKEN GENERATION FAILED :: ", err)
		return nil, err
	}
	// return the user is password verified
	return token, nil

}

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

	myLogger.Info("SERVICE :: LOGIN :: USER :: ", user)

	tokenPayload := map[string]interface{}{
		"Id":    user.Id,
		"Email": user.Email,
	}

	token, err := jwt.GenerateToken(tokenPayload)
	if err != nil {
		myLogger.Error("SERVICE :: JWT TOKEN GENERATION FAILED :: ", err)
		return nil, err
	}
	// return the user is password verified
	return token, nil

}
