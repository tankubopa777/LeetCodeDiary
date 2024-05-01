package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// postgres://jvbmquku:pY8a7Y5BphEt1ICemQn4YjmOlP5PSECo@rain.db.elephantsql.com/jvbmquku
	url := "DATABASE_URL"
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	row := db.QueryRow("INSERT INTO users (name, age) values ($1, $2) RETURNING id", "tansan", 20)
	var id int
	err = row.Scan(&id)
	if err != nil {
		log.Fatal("can't scan id", err)
	}
	log.Println("Inserted row id:", id)
}