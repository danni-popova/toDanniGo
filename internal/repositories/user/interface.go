package user

type Repository interface {
	Create(u User) (err error)
	GetPassword(email string) (pass string, err error)
	GetByID(id int) (u User, err error)
	GetByEmail(email string) (u User, err error)
}
