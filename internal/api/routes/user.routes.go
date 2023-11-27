package api

import (
	"github.com/gorilla/mux"
	api "github.com/innvtseeds/wdic-server/internal/api/handler"
)

func UserRoutes(r *mux.Router) {
	r.HandleFunc("/user", api.CreateUserHandler).Methods("POST")
	r.HandleFunc("/user/{userId}", api.GetUser).Methods("GET")
}
