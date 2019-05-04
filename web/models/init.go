package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var defaultDb *sql.DB

func New() (*sql.DB, error) {
	fmt.Println(defaultDb)
	if defaultDb != nil {
		return defaultDb, nil
	}
	var err error
	defaultDb, err = sql.Open("mysql", "go:go@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	return defaultDb, err
}

func Close() {
	if defaultDb != nil {
		defaultDb.Close()
	}
}
