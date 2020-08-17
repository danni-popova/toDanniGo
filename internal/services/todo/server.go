//package todo
//
//import (
//	"encoding/json"
//	"fmt"
//	"io/ioutil"
//	"log"
//	"net/http"
//	"strings"
//	"time"
//
//	"github.com/jmoiron/sqlx"
//	_ "github.com/lib/pq"
//
//	"github.com/gorilla/mux"
//)
//
//const (
//	host     = "localhost"
//	port     = 5432
//	user     = "postgres"
//	password = "docker"
//	dbname   = "todo"
//)
//
//type ToDo struct {
//	UserID      int       `json:"user_id,omitempty" db:"user_id"`
//	ID          int       `json:"id,omitempty" db:"todo_id"`
//	Title       string    `json:"title, omitempty" db:"title"`
//	Description string    `json:"description,omitempty" db:"description"`
//	Deadline    time.Time `json:"deadline,omitempty" db:"deadline"`
//	CreatedAt   time.Time `json:"created_at,omitempty" db:"created_at"`
//	Done        bool      `json:"done" db:"done"`
//}
//
//func dbCon() (db *sqlx.DB) {
//	// Construct database connection string
//	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
//		"password=%s dbname=%s sslmode=disable",
//		host, port, user, password, dbname)
//
//	db, err := sqlx.Connect("postgres", psqlInfo)
//	if err != nil {
//		fmt.Println(err)
//	}
//	return db
//}
//
//func getTodos(w http.ResponseWriter, r *http.Request) {
//	//TODO: Authenticate
//
//	// Retrieve from database
//	// TODO: When auth is implemented, filter by user
//	db := dbCon()
//	var td []ToDo
//	err := db.Select(&td, "SELECT * FROM todo;")
//	if err != nil {
//		fmt.Println(err)
//	}
//	// Write response
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//
//	// TODO: Check for error when marshaling
//	marshalled, err := json.Marshal(td)
//	w.Write([]byte(marshalled))
//}
//
//func getTodo(w http.ResponseWriter, r *http.Request) {
//	//TODO: Authenticate
//	pathParams := mux.Vars(r)
//	rID := pathParams["todo_id"]
//
//	// Retrieve from database
//	db := dbCon()
//	var td []ToDo
//	sqlQuery := fmt.Sprintf("SELECT * FROM todo WHERE todo_id=%s", rID)
//	err := db.Select(&td, sqlQuery)
//	if err != nil {
//		fmt.Println(err)
//	}
//	// Write response
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusOK)
//
//	// TODO: Check for error when marshaling
//	marshalled, err := json.Marshal(td[0])
//	w.Write([]byte(marshalled))
//}
//
//func createTodo(w http.ResponseWriter, r *http.Request) {
//	reqBody, _ := ioutil.ReadAll(r.Body)
//	var td ToDo
//	err := json.Unmarshal(reqBody, &td)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	// Write to database
//	db := dbCon()
//	result, err := db.NamedQuery(`INSERT INTO todo(user_id, title, description, deadline)
//										VALUES (:user_id, :title, :description, :deadline)
//										RETURNING todo_id;`, &td)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	result.Next()
//	var lastID int
//	err = result.Scan(&lastID)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusCreated)
//	w.Write([]byte(fmt.Sprintf(`{ "id" : %d }`, lastID)))
//}
//
//func updateTodo(w http.ResponseWriter, r *http.Request) {
//	reqBody, _ := ioutil.ReadAll(r.Body)
//	var toDoRequest ToDo
//	err := json.Unmarshal(reqBody, &toDoRequest)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	pathParams := mux.Vars(r)
//	rID := pathParams["todo_id"]
//
//	var statements []string
//	var params []interface{}
//	var baseQuery = "UPDATE todo SET %v WHERE todo_id=$%d"
//
//	if toDoRequest.Title != "" {
//		statements = append(statements, fmt.Sprintf("title=$%d", len(statements)+1))
//		params = append(params, toDoRequest.Title)
//	}
//	if toDoRequest.Description != "" {
//		statements = append(statements, fmt.Sprintf("description=$%d", len(statements)+1))
//		params = append(params, toDoRequest.Description)
//	}
//	if !toDoRequest.Deadline.IsZero() {
//		statements = append(statements, fmt.Sprintf("deadline=$%d", len(statements)+1))
//		params = append(params, toDoRequest.Deadline)
//	}
//	if toDoRequest.Done {
//		statements = append(statements, fmt.Sprintf("done=$%d", len(statements)+1))
//		params = append(params, true)
//	}
//
//	params = append(params, rID)
//	updateStatements := strings.Join(statements, ", ")
//	querySql := fmt.Sprintf(baseQuery, updateStatements, len(statements)+1)
//
//	db := dbCon()
//	result, err := db.Query(querySql, params...)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	result.Next()
//	var returned ToDo
//	err = result.Scan(&returned)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	// Write response
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(http.StatusAccepted)
//	marshalled, err := json.Marshal(returned)
//	w.Write([]byte(marshalled))
//}
//
//func deleteTodo(w http.ResponseWriter, r *http.Request) {
//	//TODO: Authenticate
//
//	// Validate request
//	pathParams := mux.Vars(r)
//	toDoID := pathParams["todo_id"]
//
//	// Retrieve from database
//	// TODO: check that todo is actually the user's
//	db := dbCon()
//	_, err := db.Query("DELETE FROM todo WHERE todo_id=$1", toDoID)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	w.WriteHeader(http.StatusOK)
//}
//
//func StartServer() {
//	r := mux.NewRouter()
//	r.HandleFunc("/", getTodos).Methods(http.MethodGet)
//	r.HandleFunc("/{todo_id}", getTodo).Methods(http.MethodGet)
//	r.HandleFunc("/", createTodo).Methods(http.MethodPost)
//	r.HandleFunc("/{todo_id}", updateTodo).Methods(http.MethodPut)
//	r.HandleFunc("/{todo_id}", deleteTodo).Methods(http.MethodDelete)
//	log.Fatal(http.ListenAndServe(":8080", r))
//}
