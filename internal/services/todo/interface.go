package todo

import (
	"context"
	"net/http"
)

type Service interface {
	// Get the details for a single todanni identified by the given ID.
	Get(context.Context, *GetRequest) (*Response, error)

	// List all todannis.
	List(context.Context, *ListRequest) (*ListResponse, error)

	// Create a todanni.
	Create(context.Context, *CreateRequest) (*Response, error)

	// Update a todanni.
	Update(context.Context, *UpdateRequest) (*Response, error)

	// Delete a todanni identified by the given ID.
	Delete(context.Context, *DeleteRequest) error

	// Before I figure out how to refactor these:
	GetHttp(w http.ResponseWriter, r *http.Request)

	ListHttp(w http.ResponseWriter, r *http.Request)

	CreateHttp(w http.ResponseWriter, r *http.Request)

	UpdateHttp(w http.ResponseWriter, r *http.Request)

	DeleteHttp(w http.ResponseWriter, r *http.Request)
}
