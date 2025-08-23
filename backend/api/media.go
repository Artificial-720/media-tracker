package api

import (
	"fmt"
	"net/http"
)

func allMedia(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ALL MEDIA\n")
}

func getMedia(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "SPECIFIC MEDIA\n")
}
