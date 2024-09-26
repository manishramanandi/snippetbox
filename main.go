package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	// Define a home handler function writes a byte slice containing
	// "Hello from snippetbox" as the response body
	w.Write([]byte("Hello from Snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is Showsnippet....."))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is CreateSnippet..."))
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
