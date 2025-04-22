package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type LoginEvent struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
    TenantID   string    `json:"tenant_id"`
    Username   string    `json:"username"`
    LoginStatus string   `json:"login_status"`
    Origin     string    `json:"origin"`
    Timestamp  time.Time `json:"timestamp"`
}

var db *gorm.DB


func IngestHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Ingest endpoint hit")
}

func SuspiciousHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Suspicious endpoint hit")
} 
func CreateLoginEvent(w http.ResponseWriter, r *http.Request){
	var event LoginEvent

	// Parse json request body
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	//Validate input fields
	if event.TenantID == "" || event.Username == "" || event.LoginStatus == "" || event.Origin == "" || event.Timestamp.IsZero(){
		http.Error(w, "Event Already Exists", http.StatusConflict)
	}

	//Check if login event already exists
	var existingEvent LoginEvent
	//next line queries the db to see if the event exists, if it does it is then stored in the existing evetn table
    db.Where("tenant_id = ? AND username = ? AND timestamp = ?", event.TenantID, event.Username, event.Timestamp).First(&existingEvent)

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

