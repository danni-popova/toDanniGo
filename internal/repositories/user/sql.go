package user

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

type User struct {
	UserID    int       `json:"user_id" db:"user_id"`
	Password  string    `json:"password" db:"password"`
	Email     string    `json:"email" db:"email"`
	FirstName string    `json:"first_name,omitempty" db:"first_name"`
	LastName  string    `json:"last_name,omitempty" db:"last_name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (r *repository) Create(u User) error {
	if _, err := r.db.NamedQuery(`INSERT INTO registered_user(first_name, last_name, email, password)
										VALUES (:first_name, :last_name, :email, :password)
										RETURNING user_id;`, &u); err != nil {
		return err
	}
	return nil
}

func (r *repository) Get(id int) (u User, err error) {
	if err = r.db.Select(&u, "SELECT * FROM registered_user;"); err != nil {
		return u, err
	}
	return
}

func (r *repository) GetPassword(email string) (password string, err error) {
	var user []User
	err = r.db.Select(&user, "SELECT password FROM registered_user WHERE email=$1;", email)
	if err != nil {
		return "", err
	}
	return user[0].Password, err
}
