package todo

import "time"

type GetRequest struct {
	ID string
}

type ListRequest struct {
	Done bool
}

type ListResponse struct {
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

type ToDo struct {
	UserID      int       `json:"user_id,omitempty" db:"user_id"`
	ID          int       `json:"id,omitempty" db:"todo_id"`
	Title       string    `json:"title,omitempty" db:"title"`
	Description string    `json:"description,omitempty" db:"description"`
	Deadline    time.Time `json:"deadline,omitempty" db:"deadline"`
	CreatedAt   time.Time `json:"created_at,omitempty" db:"created_at"`
	Done        bool      `json:"done" db:"done"`
}
