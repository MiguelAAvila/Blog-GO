// Members: Miguel Avila, Federico Rosado

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq" //Third party package
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
	db *sql.DB
}

type Blog struct {
	ID           int
	FirstName    string
	LastName     string
	Email        string
	Subject      string
	Message      string
	Date_Created time.Time
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
		db: db,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/createblog", app.createBlogForm)
	mux.HandleFunc("/blog-add", app.createBlog)
	mux.HandleFunc("/blogs", app.blogs)
	log.Println("Starting Server on port :4000")
	err = http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
