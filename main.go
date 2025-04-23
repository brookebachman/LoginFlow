package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
    log.Println("Starting server on :8080")

    // Add CORS support
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:8000"}, // or "*" to allow all origins
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
        AllowedHeaders: []string{"Content-Type"},
    })

    // Apply CORS middleware
    http.HandleFunc("/login-events", CreateLoginEvent)
    http.HandleFunc("/suspicious", SuspiciousHandler)

    handler := c.Handler(http.DefaultServeMux)

    log.Fatal(http.ListenAndServe(":8080", handler))
}

