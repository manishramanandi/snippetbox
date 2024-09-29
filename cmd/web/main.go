package main

import (
	"fmt"
	"log" // library for prints error or any logs
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	// Define a home handler function writes a byte slice containing
	// "Hello from snippetbox" as the response body
	w.Write([]byte("Hello from Snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and
	// try to convert it to an integer using the strconv.Atoi() function
	// if it can not be converted to an integer and the value is less than 1,
	// we return a 404 page not found response

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// Use the fmt.Fprintf() function to interpolate the id with our response
	// and write it to the httpResponseWriter
	fmt.Fprint(w, "Display a specific snippet with ID ...", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	// use r.Method to check whether the request is using POST
	// or not. Note that
	// http.MethodPost is a constant equal to the string "POST"
	if r.Method != http.MethodPost {
		// if it is not use, use the w.WriteHeader() method to send a 405
		// status code and the w.Write() method to write "Method not allowed"
		// response bosy. We then return from the function so that the
		// subsequent code is not executed
		// use the Header().Set() Method to add an 'Allow: POST' header
		// to the response header map. The first parameter is the header name,
		// and the second parameter is the header value
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(405)
		// w.Write([]byte("Method not allowed"))
		// Use the http.Error() function to send a 405 status code and
		// "Method Not Allowed" string as the response body.
		http.Error(w, "Method now Allowed", 405)
		return
	}
	w.Write([]byte("Create a new snippet...."))
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux,
	// then register the home function as the handler for the "/"URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters: the TCP network address to listen on :4000 you can change it if you want to
	// we use the log.Fatal() function to log the error message and exit simple
	// Note that any error returned by http.ListenAndServer() is always non-nil.

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
