package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// função principal
func main() {

	// Populate the fake database ->
	people = append(people, Person{ID: "1", Firstname: "Carlos", Lastname: "Silva", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Luiz", Lastname: "Roberto", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Andre", Lastname: "Santos"})
	// <- Fake DB

	router := mux.NewRouter()

	router.HandleFunc("/contact", GetPeople).Methods("GET")
	router.HandleFunc("/contact/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/contact/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/contact/{id}", DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}
func GetPerson(w http.ResponseWriter, r *http.Request)    {}
func CreatePerson(w http.ResponseWriter, r *http.Request) {}
func DeletePerson(w http.ResponseWriter, r *http.Request) {}

type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person
