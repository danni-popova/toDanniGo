package todo

import "time"

type Repository interface {
	Create(cr CreateToDoRequest) error
	Get(id string) (ToDo, error)
	List() ([]ToDo, error)
	Update(ur UpdateToDoRequest) (error, td ToDo)
	Delete(id string) error
}

type CreateToDoRequest struct {
	Title       string
	Description string
	Deadline    time.Time
}

type UpdateToDoRequest struct {
	Title       string
	Description string
	Deadline    time.Time
	Done        bool
}

type ToDo struct {
	UserID      int       `json:"user_id"`
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	CreatedAt   time.Time `json:"created_at"`
	Done        bool      `json:"done"`
}
