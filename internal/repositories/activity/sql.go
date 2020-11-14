package activity

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

type Action struct {
	ID        int       `json:"id" db:"id"`
	UserID    int       `json:"userID" db:"user_id"`
	Task      string    `json:"task" db:"task"`
	Type      string    `json:"type" db:"type"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
}

func (r *repository) List(userID int) (actions []Action, err error) {
	err = r.db.Select(&actions, "SELECT * FROM action_log WHERE user_id=$1;", userID)
	return actions, err
}

func (r *repository) Add(action Action) (err error) {
	_, err = r.db.Exec(`INSERT INTO action_log(timestamp, type, task, user_id) 
					 VALUES (:timestamp, :type, :task_id, :user_id);`, &action)
	return err
}
