package user

import (
	"net/http"
)

type Service interface {
	// Login with existing user
	Login(w http.ResponseWriter, r *http.Request)

	// Register new user
	Register(w http.ResponseWriter, r *http.Request)

	// Get user details
	GetUser(w http.ResponseWriter, r *http.Request)
}
