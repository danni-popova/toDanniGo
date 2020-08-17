package todo

import (
	"context"
)

type Service interface {
	// Get the details for a single todanni identified by the given ID.
	Get(context.Context, *GetRequest) (err error, td ToDo)

	// Get all todannis.
	List(context.Context, *ListRequest) (err error, td []ToDo)

	// Create a todanni.
	Create(context.Context, *CreateRequest) (ToDo, error)

	// Update a todanni.
	Update(context.Context, *UpdateRequest) error

	// Delete a todanni identified by the given ID.
	Delete(context.Context, *DeleteRequest) error
}
