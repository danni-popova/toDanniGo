package user

type Repository interface {
	Create(u User) (err error)
	Get(id int) (u User, err error)
	List() (u []User, err error)
	Update(u User) (usr User, err error)
	Delete(id int) error
}
