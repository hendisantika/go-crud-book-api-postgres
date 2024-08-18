package main

import (
	"fmt"
	"go-crud-book-api-postgres/router"
	"log"
	"net/http"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	r := router.Router()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)
	fmt.Println("Server running @ port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
