package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/devtest", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server set up and running as expected")
	})

	s := &http.Server{
		Addr: ":8080",
		// Handler:        myHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())

}
