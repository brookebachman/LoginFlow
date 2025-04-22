package main

import (
    "fmt"
    "net/http"
)

func IngestHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Ingest endpoint hit")
}

func SuspiciousHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Suspicious endpoint hit")
}
