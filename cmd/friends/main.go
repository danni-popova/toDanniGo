package main

import (
	"net/http"
	"os"

	"github.com/danni-popova/todannigo/internal/databases/sql"
	friendsRepo "github.com/danni-popova/todannigo/internal/databases/sql"
	"github.com/danni-popova/todannigo/internal/services/friends"
	"github.com/danni-popova/todannigo/internal/services/middleware"

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
	var svc friends.Service
	svc = friends.NewService(friendsRepo.NewRepository(db))

	// Setup router
	r := mux.NewRouter()
	r.Use(middleware.Middleware)
	r.Use(middleware.LoggingMiddleware)

	api := r.PathPrefix("/friends").Subrouter()
	api.HandleFunc("/", svc.List).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8081", r))
}
