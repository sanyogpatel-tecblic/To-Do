package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sanyogpatel-tecblic/API-Simple/model"
)

var DB *sql.DB

func GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []model.Book
	rows, err := DB.Query("SELECT * FROM books")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	var book model.Book
	row := DB.QueryRow("SELECT * FROM books WHERE id=$1", id)
	err = row.Scan(&book.ID, &book.Title, &book.Author, &book.Description)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Fatal(err)
	}
	_, err = DB.Exec("INSERT INTO books (id,title, author, description) VALUES ($1, $2, $3,$4)", book.ID, book.Title, book.Author, book.Description)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)

	id, err := strconv.Atoi(param["id"])

	if err != nil {
		log.Fatal(err)
	}
	var book model.Book

	json.NewDecoder(r.Body).Decode(&book)

	_, err = DB.Exec("update books set title=$1, author=$2,description=$3 where id=$4", book.Title, book.Author, book.Description, id)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	var book model.Book
	_, err = DB.Exec("delete from books where id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(book)
}
