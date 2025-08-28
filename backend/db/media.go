package db

import "errors"

type MediaItem struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Type string `json:"type"`
	Source string `json:"source"`
	ExternalId int `json:"external_id"`
}

func InsertMedia(item MediaItem) (int64, error) {
	result, err := db.Exec("INSERT INTO media_items (title, type, source) VALUES (?, ?, ?)", item.Title, item.Type, item.Source)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetMedia(id int64) (*MediaItem, error) {
	var item MediaItem

	row := db.QueryRow("SELECT id, title, type, source FROM media_items WHERE id = ?", id)
	err := row.Scan(&item.ID, &item.Title, &item.Type, &item.Source)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func GetAllMedia() ([]MediaItem, error) {
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

func UpdateMedia(id int64, item MediaItem) (*MediaItem, error) {
	_, err := db.Exec("UPDATE media_items SET title=?, type=?, source=? WHERE id=?", item.Title, item.Type, item.Source, id)
	if err != nil {
		return nil, err
	}

	updatedItem, err := GetMedia(id)
	if err != nil {
		return nil, err
	}
	return updatedItem, nil
}

func DeleteMedia(id int64) error {
	result, err := db.Exec("DELETE FROM media_items WHERE id=?", id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected != 1 {
		return errors.New("No rows deleted")
	}

	return nil
}
