package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/danni-popova/todannigo/internal/databases/sql"
	accountRepo "github.com/danni-popova/todannigo/internal/repositories/account"
	"github.com/danni-popova/todannigo/internal/services/account"
	"github.com/danni-popova/todannigo/internal/services/middleware"
	"github.com/gorilla/mux"
)

func main() {
	// Setup database
	db, err := sql.NewFromEnv()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log.Printf("Environment variables: %s", os.Getenv("POSTGRES_USER"))

	// Setup the service
	var svc account.Service
	svc = account.NewService(accountRepo.NewRepository(db))

	// Setup router
	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware)

	api := r.PathPrefix("/api/").Subrouter()
	api.HandleFunc("/authenticate", svc.Authenticate).Methods(http.MethodPost)
	api.HandleFunc("/register", svc.Register).Methods(http.MethodPost)
	api.HandleFunc("/account/{id}", svc.GetAccountDetails).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8081", r))
}
