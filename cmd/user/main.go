package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/danni-popova/todannigo/internal/databases/sql"
	userRepo "github.com/danni-popova/todannigo/internal/repositories/user"
	"github.com/danni-popova/todannigo/internal/services/user"
	"github.com/gorilla/mux"
)

func main() {
	// Setup database
	db, err := sql.NewFromEnv()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Setup the service
	var svc user.Service
	svc = user.NewService(userRepo.NewRepository(db))

	// Setup router
	r := mux.NewRouter()
	r.HandleFunc("/login", svc.Login).Methods(http.MethodPost)
	r.HandleFunc("/register", svc.Register).Methods(http.MethodPost)
	r.HandleFunc("/user/{id}", svc.GetUser).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", r))
}
