package api

import (
	"net/http"

	api "github.com/innvtseeds/wdic-server/internal/api/handler"
)

func UserRoutes() {
	http.HandleFunc("/user", createUser)
}

// ServerStatus is a function that handles the /devtest route
func createUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		api.CreateUserHandler(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
