package util

import (
	"database/sql"
	"fmt"
	"log"

	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Db struct{}

// Connect to DB
func (bookDb *Db) ConnectDB() (*sql.DB, error) {

	envPath := os.Getenv("API_KEY")

	log.Println("ENV_PATH", envPath)
	if envPath == "" {
		if errr := godotenv.Load("/home/aamir/Desktop/My_code/searchRecommend/.env"); errr != nil {

			panic(errr.Error())
		}
	}
	connStr := fmt.Sprintf("host=postgres user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	if err1 := db.Ping(); err1 != nil {
		panic(err1.Error())
	}

	return db, nil

}
