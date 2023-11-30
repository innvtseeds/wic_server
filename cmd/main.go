package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	api "github.com/innvtseeds/wdic-server/internal/api/routes"
	"github.com/innvtseeds/wdic-server/internal/config"

	"github.com/joho/godotenv"
)

func main() {

	// Set up routes from routes.go
	r := mux.NewRouter()
	api.SetupRoutes(r)

	config.LoadDBConfig()

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	s := &http.Server{
		Addr:           ":" + port,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())

}
