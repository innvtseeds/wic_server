package api

import (
	"github.com/gorilla/mux"
	api "github.com/innvtseeds/wdic-server/internal/api/handler"
)

func StatusSetup(r *mux.Router) {
	r.HandleFunc("/devtest", api.SetupRoutesHandler).Methods("GET")
}
