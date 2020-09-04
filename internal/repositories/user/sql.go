package user

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"time"
)

type repository struct {
	db *sqlx.DB
}

func (r repository) Create(u User) (err error) {
	panic("implement me")
}

func (r repository) Get(id int) (u User, err error) {
	panic("implement me")
}

func (r repository) List() (u []User, err error) {
	panic("implement me")
}

func (r repository) Update(u User) (usr User, err error) {
	panic("implement me")
}

func (r repository) Delete(id int) error {
	panic("implement me")
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

type User struct {
	UserID    int       `json:"id,omitempty" db:"user_id"`
	Email     string    `json:"email,omitempty" db:"email"`
	FirstName string    `json:"first_name,omitempty" db:"first_name"`
	LastName  string    `json:"last_name,omitempty" db:"last_name"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
}
