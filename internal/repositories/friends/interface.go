package friends

type Repository interface {
	// List actions in DB
	List(userID int) (f []Friendship, err error)
}
