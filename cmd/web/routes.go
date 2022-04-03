package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/createblog", app.createBlogForm)
	mux.HandleFunc("/blog-add", app.createBlog)
	mux.HandleFunc("/blogs", app.blogs)
	//create a fileserver to serve static files
	fs := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	return mux
}
