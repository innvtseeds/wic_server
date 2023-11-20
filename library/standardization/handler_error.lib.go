package apiResponse

import (
	"encoding/json"
	"net/http"

	customLogger "github.com/innvtseeds/wdic-server/library/logger"
)

var myLogger = customLogger.NewLogger()

// ErrorResponse represents a standardized error response structure.
type ErrorResponse struct {
	ErrorType string `json:"error_type"`
	Message   string `json:"message"`
}

func sendErrorResponse(w http.ResponseWriter, statusCode int, errorType string, messages ...interface{}) {
	errorMessage := ErrorResponse{
		ErrorType: errorType,
		Message:   "Error in request",
	}

	if len(messages) > 0 {
		if msg, ok := messages[0].(string); ok {
			errorMessage.Message = msg
			messages = messages[1:]
		}
	}

	// Log the error
	myLogger.Error(errorType+":", messages)

	// Set the Content-Type header
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response with the error message
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errorMessage)
}

func RequestBodyError(w http.ResponseWriter, messages ...interface{}) {
	sendErrorResponse(w, http.StatusBadRequest, "REQUEST_BODY_ERROR", messages...)
}

func UnmarshalError(w http.ResponseWriter, messages ...interface{}) {
	sendErrorResponse(w, http.StatusBadRequest, "UNMARSHALING_ERROR", messages...)
}

func ServiceResponseError(w http.ResponseWriter, messages ...interface{}) {
	sendErrorResponse(w, http.StatusBadRequest, "SERVICE_RESPONSE_ERROR", messages...)
}
