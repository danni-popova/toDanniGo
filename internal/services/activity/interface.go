package activity

import "net/http"

type Service interface {
	// List all actions for a user
	ListActions(w http.ResponseWriter, r *http.Request)

	// Record action
	RecordAction(w http.ResponseWriter, r *http.Request)
}
