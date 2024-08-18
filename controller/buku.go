package controller

import (
	"encoding/json"
	"fmt"
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

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	books, err := models.GetAllBooks()

	if err != nil {
		log.Fatalf("Cannot get data. %v", err)
	}

	var response Response
	response.Status = 1
	response.Message = "Success"
	response.Data = books

	json.NewEncoder(w).Encode(response)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Cannot convert string to int.  %v", err)
	}

	var buku models.Book

	err = json.NewDecoder(r.Body).Decode(&buku)

	if err != nil {
		log.Fatalf("Tidak bisa decode request body.  %v", err)
	}

	updatedRows := models.UpdateBook(int64(id), buku)

	msg := fmt.Sprintf("Book has been updated. Update %v rows/record", updatedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}
