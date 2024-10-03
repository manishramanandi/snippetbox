package main

import (
	"fmt"
	"html/template"
	"log"      // prints error or any logs
	"net/http" // create a http server
	"strconv"
)

// Define a home handeler function
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Use the template.ParseFiles() function read the template
	// file into a template set If there is an error, we log the details
	// error message and use the http.Error() function to see a generic 500
	// internal server error response to the user.

	ts, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// We then use the Execute() method on the template set to
	// write the template contents as the response body. the last parameter
	// to Execute() represents any dynamic data that we want to pass in, which
	// for now we will leave as nil.

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)

}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", 405)
		return
	}
	w.Write([]byte("Create a new Snippet..."))
}
