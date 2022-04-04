// Members: Miguel Avila, Federico Rosado

package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	_ "github.com/lib/pq" //Third party package
	"mavila_frosado.net/test1/pkg/models/postgresql"
)

// Database Function
func setUpDB(dsn string) (*sql.DB, error) {

	//Establish connection to the dababase
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	//test connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

//Dependencies Injection (passing)
type application struct {
	Blogs *postgresql.BlogModel
	addr  string
}

func main() {
	// create a command line flag for the port
	addr := flag.String("addr", ":8080", "HTTP network address")
	dsn := flag.String("dsn", "postgres://dbtest:dbtest@localhost/dbtest?sslmode=disable", "PostgreSQL DSN (Data Source Name)")

	flag.Parse()
	//create and initialize
	var db, err = setUpDB(*dsn)

	//Check if errors
	if err != nil {
		log.Fatal(err)
	}

	//A must before exiting
	defer db.Close()
	app := &application{
		Blogs: &postgresql.BlogModel{DB: db},
		addr:  *addr,
	}

	//Create a custom server
	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}

	log.Printf("Starting Server on port %s", *addr)
	err = srv.ListenAndServe()
	log.Fatal(err)

}
