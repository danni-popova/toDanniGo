package email

import "net/http"

type Service interface {
	// SendVerificationEmail sends an email to the specified user after registration
	SendVerificationEmail(w http.ResponseWriter, r *http.Request)

	// SendPasswordResetEmail sends an email with a unique link to reset password
	SendPasswordResetEmail(w http.ResponseWriter, r *http.Request)
}
