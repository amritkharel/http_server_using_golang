package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() *sql.DB {
	dbinfo := "user=postgres password=yourpasswordhere username=yourusernamehere host=localhost dbname=todoapp sslmode=disable"
	var err error
	db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatalf("Error opening database: %s\n", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %s\n", err)
	}

	fmt.Println("Successfully connected!")
	return db
}
