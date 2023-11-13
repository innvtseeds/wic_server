package api

import (
	"net/http"

	api "github.com/innvtseeds/wdic-server/internal/api/handler"
)

func StatusSetup() {
	http.HandleFunc("/devtest", ServerStatus)
}

// ServerStatus is a function that handles the /devtest route
func ServerStatus(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		api.SetupRoutesHandler(w, r)
	case "POST":
		// do nothing
	case "PUT":
		// do nothing
	case "DELETE":
		// do nothing
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
