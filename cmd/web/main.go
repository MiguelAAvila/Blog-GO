// Members: Miguel Avila, Federico Rosado

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq" //Third party package
	"mavila_frosado.net/test1/pkg/models/postgresql"
)

// Database Function
func setUpDB() (*sql.DB, error) {
	const (
		host     = "localhost"
		port     = 5432
		user     = "dbtest"
		password = "dbtest"
		dbname   = "dbtest"
	)

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

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
}

func main() {
	//create and initialize
	var db, err = setUpDB()

	//Check if errors
	if err != nil {
		log.Fatal(err)
	}

	//A must before exiting
	defer db.Close()
	app := &application{
		Blogs: &postgresql.BlogModel{DB: db},
	}

	//Create a custom server
	srv := &http.Server{
		Addr:    ":4000",
		Handler: app.routes(),
	}

	log.Println("Starting Server on port :4000")
	err = srv.ListenAndServe()
	log.Fatal(err)

}
