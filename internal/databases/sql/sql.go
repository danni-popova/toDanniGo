package sql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

// TODO: refactor to use environment variables
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "docker"
	dbname   = "todo"
)

func Open() (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewFromEnv() (*sqlx.DB, error) {
	return Open()
}
