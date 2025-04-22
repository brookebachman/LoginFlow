package main

import (
	"time"

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

// This function will be used to initialize the database and automatically migrate
// the schema (creating tables, etc.) when the application starts.
func MigrateDB(db *gorm.DB) {
	// Automate migration of the models
	err := db.AutoMigrate(&LoginEvent{})
	if err != nil {
		panic("Failed to migrate the database")
	}
}
