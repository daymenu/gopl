package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var defaultDb *sql.DB

func New() (*sql.DB, error) {
	if defaultDb != nil {
		return defaultDb, nil
	}
	db, err := sql.Open("mysql", "root:mysqldb1314@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

func Close() {
	if defaultDb != nil {
		defaultDb.Close()
	}
}
