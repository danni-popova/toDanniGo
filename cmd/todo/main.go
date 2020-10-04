package main

import (
	"net/http"
	"os"

	"github.com/danni-popova/todannigo/internal/services/middleware"

	"github.com/danni-popova/todannigo/internal/databases/sql"
	todoRepo "github.com/danni-popova/todannigo/internal/repositories/todo"
	"github.com/danni-popova/todannigo/internal/services/todo"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Setup database
	db, err := sql.NewFromEnv()
	if err != nil {
		//fmt.Println(err)
		log.Error(err)
		os.Exit(1)
	}

	// Setup the service
	var svc todo.Service
	svc = todo.NewService(todoRepo.NewRepository(db))

	// Setup the middleware

	// Setup router
	r := mux.NewRouter()
	r.Use(middleware.Middleware)

	api := r.PathPrefix("/todo").Subrouter()
	api.HandleFunc("/", svc.CreateHttp).Methods(http.MethodPost)
	api.HandleFunc("/{id}", svc.GetHttp).Methods(http.MethodGet)
	api.HandleFunc("/", svc.ListHttp).Methods(http.MethodGet)
	api.HandleFunc("/{id}", svc.UpdateHttp).Methods(http.MethodPatch)
	api.HandleFunc("/{id}", svc.DeleteHttp).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8081", r))
}
