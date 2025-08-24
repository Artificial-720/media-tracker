package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Artificial-720/media-tracker/database"
)

func allMedia(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ALL MEDIA\n")

	items, err := database.GetAllMedia()
	if err != nil {
		fmt.Fprintf(w, "Error %s\n", err)
	}

	for _, elm := range items {
		fmt.Fprintf(w, "%s\n", elm)
	}
}

func getMedia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "SPECIFIC MEDIA with id = %s\n", id)
}

func postMedia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "SPECIFIC MEDIA with id = %s\n", id)
}

func putMedia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "SPECIFIC MEDIA with id = %s\n", id)
}

func deleteMedia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "SPECIFIC MEDIA with id = %s\n", id)
}
