package todo

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

type ToDo struct {
	UserID      int       `json:"user_id,omitempty" db:"user_id"`
	ID          int       `json:"todo_id,omitempty" db:"todo_id"`
	Title       string    `json:"title,omitempty" db:"title"`
	Description string    `json:"description,omitempty" db:"description"`
	Deadline    time.Time `json:"deadline,omitempty" db:"deadline"`
	CreatedAt   time.Time `json:"created_at,omitempty" db:"created_at"`
	Done        bool      `json:"done" db:"done"`
}

func (r *repository) Create(ctd ToDo) error {
	// TODO(danni): Refactor later to return the created todo
	if _, err := r.db.NamedQuery(`INSERT INTO todo(user_id, title, description, deadline)
										VALUES (:user_id, :title, :description, :deadline)
										RETURNING todo_id;`, &ctd); err != nil {
		return err
	}

	return nil
}

func (r *repository) Get(id int) (td ToDo, err error) {
	var tds []ToDo
	if err = r.db.Select(&tds, "SELECT * FROM todo WHERE todo_id=$1", id); err != nil {
		return ToDo{}, err
	}
	td = tds[0]
	return
}

func (r *repository) List() (td []ToDo, err error) {
	if err = r.db.Select(&td, "SELECT * FROM todo;"); err != nil {
		return td, err
	}
	return td, nil
}

// TODO(danni): Implement in a non-gross way
func (r *repository) Update(otd ToDo) (ntd ToDo, err error) {
	panic("implement me")
}

func (r *repository) Delete(id int) error {
	_, err := r.db.Query("DELETE FROM todo WHERE todo_id=$1", id)
	if err != nil {
		return err
	}
	return nil
}
