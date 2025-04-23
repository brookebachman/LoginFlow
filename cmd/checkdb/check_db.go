package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type LoginEvent struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	TenantID   string    `json:"tenant_id" gorm:"not null"`
	Username   string    `json:"username" gorm:"not null"`
	LoginStatus string   `json:"login_status" gorm:"not null"`
	Origin     string    `json:"origin" gorm:"not null"`
	Timestamp  time.Time `json:"timestamp" gorm:"not null"`
}

func main() {
	// Get database path from environment variable or use default
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "test.db"
	}

	// Open database connection
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Count total rows
	var count int64
	if err := db.Table("events").Count(&count).Error; err != nil {
		log.Fatalf("Failed to count rows: %v", err)
	}

	fmt.Printf("Total rows in events table: %d\n", count)

	// Get all events
	var events []LoginEvent
	if err := db.Find(&events).Error; err != nil {
		log.Fatalf("Failed to fetch events: %v", err)
	}

	// Print each event
	fmt.Println("\nEvents in database:")
	for _, event := range events {
		fmt.Printf("ID: %d, Tenant: %s, User: %s, Status: %s, Origin: %s, Time: %s\n",
			event.ID, event.TenantID, event.Username, event.LoginStatus, event.Origin, event.Timestamp)
	}
} 