package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    QueryType,
		Mutation: MutationType,
	})
	if err != nil {
		log.Fatal(err)
	}
	db, err = sql.Open("postgres", "postgres://vagrant:vagrant@localhost:5432/graphql")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error: Could not establish a connection with the database")
	}

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})

	http.Handle("/graphql", h)

	serverAndPort := "127.0.0.1:8080"
	fmt.Printf("Listen on %s", serverAndPort)

	log.Fatal(http.ListenAndServe(serverAndPort, nil))
}
