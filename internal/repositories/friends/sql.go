package friends

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type repository struct {
	db *sqlx.DB
}

func (r repository) List(userID int) (f []Friendship, err error) {
	if err = r.db.Select(&f, "SELECT * FROM friends WHERE user_id=$1;", userID); err != nil {
		return f, err
	}
	return f, nil
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

type Friendship struct {
	ID       int    `json:"id" db:"id"`
	FriendID int    `json:"friendID" db:"friend_id"`
	Type     string `json:"type" db:"type"`
	Rejected bool   `json:"rejected" db:"rejected"`
}
