package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
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
var db *sql.DB

// InitDb s
func InitDb() {
	var err error

	db, err = sql.Open("postgres", "postgres://tchiak:12345@localhost:1234/default?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	var msg = addPerson(person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &address{City: "City X", State: "State X"}})
	if msg != "" {
		log.Fatal(msg)
	}
	msg = addPerson(person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &address{City: "City Z", State: "State Y"}})
	if msg != "" {
		log.Fatal(msg)
	}
	msg = addPerson(person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})
	if msg != "" {
		log.Fatal(msg)
	}

	fmt.Println("Intialized")
	people = append(people, person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &address{City: "City X", State: "State X"}})
	people = append(people, person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &address{City: "City Z", State: "State Y"}})
	people = append(people, person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})
}

func addPerson(entry person) (errorMessage string) {
	var err error
	var result sql.Result

	result, err = db.Exec("INSERT INTO public.Person (firstName, lastName) VALUES ( $1, $2 ) RETURNING id", entry.Firstname, entry.Lastname)
	if err != nil {
		fmt.Println("INSERT into person error")
		return err.Error()
	}
	var id, _ = result.LastInsertId()
	fmt.Println(id)
	if err == nil && entry.Address != nil {
		result, err = db.Exec("INSERT INTO Address (city, state, personid) VALUES($1, $2, $3)", entry.Address.City, entry.Address.State, id)
	}
	if err != nil {
		return err.Error()
	}

	return
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
