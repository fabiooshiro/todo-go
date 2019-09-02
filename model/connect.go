package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

var con *sql.DB

func Connect() *sql.DB {
	dbConfig := os.Getenv("DATABASE_CONFIG")
	if dbConfig == "" {
		dbConfig = "user:password@tcp(localhost:3306)/db"
	}
	db, err := sql.Open("mysql", dbConfig)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database")
	con = db
	return db
}
