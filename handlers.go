package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"gorm.io/gorm"
)

// SuspiciousHandler handles GET requests for suspcious events
func SuspiciousHandler(w http.ResponseWriter, r *http.Request) {
	// Query params for filtering (e.g., tenant_id, threshold, time window)
	tenantID := r.URL.Query().Get("tenant_id")
	if tenantID == "" {
		http.Error(w, "Tenant ID is required", http.StatusBadRequest)
		return
	}

	threshold := 5           // 5 failed attempts
	timeWindow := 10 * time.Minute // 10 minutes window
	now := time.Now()

	// Query for failed login events within the time window
	var events []LoginEvent
	if err := db.Where("tenant_id = ? AND login_status = ? AND timestamp >= ?", tenantID, "failure", now.Add(-timeWindow)).Find(&events).Error; err != nil {
		http.Error(w, "Database error while fetching events", http.StatusInternalServerError)
		return
	}

	// Group by origin and count failed attempts
	originFailedCounts := make(map[string]int)
	for _, event := range events {
		originFailedCounts[event.Origin]++
	}

	// Collect origins with more than the threshold of failures
	var suspiciousOrigins []string
	for origin, count := range originFailedCounts {
		if count >= threshold {
			suspiciousOrigins = append(suspiciousOrigins, origin)
		}
	}

	// Send the suspicious origins as response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(suspiciousOrigins)
}

func CreateLoginEvent(w http.ResponseWriter, r *http.Request) {
	if db == nil {
		log.Println("Database connection is nil")
		http.Error(w, "Database not initialized", http.StatusInternalServerError)
		return
	}

	var event LoginEvent

	// Parse json request body
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	log.Printf("Incoming event: %+v", event)

	// Check if timestamp is correctly parsed
	log.Printf("Timestamp: %v\n", event.Timestamp)
	
	//Validate input fields
	if event.TenantID == "" || event.Username == "" || event.LoginStatus == "" || event.Origin == "" || event.Timestamp.IsZero() {
		http.Error(w, "Invalid event data", http.StatusBadRequest)
	}

	//Check if login event already exists
	var existingEvent LoginEvent
	//next line queries the db to see if the event exists, if it does it is then stored in the existing evetn table
	if err := db.Where("tenant_id = ? AND username = ? AND timestamp = ?", event.TenantID, event.Username, event.Timestamp).First(&existingEvent).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("Event does not exist, proceeding to create new event.")
		} else {
			log.Println("Error querying the database:", err)
			http.Error(w, "Database error while checking existing events", http.StatusInternalServerError)
			return
		}
	}

	if existingEvent.ID != 0 {
		http.Error(w, "Event already exists", http.StatusConflict)
		return
	}

	// If the event is not found, meaning its a new login event, save the new event into the database
	if err := db.Create(&event).Error; err != nil {
		//if it cannot insert the row into the db it sends a 500 error to the cliet
		http.Error(w, "Failed to store event", http.StatusInternalServerError)
		return
	}

	//If successful this sends a 201 to client saying it was created
	w.WriteHeader(http.StatusCreated)
	//this converts the event object into json
	json.NewEncoder(w).Encode(event)
}



