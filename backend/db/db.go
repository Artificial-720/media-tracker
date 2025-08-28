package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Open(path string) error {
	database, err := sql.Open("sqlite3", path)
	if err == nil {
		db = database
	}
	return err
}

func Close() {
	db.Close()
}
