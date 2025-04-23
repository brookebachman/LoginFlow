package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	// Initialize the database
	if err := InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	log.Println("Starting server on :8080")

	// Add CORS support
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type"},
	})

	// Apply CORS middleware
	http.HandleFunc("/login-events", CreateLoginEvent)
	http.HandleFunc("/suspicious", SuspiciousHandler)

	handler := c.Handler(http.DefaultServeMux)

	log.Fatal(http.ListenAndServe(":8080", handler))
}

