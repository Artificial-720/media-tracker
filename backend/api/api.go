package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func base(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from base")
}

func RegisterRoutes(router *mux.Router) {
	log.Printf("Registering API routes...\n")
	router.HandleFunc("/", base)

	mediaRouter := router.PathPrefix("/media").Subrouter()
	mediaRouter.HandleFunc("/", allMedia).Methods("GET")
	mediaRouter.HandleFunc("/{id}", getMedia).Methods("GET")

}
