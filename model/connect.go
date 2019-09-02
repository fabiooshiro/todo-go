package model

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/db")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database")
	con = db
	return db
}
