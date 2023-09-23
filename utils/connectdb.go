package util

import (
	"database/sql"
	"fmt"

	//"time"
	//"encoding/csv"

	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type BookDb struct{}

// Connect to DB
func (bookDb *BookDb) ConnectDB() (*sql.DB, error) {

	if err := godotenv.Load(); err != nil {

		panic(err.Error())
	}

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db, nil

}
