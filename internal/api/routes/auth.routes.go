package api

import (
	"net/http"

	api "github.com/innvtseeds/wdic-server/internal/api/handler"
)

func AuthRoutes() {
	http.HandleFunc("/auth/register", register)
	http.HandleFunc("/auth/login", login)
}

func register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		api.Register(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

// Login Function
func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		api.Login(w, r)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
