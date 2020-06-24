package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "docker"
	dbname   = "todo"
)

type ToDoRequest struct {
	UserID      int       `json:"user_id" db:"user_id"`
	ID          int       `json:"id" db:"todo_id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description,omitempty" db:"description"`
	Deadline    time.Time `json:"deadline,omitempty" db:"deadline"`
}

type ToDo struct {
	ID          int          `json:"id" db:"todo_id"`
	UserID      int          `json:"user_id" db:"user_id"`
	Title       string       `json:"title" db:"title"`
	Description string       `json:"description,omitempty" db:"description"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	Deadline    sql.NullTime `json:"deadline,omitempty" db:"deadline"`
	Done        bool         `json:"done" db:"done"`
}

func dbCon() (db *sqlx.DB) {
	// Construct database connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	//TODO: Authenticate

	// Retrieve from database
	// TODO: When auth is implemented, filter by user
	db := dbCon()
	var td []ToDo
	err := db.Select(&td, "SELECT * FROM todo;")
	if err != nil {
		fmt.Println(err)
	}
	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// TODO: Check for error when marshaling
	marshalled, err := json.Marshal(td)
	w.Write([]byte(marshalled))
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	//TODO: Authenticate

	//TODO: Validate request - has to contain a title
	// user and user has to exist in the db
	reqBody, _ := ioutil.ReadAll(r.Body)
	var td ToDoRequest
	json.Unmarshal(reqBody, &td)

	// Write to database
	db := dbCon()
	result, err := db.NamedQuery(`INSERT INTO todo(user_id, title, description) VALUES (:user_id, :title, :description) RETURNING todo_id;`, &td)
	if err != nil {
		fmt.Println(err)
	}

	var lastID int
	result.Next() // I hate this line, but it's needed for the Scan to work
	err = result.Scan(&lastID)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf(`{ "id" : %d }`, lastID)))
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	//TODO: Authenticate

	// Validate request
	reqBody, _ := ioutil.ReadAll(r.Body)
	var td ToDoRequest
	json.Unmarshal(reqBody, &td)

	// Build query
	var sb strings.Builder
	sb.WriteString(`UPDATE todo SET `)

	if td.Title != "" {
		sb.WriteString(" title=" + td.Title)
	}

	if td.Description != "" {
		sb.WriteString(" description=" + td.Description)
	}

	if !td.Deadline.IsZero() {
		sb.WriteString(" deadline=" + td.Deadline.String())
	}

	sb.WriteString(" WHERE todo_id=" + string(td.ID) + ";")

	// Write to database
	db := dbCon()
	_, err := db.Query(sb.String())
	if err != nil {
		fmt.Println(err)
	}

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(`{"message": "updateTodo called"}`))
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	//TODO: Authenticate

	// Validate request

	// Retrieve from database

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "deleteTodo called"}`))
}

func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/", getTodos).Methods(http.MethodGet)
	r.HandleFunc("/", createTodo).Methods(http.MethodPost)
	r.HandleFunc("/", updateTodo).Methods(http.MethodPut)
	r.HandleFunc("/", deleteTodo).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8080", r))
}
