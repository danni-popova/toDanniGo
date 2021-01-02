package todo

import (
	"database/sql"
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

type ToDo struct {
	UserID      int            `json:"userID" db:"user_id"`
	ID          int            `json:"id"     db:"id"`
	Title       string         `json:"title"  db:"title"`
	Description string         `json:"description,omitempty" db:"description"`
	Done        bool           `json:"done"   db:"done"`
	CreatedAt   time.Time      `json:"createdAt,omitempty" db:"created_at"`
	UpdatedAt   sql.NullString `json:"updatedAt" db:"updated_at"`
	DeletedAt   sql.NullString `json:"deletedAt" db:"deleted_at"`
	Deadline    sql.NullString `json:"deadline,omitempty" db:"deadline"`
}

func (r *repository) Create(ctd ToDo) (td ToDo, err error) {
	result, err := r.db.NamedQuery(`INSERT INTO todo(user_id, title, description, deadline)
										VALUES (:user_id, :title, :description, :deadline)
										RETURNING *;`, &ctd)
	if err != nil {
		return td, err
	}

	result.Next()
	err = result.StructScan(&td)
	return td, err
}

func (r *repository) Get(todoID, userID int) (td ToDo, err error) {
	var tds []ToDo
	if err = r.db.Select(&tds,
		`SELECT * FROM todo WHERE id=$1 AND user_id=$2`,
		todoID, userID); err != nil {
		return ToDo{}, err
	}
	td = tds[0]
	return
}

func (r *repository) List(userID int) (td []ToDo, err error) {
	if err = r.db.Select(&td, "SELECT * FROM todo WHERE user_id=$1;", userID); err != nil {
		return td, err
	}
	return td, nil
}

// TODO(danni): Allow to update other columns of todos
func (r *repository) Update(todoID, userID int) (err error) {
	_, err = r.db.Exec(`UPDATE todo SET done=$1 WHERE id=$2 AND user_id=$3`,
		true, todoID, userID)

	if err != nil {
		return err
	}
	return err
}

func (r *repository) Delete(todoID, userID int) error {
	_, err := r.db.Exec("DELETE FROM todo WHERE id=$1 AND user_id=$2", todoID, userID)
	if err != nil {
		return err
	}
	return nil
}
