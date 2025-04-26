package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var (
	HOST     = os.Getenv("DATABASE_HOST")
	PORT     = os.Getenv("DATABASE_PORT")
	USER     = os.Getenv("DATABASE_USER")
	PASSWORD = os.Getenv("DATABASE_PASSWORD")
	DB_NAME  = os.Getenv("DATABASE_NAME")
)

func ConnectDB() (*sql.DB, error) {

	var psqlInfo string = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, USER, PASSWORD, DB_NAME,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db, nil
}
