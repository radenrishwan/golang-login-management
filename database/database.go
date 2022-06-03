package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"login-management/helper"
	"os"
)

func GetDatabase() *sql.DB {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := ""

	if os.Getenv("ENVIRONMENT") == "production" {
		database = os.Getenv("DB_DATABASE")
	} else {
		database = os.Getenv("DB_DATABASE_TEST")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)

	db, err := sql.Open("mysql", dsn)
	helper.PanicIfError(err)

	return db
}
