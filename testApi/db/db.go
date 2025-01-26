package db

import (
	"fmt"
	"log"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
)

var DB *pg.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	options := &pg.Options{
		User:     username,
		Password: password,
		Database: dbName,
		Addr:     "localhost:5432",
	}

	DB = pg.Connect(options)

	_, err = DB.Exec("SELECT 1")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	fmt.Println("Successfully connected to the database!")
}
