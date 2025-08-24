package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/Artificial-720/media-tracker/db"
)

func allMedia(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ALL MEDIA\n")

	items, err := db.GetAllMedia()
	if err != nil {
		fmt.Fprintf(w, "Error %s\n", err)
	}

	for _, elm := range items {
		fmt.Fprintf(w, "%v\n", elm)
	}
}

func getMedia(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Error %s\n", err)
		return
	}

	fmt.Fprintf(w, "SPECIFIC MEDIA with id = %d\n", id)

	item, err := db.GetMedia(id)
	if err != nil {
		fmt.Fprintf(w, "Error %s\n", err)
		return
	}

	fmt.Fprintf(w, "%v\n", item)
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
