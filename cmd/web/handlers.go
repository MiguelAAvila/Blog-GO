// Members: Miguel Avila, Federico Rosado

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"unicode/utf8"
)

//Home Page
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/index.tmpl")
	port := app.addr

	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	data := &templateData{
		Port: port,
	}
	err = ts.Execute(w, data)
	if err != nil {
		log.Panicln(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}
}

//Displays SingUp Form
func (app *application) createBlogForm(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/html/form.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Panicln(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}

}

//blog Page
func (app *application) blogs(w http.ResponseWriter, r *http.Request) {
	blogs, err := app.Blogs.Read()
	port := app.addr

	if err != nil {
		log.Panicln(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
	}
	//instance of templateData
	data := &templateData{
		Blogs: blogs,
		Port:  port,
	}

	//Body part of tmpl
	ts, err := template.ParseFiles("./ui/html/blog.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, data)

	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}

}

//Extract, Validate and Write to the blogs table
func (app *application) createBlog(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	err := r.ParseForm()
	if err != nil {
		http.Error(w,
			http.StatusText(http.StatusBadRequest),
			http.StatusBadRequest)
		return
	}
	firstname := r.PostForm.Get("firstname")
	lastname := r.PostForm.Get("lastname")
	email := r.PostForm.Get("email")
	subject := r.PostForm.Get("subject")
	message := r.PostForm.Get("message")

	//Validate Form Fields
	errors := make(map[string]string)
	//Check each Fields
	if strings.TrimSpace(firstname) == "" {
		errors["firstname"] = "This field cannot be left blank"
	} else if utf8.RuneCountInString(firstname) > 20 {
		errors["firstname"] = "No more than 20 characters"
	}
	//Check each Fields
	if strings.TrimSpace(lastname) == "" {
		errors["lastname"] = "This field cannot be left blank"
	} else if utf8.RuneCountInString(lastname) > 20 {
		errors["lastname"] = "No more than 20 characters"
	}
	if strings.TrimSpace(email) == "" {
		errors["email"] = "This field cannot be left blank"
	} else if utf8.RuneCountInString(email) > 25 {
		errors["email"] = "No more than 25 characters"
	}
	if strings.TrimSpace(subject) == "" {
		errors["subject"] = "This field cannot be left blank"
	} else if utf8.RuneCountInString(subject) > 50 {
		errors["subject"] = "No more than 50 characters"
	}
	if strings.TrimSpace(message) == "" {
		errors["message"] = "This field cannot be left blank"
	} else if utf8.RuneCountInString(subject) > 500 {
		errors["message"] = "No more than 500 characters"
	}
	//Check if errors in the map
	if len(errors) > 0 {
		ts, err := template.ParseFiles("./ui/html/form.page.tmpl")
		err = ts.Execute(w, &templateData{
			ErrorsFromForm: errors,
			FormData:       r.PostForm,
		})
		if err != nil {
			log.Println(err.Error())
			http.Error(w,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			return
		}

		if err != nil {
			log.Panicln(err.Error())
			http.Error(w,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
			return
		}
		return
	}

	//inser a blog
	id, err := app.Blogs.Insert(firstname, lastname, email, subject, message)
	if err != nil {
		log.Println(err.Error())
		http.Error(w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Row with id %d has been inserted.", id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}
