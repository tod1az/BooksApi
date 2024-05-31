package dataBase

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func GetDBConnection() (*sql.DB, error) {
	connStr := os.Getenv("DB_STRING_CONN")
	fmt.Print(connStr)
	return sql.Open("postgres", connStr)
}