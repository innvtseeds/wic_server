package api

import (
	"encoding/json"
	"io"
	"net/http"

	authHandlerDTO "github.com/innvtseeds/wdic-server/internal/dto/handler/auth"
	authServiceDTO "github.com/innvtseeds/wdic-server/internal/dto/service/auth"
	userServiceDTO "github.com/innvtseeds/wdic-server/internal/dto/service/user"
	"github.com/innvtseeds/wdic-server/internal/service"
	"github.com/innvtseeds/wdic-server/library/jwt"
	customLogger "github.com/innvtseeds/wdic-server/library/logger"
	apiResponse "github.com/innvtseeds/wdic-server/library/standardization"
)

var myLogger = customLogger.NewLogger()

func Register(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		apiResponse.RequestBodyError(w, err)
	}

	var requestBody authHandlerDTO.Register_RequestBody

	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		apiResponse.UnmarshalError(w, err)
	}

	createUserBody := userServiceDTO.UserCreate_RequestBody{
		Email:    requestBody.Email,
		Password: requestBody.Password,
	}

	myLogger.Info("SERVICE_PAYLOAD :: REGISTER :: ", createUserBody)

	response, err := service.CreateUserService(&createUserBody)
	if err != nil {

		apiResponse.ServiceResponseError(w, err)
	}

	apiResponse.StandardResponse(w, response)
}

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		apiResponse.RequestBodyError(w, err)
		return
	}

	var requestBody authHandlerDTO.Login_RequestBody

	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		apiResponse.UnmarshalError(w, err)
	}

	loginUserBody := authHandlerDTO.Login_RequestBody{
		Email:    requestBody.Email,
		Password: requestBody.Password,
	}

	myLogger.Info("SERVICE_PAYLAOD :: LOGIN :: ", loginUserBody)
	user, err := service.Login((*authServiceDTO.Login_ServiceRequestBody)(&loginUserBody))

	if err != nil {
		apiResponse.ServiceResponseError(w, err)
	}

	apiResponse.StandardResponse(w, user)

}

func Test(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)

	if err != nil {
		apiResponse.RequestBodyError(w, err)
		return
	}

	type TestPayload struct {
		Token string `json:"token"`
	}

	var requestBody TestPayload

	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		apiResponse.UnmarshalError(w, err)
	}

	myLogger.Info("BODY::", requestBody.Token)

	token, err := jwt.DecodeToken(requestBody.Token)

	if err != nil {
		myLogger.Error("HANDLER :: AUTH TEST FAILED :: ", err)
		apiResponse.ServiceResponseError(w, err)
	}

	apiResponse.StandardResponse(w, token)
}
