package main

import (
	"net/http"
	"os"

	"github.com/danni-popova/todannigo/internal/databases/sql"
	activityRepo "github.com/danni-popova/todannigo/internal/repositories/activity"
	"github.com/danni-popova/todannigo/internal/services/activity"
	"github.com/danni-popova/todannigo/internal/services/middleware"
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
	var svc activity.Service
	svc = activity.NewService(activityRepo.NewRepository(db))

	// Setup router
	r := mux.NewRouter()
	r.Use(middleware.Middleware)
	r.Use(middleware.LoggingMiddleware)

	api := r.PathPrefix("/activity").Subrouter()
	api.HandleFunc("/", svc.ListActions).Methods(http.MethodGet)
	api.HandleFunc("/", svc.RecordAction).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":8082", r))
}
