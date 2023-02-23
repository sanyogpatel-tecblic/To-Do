package model

// import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}
