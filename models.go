package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func init() {
    var err error
    db, err = sql.Open("sqlite3", "./db/database.db")
    if err != nil {
        log.Fatal(err)
    }

    // Optional: create table if not exists
    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS logins (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            tenant_id TEXT,
            user TEXT,
            status TEXT,
            origin TEXT,
            timestamp TEXT
        );
    `)
    if err != nil {
        log.Fatal(err)
    }
}
