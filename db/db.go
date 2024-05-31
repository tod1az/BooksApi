package dataBase

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func GetDBConnection() (*sql.DB, error) {
	connStr := os.Getenv("DB_STRING_CONN")
	log.Fatal(connStr)
	return sql.Open("postgres", connStr)
}