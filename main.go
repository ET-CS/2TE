package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// indexTemplate is the HTML template we use to present the index page.
var (
	indexTemplate   = template.Must(template.ParseFiles("index.min.html"))
)

// index sets up a session for the current user and serves the index page
func index(w http.ResponseWriter, r *http.Request) *appError {
	// Fill in the missing fields in index.html
	var data = struct {
		Title, ApplicationName string
	}{"MyTitle", "2TE Example"}

	// Render and serve the HTML
	err := indexTemplate.Execute(w, data)
	if err != nil {
		log.Println("error rendering template:", err)
		return &appError{err, "Error rendering template", 500}
	}
	return nil
}

// appHandler is to be used in error handling
type appHandler func(http.ResponseWriter, *http.Request) *appError

type appError struct {
	Err     error
	Message string
	Code    int
}

// serveHTTP formats and passes up an error
func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		log.Println(e.Err)
		http.Error(w, e.Message, e.Code)
	}
}

// Here everything starts
func main() {
	// Serve the index.html page
	http.Handle("/", appHandler(index))

	// Starting server
	port := "8080"
	fmt.Println("Starting server on port: " + port + "...")
	err := http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
