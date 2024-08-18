package controller

import (
	"encoding/json"
	"go-crud-book-api-postgres/models"
	"log"
	"net/http"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type Response struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []models.Book `json:"data"`
}

func AddNewBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		log.Fatalf("Cannot decode from request body.  %v", err)
	}

	insertID := models.AddNewBook(book)

	res := response{
		ID:      insertID,
		Message: "Add New Book successfully!",
	}

	json.NewEncoder(w).Encode(res)
}
