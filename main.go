package main

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	"net/http"
	"os"

	_ "modernc.org/sqlite"

	database "aurora/database/gen"
)

//go:embed schema.sql
var ddl string

var queries *database.Queries
var ctx context.Context

func home(w http.ResponseWriter, r *http.Request) {

	var n int
	as, err := queries.ListAuthors(ctx)
	if err != nil {
		n = -1
	}
	n = len(as)

	fmt.Fprintf(w, "Hello, World!, n: %d", n)
}

func main() {
	os.Remove("./database/db_file.db")

	ctx = context.Background()

	db, err := sql.Open("sqlite", "./database/db_file.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// create tables
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		panic(err)
	}

	queries = database.New(db)

	r := NewRouter()

	r.GET("/home", home)

	s := &http.Server{
		Handler: r.ServeMux,
		Addr:    ":8080",
	}

	err = s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"htmx/views"
// )

// type Count struct {
// 	Count int
// }

// func increment(w http.ResponseWriter, r *http.Request) {
// 	// Generate the component with a name
// 	component := views.Hello("Fucking hell this hurt")

// 	// Render the component using Templ's handler
// 	templ.Handler(component).ServeHTTP(w, r)
// }

// func main() {
// 	fmt.Println("Hello, world!")

// 	component := views.Hello("FUcking hell this hurt")
// 	http.Handle("/", templ.Handler(component))

// 	http.HandleFunc("/count", increment)

// 	log.Println("Server running on http://localhost:8080")
// 	http.ListenAndServe(":8080", nil)
// }
