package router

import (
	"github.com/gorilla/mux"
	"go-crud-book-api-postgres/controller"
)

func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/books", controller.AddNewBook).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/books", controller.GetAllBooks).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/books/{id}", controller.GetBookByID).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/books/{id}", controller.UpdateBook).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/books/{id}", controller.DeleteBook).Methods("DELETE", "OPTIONS")

	return router
}
