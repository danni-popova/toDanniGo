package todo

import (
	"net/http"
)

type Service interface {
	// Get the details for a single todanni identified by the given ID.
	GetHttp(w http.ResponseWriter, r *http.Request)

	// List all todannis.
	ListHttp(w http.ResponseWriter, r *http.Request)

	// Create a todanni.
	CreateHttp(w http.ResponseWriter, r *http.Request)

	// Update a todanni.
	UpdateHttp(w http.ResponseWriter, r *http.Request)

	// Delete a todanni identified by the given ID.
	DeleteHttp(w http.ResponseWriter, r *http.Request)
}
