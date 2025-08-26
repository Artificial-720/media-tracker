package api

import (
	"encoding/json"
	"net/http"

	"github.com/Artificial-720/media-tracker/auth"
	"github.com/Artificial-720/media-tracker/db"
)

type credentials struct {
	Username string
	Password string
}

func postLogin(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var creds credentials
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := db.GetUserByUsername(creds.Username)
	if err != nil  || !auth.CheckPasswordHash(user.PasswordHash, creds.Password){
		writeError(w, http.StatusUnauthorized, "invalid username or password")
		return
	}

	token, err := auth.GenerateJWT(user.Username)
	writeResponse(w, http.StatusOK, map[string]string{"token": token})
}

