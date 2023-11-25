package apiResponse

import (
	"encoding/json"
	"net/http"
)

type Metadata struct {
	Page       int `json:"page"`
	PerPage    int `json:"perPage"`
	TotalPages int `json:"totalPages"`
}

type ResponseObject struct {
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
	Metadata *Metadata   `json:"metadata,omitempty"`
}

func StandardResponse(w http.ResponseWriter, response ...interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	response_dto := ResponseObject{
		Message: "Success",
		Data:    response,
	}

	if err := json.NewEncoder(w).Encode(response_dto); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

}
