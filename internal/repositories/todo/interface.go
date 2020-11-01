package todo

type Repository interface {
	Create(ctd ToDo) (td ToDo, err error)
	Get(todoID, userID int) (td ToDo, err error)
	List(userID int) (td []ToDo, err error)
	Update(todoID, userID int) (err error)
	Delete(todoID, userID int) error
}
