package main

import (
	"database/sql"
	"encoding/gob"
	"log"
	"net/http"

	"github.com/sanyogpatel-tecblic/To-Do/config"
	"github.com/sanyogpatel-tecblic/To-Do/model"
)

var app config.AppConfig
var db *sql.DB

// main is the main function
func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Fatal(http.ListenAndServe(":8000", nil))

}

func run() (*sql.DB, error) {
	var err error
	// what am I going to put in the session
	gob.Register(model.Book{})

	//connect to database
	log.Println("connecting to database...")
	db, err = sql.Open("postgres", "postgresql://postgres:root@localhost/books?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database!")
	return db, nil
}
