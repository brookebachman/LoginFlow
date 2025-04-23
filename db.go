package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// Initialize the database connection
func InitDB() error {
	var err error
	// Open a connection to the database
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Printf("Error initializing database: %v", err)
		return err
	}

	// Test the connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Error getting database instance: %v", err)
		return err
	}

	// Ping the database to ensure connection is alive
	if err := sqlDB.Ping(); err != nil {
		log.Printf("Error pinging database: %v", err)
		return err
	}

	log.Println("Database connected successfully")
	return nil
}

// GetDB returns the current database connection
func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("Database connection is nil")
	}
	return db
}
