package models

import (
	"database/sql"
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

func GetAllBooks() ([]Book, error) {
	db := config.CreateConnection()

	defer db.Close()

	var books []Book

	sqlStatement := `SELECT * FROM book`

	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Cannot execute query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var book Book

		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublishedAt)

		if err != nil {
			log.Fatalf("Cannot get data. %v", err)
		}

		books = append(books, book)

	}

	return books, err
}

func GetBookByID(id int64) (Book, error) {
	db := config.CreateConnection()

	defer db.Close()

	var book Book

	sqlStatement := `SELECT * FROM book WHERE id=$1`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.PublishedAt)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Data not found!")
		return book, nil
	case nil:
		return book, nil
	default:
		log.Fatalf("Can not get data. %v", err)
	}

	return book, err
}

func UpdateBook(id int64, book Book) int64 {
	db := config.CreateConnection()

	defer db.Close()

	sqlStatement := `UPDATE book SET title=$2, author=$3, publishedat=$4 WHERE id=$1`

	res, err := db.Exec(sqlStatement, id, book.Title, book.Author, book.PublishedAt)

	if err != nil {
		log.Fatalf("Cannot execute query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error when check rows/data. %v", err)
	}

	fmt.Printf("Total updated rows/record  %v\n", rowsAffected)

	return rowsAffected
}
