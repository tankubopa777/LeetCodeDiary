package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	url := "DATABASE_URL"
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	createTb := `
	CREATE TABLE IF NOT EXISTS users ( id SERIAL PRIMARY KEY, name TEXT, age INT);
	`
	_, err = db.Exec(createTb)

	if err != nil {
		log.Fatal("Error create table: ", err)
	}


	log.Println("Connected to database")
}