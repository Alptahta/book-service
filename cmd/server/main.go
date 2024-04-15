package main

import (
	"fmt"
	"log"
	"net/http"
)

const HTTP_SERVER_PORT = 8080

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", HTTP_SERVER_PORT),
		Handler: mux,
	}

	log.Printf("Starting HTTP server at port %d", HTTP_SERVER_PORT)
	log.Fatal(srv.ListenAndServe())
}
