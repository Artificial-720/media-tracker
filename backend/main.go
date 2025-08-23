package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"

	"github.com/Artificial-720/media-tracker/api"
	"github.com/Artificial-720/media-tracker/middleware"
)

func main() {
	log.Println("Starting Server")

	router := mux.NewRouter()

	// Setup API
	apiRouter := router.PathPrefix("/api").Subrouter()
	api.RegisterRoutes(apiRouter)

	// Frontend static files
	staticPath := filepath.Join("..", "frontend")
	fileServer := http.FileServer(http.Dir(staticPath))
	router.PathPrefix("/").Handler(fileServer)

	// Setup Middleware
	router.Use(middleware.LoggingMiddleware)

	http.ListenAndServe(":8080", router)
}
