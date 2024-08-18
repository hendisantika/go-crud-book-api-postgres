package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-crud-book-api-postgres/models"
	"log"
	"net/http"
	"strconv"
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

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Canot convert from string to int.  %v", err)
	}

	buku, err := models.GetBookByID(int64(id))

	if err != nil {
		log.Fatalf("Cannot get book data. %v", err)
	}

	json.NewEncoder(w).Encode(buku)
}
