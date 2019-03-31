package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tchiak/rest-api/handlers"
)

// our main function
func main() {

	router := mux.NewRouter()
	handlers.InitDb()
	router.HandleFunc("/people", handlers.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", handlers.GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", handlers.CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", handlers.DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))

}
