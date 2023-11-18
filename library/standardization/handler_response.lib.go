package apiResponse

import (
	"encoding/json"
	"net/http"
)

func StandardResponse(w http.ResponseWriter, response ...interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}

}
