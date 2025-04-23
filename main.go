package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
    // Initialize the database connection
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	log.Println("Database connected successfully")

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

