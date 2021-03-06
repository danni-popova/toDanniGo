package sql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

// TODO: refactor to use environment variables
const (
	host     = "postgresql"
	port     = 5432
	user     = "docker"
	password = "docker"
	dbname   = "todo"
)

// Open - creates a connection to the database
func Open() (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return db, nil
}

func NewFromEnv() (*sqlx.DB, error) {
	return Open()
}
