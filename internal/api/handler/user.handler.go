package api

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	userHandlerDTO "github.com/innvtseeds/wdic-server/internal/dto/handler/user"
	userServiceDTO "github.com/innvtseeds/wdic-server/internal/dto/service/user"
	sharedDTO "github.com/innvtseeds/wdic-server/internal/dto/shared"

	"github.com/innvtseeds/wdic-server/internal/service"
	apiResponse "github.com/innvtseeds/wdic-server/library/standardization"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	var requestBody userHandlerDTO.CreateUser_RequestBody

	// Unmarshal the JSON into the struct
	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		http.Error(w, "Error unmarshalling JSON", http.StatusBadRequest)
		return
	}

	createUserBody := userServiceDTO.UserCreate_RequestBody{
		Email:    requestBody.Email,
		Password: requestBody.Password,
	}

	response, err := service.CreateUserService(&createUserBody)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)

}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	queries := r.URL.Query()
	pageStr := queries.Get("page")
	perPageStr := queries.Get("per_page")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage <= 0 {
		perPage = 10
	}

	myLogger.Info("PAGINATION", page, perPage)

	pagination := sharedDTO.PaginationStruct{
		Page:     int64(page),
		PageSize: int64(perPage),
	}

	users, err := service.GetUsers(&pagination)
	if err != nil {
		apiResponse.ServiceResponseError(w, err)
	}

	apiResponse.StandardResponse(w, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId := vars["userId"]

	myLogger.Info("USERID", userId)

	apiResponse.StandardResponse(w, userId)
}
