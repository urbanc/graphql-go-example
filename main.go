package main

import (
	"database/sql"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	_ "github.com/lib/pq"
)

var db *sql.DB

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("request body: %s", dump)

		next.ServeHTTP(w, r)

	})
}

func main() {
	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    QueryType,
		Mutation: MutationType,
	})
	if err != nil {
		log.Fatal(err)
	}
	db, err = sql.Open("postgres", "postgres://vagrant:vagrant@localhost:5432/graphql?sslmode=disable")
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

	http.Handle("/graphql", logMiddleware((h)))

	// serve a graphiql IDE
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	serverAndPort := "127.0.0.1:8080"
	log.Printf("Listen on %s\n", serverAndPort)

	log.Fatal(http.ListenAndServe(serverAndPort, nil))
}
