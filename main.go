package main

import (
	"log"
	"net/http"
)

func main() {
    log.Println("Starting server on :8080")

    http.HandleFunc("/login-events", CreateLoginEvent)
    http.HandleFunc("/suspicious", SuspiciousHandler)

    log.Fatal(http.ListenAndServe(":8080", nil))
}

