package email

import "github.com/jmoiron/sqlx"

type repository struct {
	db *sqlx.DB
}

type VerificationRecord struct {
	ID     int
	UserID int
	UUID   string
}

func (r repository) InsertVerificationRecord(record VerificationRecord) error {
	panic("implement me")
}

func (r repository) UpdateVerificationRecord(recordUUID string) error {
	panic("implement me")
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}
