package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/todannigo/internal/databases/sql"
	todo2 "github.com/todannigo/internal/repositories/todo"
	"github.com/todannigo/internal/services/todo"
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
	svc = todo.NewService(todo2.NewRepository(db))

	// Router
	r := mux.NewRouter()
	r.HandleFunc("/{id}", svc.GetHttp).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", r))
}
