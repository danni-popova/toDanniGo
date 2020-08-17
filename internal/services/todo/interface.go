package todo

import "context"

type Service interface {
	// Get the details for a single todanni
	Get(ctx context.Context, id string) (err error, td ToDo)

	// Get all todannis
	List(ctx context.Context) (err error, td []ToDo)

	// Create a todanni
	Create(ctx context.Context, td ToDo) (td ToDo, err error)

	// Update a todanni
	Update(ctx context.Context, td ToDo) (err error)

	// Delete a todanni
	Delete(ctx context.Context, id string) (err error)
}
