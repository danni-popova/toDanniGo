package todo

import "github.com/todannigo/internal/repositories/todo"

type ToDo struct {
	repo todo.Repository
}

func (td *ToDo) Create() (err error) {

	return err
}
