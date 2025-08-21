package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting Server")

	router := mux.NewRouter()

	// Frontend static files
	staticPath := filepath.Join("..", "frontend")
	fileServer := http.FileServer(http.Dir(staticPath))
	router.PathPrefix("/").Handler(fileServer)

	http.ListenAndServe(":8080", router)
}
