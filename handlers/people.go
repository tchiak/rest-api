package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *address `json:"address,omitempty"`
}
type address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []person

// InitDb s
func InitDb() {
	people = append(people, person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &address{City: "City X", State: "State X"}})
	people = append(people, person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &address{City: "City Z", State: "State Y"}})
	people = append(people, person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})
}

// GetPeople s
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

// GetPerson d
func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

// CreatePerson d
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// DeletePerson d
func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
		}
		break
	}
	json.NewEncoder(w).Encode(people)
}
