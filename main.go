package main

import (
	"log"
	"net/http"
)

const (
	addr = "0.0.0.0:80"
)

func errorHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("error: %v", r.URL)
	http.Error(w, "Error handler", 500)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("not found: %v", r.URL)
	http.Error(w, "Not found handler", 404)
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("ok: %v", r.URL)
	w.Write([]byte("ok"))
}

func main() {
	server := &http.Server{Addr: addr}

	http.HandleFunc("/error/", errorHandler)
	http.HandleFunc("/notfound/", notFoundHandler)
	http.HandleFunc("/", okHandler)

	log.Printf("Listening on %v", addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
