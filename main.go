package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// função principal
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/contact", GetPeople).Methods("GET")
	router.HandleFunc("/contact/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/contact/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/contact/{id}", DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
