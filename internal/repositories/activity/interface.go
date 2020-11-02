package activity

type Repository interface {
	// List actions in DB
	List(userID int) ([]Action, error)

	// Add action to DB
	Add(action Action) error
}
