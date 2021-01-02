package main

import (
	"net/http"
	"os"

	"github.com/danni-popova/todannigo/internal/databases/sql"
	projectsRepo "github.com/danni-popova/todannigo/internal/repositories/projects"
	"github.com/danni-popova/todannigo/internal/services/middleware"
	"github.com/danni-popova/todannigo/internal/services/projects"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Setup database
	db, err := sql.NewFromEnv()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// Setup the service
	var svc projects.Service
	svc = projects.NewService(projectsRepo.NewRepository(db))

	// Setup router
	r := mux.NewRouter()
	r.Use(middleware.Middleware)
	r.Use(middleware.LoggingMiddleware)

	api := r.PathPrefix("/api/project").Subrouter()
	api.HandleFunc("/", svc.ListProjects).Methods(http.MethodGet)
	api.HandleFunc("/", svc.CreateProject).Methods(http.MethodPost)
	api.HandleFunc("/{id}", svc.AddMember).Methods(http.MethodPatch)
	log.Fatal(http.ListenAndServe(":8084", r))
}
