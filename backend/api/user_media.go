package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Artificial-720/media-tracker/auth"
	"github.com/Artificial-720/media-tracker/db"
	"github.com/gorilla/mux"
)

func allUserMedia(w http.ResponseWriter, r *http.Request) {
	a := r.Context().Value("auth").(auth.Auth)
	user, err := db.GetUserByUsername(a.Username)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error getting user account")
		return
	}

	items, err := db.GetAllUserMedia(user.ID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeResponse(w, http.StatusOK, items)
}

func getUserMedia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	item, err := db.GetUserMedia(id)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}

	writeResponse(w, http.StatusOK, item)
}

func postUserMedia(w http.ResponseWriter, r *http.Request) {
	a := r.Context().Value("auth").(auth.Auth)
	user, err := db.GetUserByUsername(a.Username)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "error getting user account")
		return
	}

	var item db.UserMedia
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	item.UserID = user.ID

	id, err := db.InsertUserMedia(item)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	createdEntry, err := db.GetUserMedia(id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeResponse(w, http.StatusCreated, createdEntry)
}

func putUserMedia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	var item db.UserMedia
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	item.ID = id
	updatedItem, err := db.UpdateUserMedia(id, item)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeResponse(w, http.StatusOK, updatedItem)
}

func deleteUserMedia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.DeleteUserMedia(id); err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}

	writeResponse(w, http.StatusNoContent, nil)
}
