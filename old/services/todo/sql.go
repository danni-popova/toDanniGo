package main

import (
	"database/sql"
"fmt"
_ "github.com/lib/pq"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "docker"
	dbname   = "todo"
)

type ToDo struct {
	ID 			int 		`json:"id"`
	UserID 		int 		`json:"user_id"`
	Title 		string 		`json:"title"`
	Description string 		`json:"description"`
	CreatedAt 	time.Time 	`json:"created_at"`
	Deadline	time.Time 	`json:"deadline"`
	Done 		bool 		`json:"done"`
}

func getToDoByUser(userID string) ([]ToDo, error){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var returnedToDo []ToDo

	rows, err := db.Query("SELECT * FROM todo WHERE user_id = $1", userID)
	for rows.Next(){
		var td ToDo
		err = rows.Scan(&td.ID,
						&td.UserID,
						&td.Title,
						&td.Description,
						&td.CreatedAt,
						&td.Deadline,
						&td.Done)

		if err != nil{
			panic(err)
		}
		returnedToDo = append(returnedToDo, td)
	}
	return returnedToDo, nil
}

func deleteToDoFromDB(id string) error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//_, err == db.Query("DELETE FROM todo WHERE todo_id = $1", id)

	return nil
}
