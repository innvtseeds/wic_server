package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/innvtseeds/wdic-server/internal/config"
	"github.com/innvtseeds/wdic-server/internal/routes"

	"github.com/joho/godotenv"
)

func main() {

	// Set up routes from routes.go
	routes.SetupRoutes()

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
		Addr: ":" + port,
		// Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())

}
