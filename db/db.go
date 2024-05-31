package dataBase

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func GetDBConnection() (*sql.DB, error) {
	connStr := os.Getenv("DB_STRING_CONN")
	return sql.Open("postgres", connStr)
}