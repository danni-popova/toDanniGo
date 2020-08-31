package todo

type Repository interface {
	Create(ctd ToDo) (err error)
	Get(id int) (td ToDo, err error)
	List() (td []ToDo, err error)
	Update(otd ToDo) (ntd ToDo, err error)
	Delete(id string) error
}
