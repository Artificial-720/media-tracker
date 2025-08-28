package db

import "errors"

type UserMedia struct {
	ID int64 `json:"id"`
	UserID int64 `json:"user_id"`
	MediaID int64 `json:"media_id"`
	Status string `json:"status"`
	Note string `json:"note"`
}

func InsertUserMedia(item UserMedia) (int64, error) {
	result, err := db.Exec("INSERT INTO user_media (user_id, media_id, status, note) VALUES (?, ?, ?, ?)", item.UserID, item.MediaID, item.Status, item.Note)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetUserMedia(id int64) (*UserMedia, error) {
	var item UserMedia

	row := db.QueryRow("SELECT id, user_id, media_id, status, note FROM user_media WHERE id = ?", id)
	err := row.Scan(&item.ID, &item.UserID, &item.MediaID, &item.Status, &item.Note)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func GetAllUserMedia(userID int64) ([]UserMedia, error) {
	rows, err := db.Query("SELECT id, user_id, media_id, status, note FROM user_media WHERE user_id=?", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []UserMedia

	for rows.Next() {
		var item UserMedia
		err := rows.Scan(&item.ID, &item.UserID, &item.MediaID, &item.Status, &item.Note)
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

func UpdateUserMedia(id int64, item UserMedia) (*UserMedia, error) {
	_, err := db.Exec("UPDATE user_media SET user_id=?, media_id=?, status=?, note=? WHERE id=?", item.UserID, item.MediaID, item.Status, item.Note, item.ID)
	if err != nil {
		return nil, err
	}

	updatedItem, err := GetUserMedia(id)
	if err != nil {
		return nil, err
	}
	return updatedItem, nil
}

func DeleteUserMedia(id int64) error {
	result, err := db.Exec("DELETE FROM user_media WHERE id=?", id)
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
