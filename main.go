package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/base-dev/handler"
	_ "github.com/lib/pq"
)

var db *sql.DB

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

		//dump, err := httputil.DumpRequest(r, true)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//
		//log.Printf("request body: %s", dump)

		next.ServeHTTP(w, r)

	})
}

func main() {
	h := handler.New(&handler.Config{
		Schema: &Schema,
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
