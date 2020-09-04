package main

import (
	"fmt"

	"github.com/danni-popova/todannigo/internal/databases/sql"
	todoRepo "github.com/danni-popova/todannigo/internal/repositories/todo"
	"github.com/danni-popova/todannigo/internal/services/todo"
	"github.com/gorilla/mux"

	"log"
	"net/http"
	"os"
)

func main() {
	// Setup database
	db, err := sql.NewFromEnv()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Setup the service
	var svc todo.Service
	svc = todo.NewService(todoRepo.NewRepository(db))

	// Setup router
	r := mux.NewRouter()
	r.HandleFunc("/", svc.CreateHttp).Methods(http.MethodPost)
	r.HandleFunc("/{id}", svc.GetHttp).Methods(http.MethodGet)
	r.HandleFunc("/", svc.ListHttp).Methods(http.MethodGet)
	r.HandleFunc("/{id}", svc.UpdateHttp).Methods(http.MethodPatch)
	r.HandleFunc("/{id}", svc.DeleteHttp).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8080", r))
}
