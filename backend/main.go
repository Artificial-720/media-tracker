package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/Artificial-720/media-tracker/api"
	"github.com/Artificial-720/media-tracker/auth"
	"github.com/Artificial-720/media-tracker/db"
	"github.com/Artificial-720/media-tracker/middleware"
)

func main() {
	log.Println("Starting Server")

	err := db.Open("./media.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Loading .env")
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	secretKey := os.Getenv("SECRET_KEY")
	auth.InitAuth(secretKey)

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
