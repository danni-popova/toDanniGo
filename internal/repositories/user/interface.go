package user

type Repository interface {
	Create(u User) (err error)
	GetPassword(email string) (pass string, err error)
	Get(id int) (u User, err error)
}
