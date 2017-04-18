package main

import (
	"database/sql"
	"log"
)

func InitDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://vagrant:vagrant@localhost:5432/graphql?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error: Could not establish a connection with the database")
	}

	return db
}
