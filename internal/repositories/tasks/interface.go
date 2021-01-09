package tasks

type Repository interface {
	InsertTask(task Task) (Task, error)
	SelectTasksByProjectID(projectID int) (tasks []Task, err error)
	UpdateTask(task Task) (Task, error)
	DeleteTask(id int) error
}
