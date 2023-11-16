package api

import (
	"encoding/json"
	"io"
	"net/http"

	userHandlerDTO "github.com/innvtseeds/wdic-server/internal/dto/handler/user"
	userServiceDTO "github.com/innvtseeds/wdic-server/internal/dto/service/user"
	"github.com/innvtseeds/wdic-server/internal/service"
	"github.com/innvtseeds/wdic-server/library/logger"
)

var myLogger = logger.NewLogger()

func Register(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		myLogger.Error("REQUEST_BODY_ERROR :: ", err)
		http.Error(w, "Error in ready request body", http.StatusBadRequest)
		return
	}

	var requestBody userHandlerDTO.CreateUser_RequestBody

	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		myLogger.Error("UNMARHALING_ERROR :: ", err)
		http.Error(w, "Error unmashalling JSON", http.StatusBadRequest)
		return
	}

	createUserBody := userServiceDTO.UserCreate_RequestBody{
		Email:    requestBody.Email,
		Password: requestBody.Password,
	}

	myLogger.Info("SERVICE_PAYLOAD :: REGISTER :: ", createUserBody)

	response, err := service.CreateUserService(&createUserBody)
	if err != nil {

		myLogger.Error("SERVICE_ERROR :: ", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}

// func Login(w http.ResponseWriter, r *http.Request) {
// 	body, err := io.ReadAll(r.Body)
// 	if err != nil {
// 		http.Error(w, "Error in ready request body", http.StatusBadRequest)
// 		return
// 	}
// }
