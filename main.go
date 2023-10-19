package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// função principal
func main() {

	// Populate the fake database ->
	people = append(people, Person{ID: "1", Firstname: "Carlos", Lastname: "Silva", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Luiz", Lastname: "Vinicius", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Vinicius", Lastname: "Alberto", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "4", Firstname: "Ana", Lastname: "Luiza", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "5", Firstname: "Barbara", Lastname: "Santos", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "6", Firstname: "Eduarda", Lastname: "Silva", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "7", Firstname: "Andre", Lastname: "Santos"})
	// <- Fake DB

	router := mux.NewRouter()

	router.HandleFunc("/contact", GetPeople).Methods("GET")
	router.HandleFunc("/contact/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/contact/", CreatePerson).Methods("POST")
	router.HandleFunc("/contact/{id}", UpdatePerson).Methods("POST")
	router.HandleFunc("/contact/{id}", DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}
func GetPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
		//json.NewEncoder(w).Encode(&Person{})      Gerando um campo vazio pra cada iteracao falsa do forIf
	}
}
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newPerson Person
	json.NewDecoder(r.Body).Decode(&newPerson)
	newPerson.ID = strconv.Itoa(len(people) + 1)
	people = append(people, newPerson)
	json.NewEncoder(w).Encode(newPerson)
}
func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			var newPerson Person
			json.NewDecoder(r.Body).Decode(&newPerson)
			newPerson.ID = params["id"]
			people = append(people, newPerson)
			json.NewEncoder(w).Encode(newPerson)
			return
		}
	}
}
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

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

/// my file contains 73 lines before adapting the last part (u from CRuD) and also pretty
