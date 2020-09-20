package todo

import "time"

// Stick all of the request and response structures somewhere e.g here
type GetRequest struct {
	ID int `json:"todo_id"`
}

type Response struct {
	ID          int       `json:"todo_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	Deadline    time.Time `json:"deadline"`
	CreatedAt   time.Time `json:"created_at"`
}

type ListRequest struct {
	Done bool `json:"done"`
}

type ListResponse struct {
	Response []Response
}

type CreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Deadline    time.Time
}

type UpdateRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Done        bool      `json:"done"`
}

type DeleteRequest struct {
	ID int `json:"id"`
}

type UnsuccessfulResponse struct {
	Error string `json:"error"`
}
