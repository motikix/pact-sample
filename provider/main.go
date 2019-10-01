package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/syuni/pact-sample/provider/handlers"
	"github.com/syuni/pact-sample/provider/store"
)

func main() {
	// setup
	if err := store.Load(); err != nil {
		log.Fatalf("error occurred when load db. cause: %s", err.Error())
	}
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.ListPerson).Methods("GET")
	r.HandleFunc("/{id}", handlers.GetPerson).Methods("GET")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("error occurred when launch server. cause: %s", err.Error())
	}
}
