package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func writeResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func writeError(w http.ResponseWriter, status int, msg string) {
	writeResponse(w, status, map[string]string{"error": msg})
}

func base(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from base")
}

func RegisterRoutes(router *mux.Router) {
	log.Printf("Registering API routes...\n")
	router.HandleFunc("/", base)

	mediaRouter := router.PathPrefix("/media").Subrouter()
	mediaRouter.HandleFunc("", allMedia).Methods("GET")
	mediaRouter.HandleFunc("", postMedia).Methods("POST")
	mediaRouter.HandleFunc("/{id}", getMedia).Methods("GET")
	mediaRouter.HandleFunc("/{id}", putMedia).Methods("PUT")
	mediaRouter.HandleFunc("/{id}", deleteMedia).Methods("DELETE")

	authRouter := router.PathPrefix("/auth").Subrouter()
	authRouter.HandleFunc("/login", postLogin).Methods("POST")
	// authRouter.HandleFunc("/register", postRegister).Methods("POST")

	// userMediaRouter := router.PathPrefix("/user/media").Subrouter()
	// userMediaRouter.HandleFunc("", allUserMedia).Methods("GET")
	// userMediaRouter.HandleFunc("", postUserMedia).Methods("POST")
	// userMediaRouter.HandleFunc("/{id}", getUserMedia).Methods("GET")
	// userMediaRouter.HandleFunc("/{id}", putUserMedia).Methods("PUT")
	// userMediaRouter.HandleFunc("/{id}", deleteUserMedia).Methods("DELETE")
}
