package todo

import "time"

type GetRequest struct {
	ID string
}

type ToDoResponse struct {
}

type ListRequest struct {
	Done bool
}

type ListResponse struct {
	Response []ToDoResponse
}

type CreateRequest struct {
	Title       string
	Description string
	Deadline    time.Time
}

type UpdateRequest struct {
	Title       string
	Description string
	Deadline    time.Time
	Done        bool
}

type DeleteRequest struct {
	ID string
}
