package api

import "github.com/gorilla/mux"

func SetupRoutes(r *mux.Router) {
	StatusSetup(r)
	UserRoutes(r)
	AuthRoutes(r)
}
