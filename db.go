package main

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

// Initialize the database connection
func InitDB() error {
	var err error
	
	// Get database path from environment variable or use default
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "test.db"
	}

	// Open a connection to the database
	db, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
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

	// Auto migrate the schema
	if err := db.AutoMigrate(&LoginEvent{}); err != nil {
		log.Printf("Error migrating database: %v", err)
		return err
	}

	log.Println("Database connected and migrated successfully")
	return nil
}

// GetDB returns the current database connection
func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("Database connection is nil")
	}
	return db
}
