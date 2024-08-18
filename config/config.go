package config

import (
	"github.com/joho/godotenv" // package used to read the .env file
	_ "github.com/lib/pq"      // postgres golang driver
	"go-crud-book-api-postgres/db/postgres"
	"log"
)

func Initialize() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	postgres.CreateConnection()
}
