package main

import (
"database/sql"
"fmt"
_ "github.com/lib/pq"
)
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "docker"
	dbname   = "todo"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("SELECT title FROM todo")
	for rows.Next(){
		var title string
		err = rows.Scan(&title)
		if err != nil{
			panic(err)
		}
		fmt.Println(title)
	}

	defer db.Close()
}