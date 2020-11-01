package activity

type Repository interface {
	List(userID int) ([]Action, error)
}
