package dataBase

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	// Cargar el archivo .env
	if err := godotenv.Load("../.env"); err != nil {
		fmt.Printf("Error cargando archivo .env: %v", err)
	}
}

func GetDBConnection() (*sql.DB, error) {
	connStr := os.Getenv("DB_STRING_CONN")
	return sql.Open("postgres", connStr)
}