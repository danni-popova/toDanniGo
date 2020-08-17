package todo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(cr CreateToDoRequest) error {

	return nil
}

func (r *repository) Get(id string) (ToDo, error) {
	var toDo ToDo

	if err := r.db.Get(&toDo,
		"SELECT * FROM todo WHERE todo_id=$1",
		id); err != nil {
		return ToDo{}, fmt.Errorf("sql get: #{err}")
	}

	return toDo, nil
}

func (r *repository) List() ([]ToDo, error) {
	var td []ToDo

	// TODO(danni): Filter on user when auth is implemented
	if err := r.db.Select(&td, "SELECT * FROM todo;"); err != nil {
		return nil, err
	}

	return td, nil
}

func (r *repository) Update(ur UpdateToDoRequest) (error, td ToDo) {
	return nil
}

func (r *repository) Delete(id string) error {
	return nil
}
