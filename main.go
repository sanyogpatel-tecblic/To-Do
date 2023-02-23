package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func main() {
	var err error
	DB, err = sql.Open("postgres", "postgresql://postgres:root@localhost/books?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("connected to database")
	defer DB.Close()

}
