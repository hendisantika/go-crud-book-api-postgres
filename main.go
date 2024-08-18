package main

import (
	"fmt"
	"go-crud-book-api-postgres/config"
	"go-crud-book-api-postgres/router"
	"log"
	"net/http"
)

func main() {
	// initialize config
	config.Initialize()

	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Server running @ port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
