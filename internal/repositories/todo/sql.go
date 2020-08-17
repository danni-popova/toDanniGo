package todo

import (
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
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

func (r *repository) Get(id string) (td ToDo, err error) {
	if err = r.db.Select(&td, "SELECT * FROM todo WHERE todo_id=$1;", id); err != nil {
		return ToDo{}, err
	}
	return td, nil
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

func (r *repository) Delete(id string) error {
	_, err := r.db.Query("DELETE FROM todo WHERE todo_id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}
