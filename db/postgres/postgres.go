package postgres

import (
	"database/sql"
	"fmt"
	"os"
)

func CreateConnection() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to DB!")

	return db
}
