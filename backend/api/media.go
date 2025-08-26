package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/Artificial-720/media-tracker/db"
)

func allMedia(w http.ResponseWriter, r *http.Request) {
	items, err := db.GetAllMedia()
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeResponse(w, http.StatusOK, items)
}

func getMedia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	item, err := db.GetMedia(id)
	if err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}

	writeResponse(w, http.StatusOK, item)
}

func postMedia(w http.ResponseWriter, r *http.Request) {
	var item db.MediaItem
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := db.InsertMedia(item)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	item.ID = id
	writeResponse(w, http.StatusCreated, item)
}

func putMedia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	var item db.MediaItem
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	item.ID = id
	updatedItem, err := db.UpdateMedia(id, item)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeResponse(w, http.StatusOK, updatedItem)
}

func deleteMedia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := db.DeleteMedia(id); err != nil {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}

	writeResponse(w, http.StatusNoContent, nil)
}
