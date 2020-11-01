package activity

import "github.com/jmoiron/sqlx"

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

type Action struct {
}

func (r *repository) List(userID int) ([]Action, error) {
	panic("Implement me")
}
