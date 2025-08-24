package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type MediaItem struct {
	ID int
	Title string
	Type string
	Source string
	ExternalId int
}


func InsertMedia() {}
func GetMedia() {}

func GetAllMedia() ([]MediaItem, error) {
	db, err := sql.Open("sqlite3", "./media.db")
	if err != nil {
		return nil, err
	}
	// this is a new keyword in GO 'defer' this will run
	// a function call at the end of this function call
	defer db.Close()

	rows, err := db.Query("SELECT id, title, type, source FROM media_items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []MediaItem

	for rows.Next() {
		var item MediaItem
		// todo handle external id
		err := rows.Scan(&item.ID, &item.Title, &item.Type, &item.Source)
		if err != nil {
			return items, err
		}
		items = append(items, item)
	}
	err = rows.Err()
	if err != nil {
		return items, err
	}
	return items, nil
}

func UpdateMedia() {}
func DeleteMedia() {}
