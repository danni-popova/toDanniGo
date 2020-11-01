package todo

import (
	"time"
)

// Stick all of the request and response structures somewhere e.g here
type GetRequest struct {
	ID int `json:"ID"`
}

type Response struct {
	ID          int       `json:"ID"`
	Title       string    `json:"Title"`
	Description string    `json:"Description"`
	Deadline    string    `json:"Deadline"`
	Done        bool      `json:"Done"`
	CreatedAt   time.Time `json:"CreatedAt"`
}

type ListRequest struct {
	Done bool `json:"done"`
}

type ListResponse struct {
	Response []Response
}

type CreateRequest struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Deadline    time.Time
}

type UpdateRequest struct {
	Title       string    `json:"Title"`
	Description string    `json:"Description"`
	Deadline    time.Time `json:"Deadline"`
	Done        bool      `json:"Done"`
}

type DeleteRequest struct {
	ID int `json:"ID"`
}

type UnsuccessfulResponse struct {
	Error string `json:"Error"`
}
