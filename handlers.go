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
	if err := GetDB().Where("tenant_id = ? AND login_status = ? AND timestamp >= ?", tenantID, "failure", now.Add(-timeWindow)).Find(&events).Error; err != nil {
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
	if GetDB() == nil {
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
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	//Check if login event already exists within a 5-second window
	var existingEvent LoginEvent
	log.Printf("Checking for duplicate event with tenant_id=%s, username=%s, timestamp=%v", 
		event.TenantID, event.Username, event.Timestamp)
	
	// Create a 5-second window around the event timestamp
	timeWindow := 5 * time.Second
	startTime := event.Timestamp.Add(-timeWindow)
	endTime := event.Timestamp.Add(timeWindow)
	
	result := GetDB().Where("tenant_id = ? AND username = ? AND timestamp BETWEEN ? AND ?", 
		event.TenantID, event.Username, startTime, endTime).First(&existingEvent)
	
	if result.Error == nil {
		// If no error, record was found (duplicate exists)
		log.Printf("Duplicate event found: %+v", existingEvent)
		http.Error(w, "Event already exists", http.StatusConflict)
		return
	} else if result.Error != gorm.ErrRecordNotFound {
		// If error is not "record not found", it's a database error
		log.Printf("Error querying the database: %v", result.Error)
		http.Error(w, "Database error while checking existing events", http.StatusInternalServerError)
		return
	} else {
		log.Printf("No duplicate found, proceeding to create new event")
	}

	// If we get here, no duplicate was found, create new event
	if err := GetDB().Create(&event).Error; err != nil {
		log.Printf("Error creating event: %v", err)
		http.Error(w, "Failed to store event", http.StatusInternalServerError)
		return
	}

	//If successful this sends a 201 to client saying it was created
	w.WriteHeader(http.StatusCreated)
	//this converts the event object into json
	json.NewEncoder(w).Encode(event)
}



