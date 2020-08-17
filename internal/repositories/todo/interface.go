package todo

import (
	"github.com/todannigo/internal/services/todo"
)

type Repository interface {
	Create(ctd todo.ToDo) (err error)
	Get(id string) (td todo.ToDo, err error)
	List() (td []todo.ToDo, err error)
	Update(otd todo.ToDo) (ntd todo.ToDo, err error)
	Delete(id string) error
}
