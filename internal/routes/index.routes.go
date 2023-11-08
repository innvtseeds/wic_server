package routes

import (
	"fmt"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/devtest", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server set up and running as expected")
	})
}
