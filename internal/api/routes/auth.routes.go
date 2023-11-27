package api

import (
	"github.com/gorilla/mux"
	api "github.com/innvtseeds/wdic-server/internal/api/handler"
)

func AuthRoutes(r *mux.Router) {
	r.HandleFunc("/auth/register", api.Register).Methods("POST")
	r.HandleFunc("/auth/login", api.Login).Methods("POST")
	r.HandleFunc("/auth/test", api.Test).Methods("POST")
}
