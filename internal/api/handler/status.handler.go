package api

import (
	"fmt"
	"net/http"
)

func SetupRoutesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server set up and running as expected")
}
