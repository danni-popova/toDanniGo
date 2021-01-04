package account

import "net/http"

type Service interface {
	// Authenticate user with provided credentials
	Authenticate(w http.ResponseWriter, r *http.Request)

	// Register a new user
	Register(w http.ResponseWriter, r *http.Request)

	// GetAccountDetails for an existing account identified by ID
	GetAccountDetails(w http.ResponseWriter, r *http.Request)

	// UpdateUserDetails

	// VerifyEmailAddress

	// ResetPassword

}
