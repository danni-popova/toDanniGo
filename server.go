package main

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
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
	UserID      int       `json:"user_id,omitempty" db:"user_id"`
	ID          int       `json:"id,omitempty" db:"todo_id"`
	Title       string    `json:"title, omitempty" db:"title"`
	Description string    `json:"description,omitempty" db:"description"`
	Deadline    time.Time `json:"deadline,omitempty" db:"deadline"`
	CreatedAt   time.Time `json:"created_at,omitempty" db:"created_at"`
	Done        bool      `json:"done" db:"done"`
}

type ToDo struct {
	ID          int       `json:"id" db:"todo_id"`
	UserID      int       `json:"user_id" db:"user_id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description,omitempty" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	Deadline    time.Time `json:"deadline,omitempty" db:"deadline"`
	Done        bool      `json:"done" db:"done"`
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
	reqBody, _ := ioutil.ReadAll(r.Body)
	var td ToDoRequest
	err := json.Unmarshal(reqBody, &td)
	if err != nil {
		fmt.Println(err)
	}

	// Write to database
	db := dbCon()
	result, err := db.NamedQuery(`INSERT INTO todo(user_id, title, description, deadline) 
										VALUES (:user_id, :title, :description, :deadline) 
										RETURNING todo_id;`, &td)
	if err != nil {
		fmt.Println(err)
	}

	result.Next()
	var lastID int
	err = result.Scan(&lastID)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf(`{ "id" : %d }`, lastID)))
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var toDoRequest ToDoRequest
	err := json.Unmarshal(reqBody, &toDoRequest)
	if err != nil {
		fmt.Println(err)
	}

	// Get the id from the request
	pathParams := mux.Vars(r)
	rID := pathParams["todo_id"]

	var oldTd ToDoRequest
	db := dbCon()
	err = db.QueryRowx("SELECT * FROM todo WHERE todo_id=$1", rID).StructScan(&oldTd)
	if err != nil {
		fmt.Println(err)
	}

	if toDoRequest.Title != "" {
		oldTd.Title = toDoRequest.Title
	}
	if toDoRequest.Description != "" {
		oldTd.Description = toDoRequest.Description
	}
	if !toDoRequest.Deadline.IsZero() {
		oldTd.Deadline = toDoRequest.Deadline
	}
	if toDoRequest.Done {
		oldTd.Done = true
	}

	// Write to database
	_, err = db.NamedQuery(`UPDATE todo 
								  SET title=:title, 
                				  description=:description, 
                				  deadline=:deadline,
								  done=:done
                				  WHERE todo_id=:todo_id`,
		&oldTd)
	if err != nil {
		fmt.Println(err)
	}

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	marshalled, err := json.Marshal(oldTd)
	w.Write([]byte(marshalled))
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	//TODO: Authenticate

	// Validate request
	pathParams := mux.Vars(r)
	toDoID := pathParams["todo_id"]

	// Retrieve from database
	// TODO: check that todo is actually the user's
	db := dbCon()
	_, err := db.Query("DELETE FROM todo WHERE todo_id=$1", toDoID)
	if err != nil {
		fmt.Println(err)
	}

	// Write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func StartServer() {
	r := mux.NewRouter()
	r.HandleFunc("/", getTodos).Methods(http.MethodGet)
	r.HandleFunc("/", createTodo).Methods(http.MethodPost)
	r.HandleFunc("/{todo_id}/", updateTodo).Methods(http.MethodPut)
	r.HandleFunc("/{todo_id}/", deleteTodo).Methods(http.MethodDelete)
	log.Fatal(http.ListenAndServe(":8080", r))
}
