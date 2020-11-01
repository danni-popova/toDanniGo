package user

import (
	"database/sql"
	"errors"
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
	UserID         int            `json:"id" db:"id"`
	Password       string         `json:"password" db:"password"`
	Email          string         `json:"email" db:"email"`
	ProfilePicture string         `json:"profilePicture" db:"profile_picture"`
	CreatedAt      time.Time      `json:"createdAt" db:"created_at"`
	UpdatedAt      sql.NullString `json:"updatedAt" db:"updated_at"`
	DeletedAt      sql.NullString `json:"deletedAt" db:"deleted_at"`
	FirstName      string         `json:"firstName" db:"first_name"`
	LastName       string         `json:"lastName" db:"last_name"`
}

func (r *repository) Create(u User) error {
	if _, err := r.db.NamedQuery(`INSERT INTO registered_user(first_name, last_name, email, password)
										VALUES (:first_name, :last_name, :email, :password)
										RETURNING id;`, &u); err != nil {
		return err
	}
	return nil
}

func (r *repository) GetByID(id int) (u User, err error) {
	var usr []User
	if err = r.db.Select(&usr, "SELECT * FROM registered_user WHERE id=$1;", id); err != nil {
		return u, err
	}

	if len(usr) == 0 {
		return u, errors.New("query returned no results")
	}

	return usr[0], nil
}

func (r *repository) GetByEmail(email string) (u User, err error) {
	var usr []User
	if err = r.db.Select(&usr, "SELECT * FROM registered_user WHERE email=$1;", email); err != nil {
		return u, err
	}

	if len(usr) == 0 {
		return u, errors.New("query returned no results")
	}

	return usr[0], nil
}

func (r *repository) GetPassword(email string) (password string, err error) {
	var user []User
	err = r.db.Select(&user, "SELECT password FROM registered_user WHERE email=$1;", email)
	if err != nil {
		return "", err
	}
	return user[0].Password, err
}
