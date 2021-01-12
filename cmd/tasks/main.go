package main

import (
	"net/http"
	"os"

	"github.com/danni-popova/todannigo/internal/databases/gorm"
	taskRepo "github.com/danni-popova/todannigo/internal/repositories/tasks"
	"github.com/danni-popova/todannigo/internal/services/middleware"
	"github.com/danni-popova/todannigo/internal/services/tasks"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Setup database
	db, err := gorm.Open()
	if err != nil {
		//fmt.Println(err)
		log.Error(err)
		os.Exit(1)
	}

	// Setup the service
	var svc tasks.Service
	svc = tasks.NewService(taskRepo.NewRepository(db))

	// Setup router
	r := mux.NewRouter()
	r.Use(middleware.Middleware)
	r.Use(middleware.LoggingMiddleware)

	api := r.PathPrefix("/api/task").Subrouter()
	api.HandleFunc("/", svc.Create).Methods(http.MethodPost)
	api.HandleFunc("/", svc.List).Queries("project", "{project}").Methods(http.MethodGet)
	api.HandleFunc("/{id}", svc.Update).Methods(http.MethodPatch)
	api.HandleFunc("/{id}", svc.Delete).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8083", r))
}
