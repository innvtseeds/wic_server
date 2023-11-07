package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/innvtseeds/wdic-server/internal/config"
	"github.com/joho/godotenv"
)

func main() {

	http.HandleFunc("/devtest", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server set up and running as expected")
	})

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
