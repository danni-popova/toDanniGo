package todo

import (
	"context"
)

type Service interface {
	// Get the details for a single todanni identified by the given ID.
	Get(context.Context, *GetRequest) (*ToDoResponse, error)

	// List all todannis.
	List(context.Context, *ListRequest) (*ListResponse, error)

	// Create a todanni.
	Create(context.Context, *CreateRequest) (*ToDoResponse, error)

	// Update a todanni.
	Update(context.Context, *UpdateRequest) (*ToDoResponse, error)

	// Delete a todanni identified by the given ID.
	Delete(context.Context, *DeleteRequest) error
}
