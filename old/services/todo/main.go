package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func getAllToDos(w http.ResponseWriter, r *http.Request) {

	// Query database
	var todo []ToDo
	todo, err := getToDoByUser("2")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	marshalled, err := json.Marshal(todo)
	w.Write([]byte(marshalled))
}

func createToDo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(``))
}

func updateToDo(w http.ResponseWriter, r *http.Request) {
	// Validate user

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(``))
}

func deleteToDo(w http.ResponseWriter, r *http.Request) {
	// Validate user

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(``))
}

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/todo").Subrouter()
	api.HandleFunc("", getAllToDos).Methods(http.MethodGet)
	api.HandleFunc("", createToDo).Methods(http.MethodPost)
	api.HandleFunc("/{todoid}", updateToDo).Methods(http.MethodPut)
	api.HandleFunc("/{todoid}", deleteToDo).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8080", r))
}