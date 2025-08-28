package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Username string
	Valid bool
}

var secret = []byte("this-is-my-key")

func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": username,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) Auth {
	var a Auth
	a.Valid = false

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return secret, nil
	})
	if err == nil && token.Valid {
		sub, err := token.Claims.GetSubject()
		if err == nil {
			a.Username = sub
			a.Valid = true
		}
	}

	return a
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
