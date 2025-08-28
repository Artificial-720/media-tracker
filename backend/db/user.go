package db

import "time"

type User struct {
	ID int64
	Username string
	PasswordHash string
	CreatedAt time.Time
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	row := db.QueryRow("SELECT * FROM users WHERE username=?", username)
	err := row.Scan(&user.ID, &user.Username, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
