package api

import (
	"encoding/json"
	"io"
	"net/http"

	userHandlerDTO "github.com/innvtseeds/wdic-server/internal/dto/handler/user"
	userServiceDTO "github.com/innvtseeds/wdic-server/internal/dto/service/user"
	"github.com/innvtseeds/wdic-server/internal/service"
	lib "github.com/innvtseeds/wdic-server/library/logger"
	apiResponse "github.com/innvtseeds/wdic-server/library/standardization"
)

var myLogger = lib.NewLogger()

func Register(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		apiResponse.RequestBodyError(w, err)
	}

	var requestBody userHandlerDTO.CreateUser_RequestBody

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

// func Login(w http.ResponseWriter, r *http.Request) {
// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, "Error in ready request body", http.StatusBadRequest)
// 		return
// 	}
// }
