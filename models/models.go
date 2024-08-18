package models

import (
	"fmt"
	"go-crud-book-api-postgres/config"
	"log"
)

type Book struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishedAt string `json:"published_at"`
}

func AddNewBook(book Book) int64 {
	db := config.CreateConnection()

	defer db.Close()

	sqlStatement := `INSERT INTO buku (title, author, published_at) VALUES ($1, $2, $3) RETURNING id`

	var id int64

	err := db.QueryRow(sqlStatement, book.Title, book.Author, book.PublishedAt).Scan(&id)

	if err != nil {
		log.Fatalf("Cannot execute query. %v", err)
	}

	fmt.Printf("Insert single record %v", id)

	return id
}
