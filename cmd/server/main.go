package main

import (
	"log"
	"net/http"

	"github.com/colebowl/coles-stream/internal/db"
	"github.com/colebowl/coles-stream/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize database
	if err := db.Init(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	r := mux.NewRouter()

	// Set up routes
	r.HandleFunc("/", handlers.StreamHandler).Methods("GET")
	r.HandleFunc("/post/new", handlers.NewPostHandler).Methods("GET", "POST")
	r.HandleFunc("/post/{id}/edit", handlers.EditPostHandler).Methods("GET", "POST")
	r.HandleFunc("/auth", handlers.AuthHandler).Methods("GET", "POST")

	// Serve static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
