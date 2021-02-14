package main

import (
	"fmt"
	"log"
	"net/http"
)

// This is a simple example of a net/http package
func main() {
	// Set up two routes to bind two handlers separately
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	// Bind tho port and use DefaultServeMux as the multiplexer
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handler echoes r.URL.Path
func indexHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "URL.Path = %q\n", r.URL.Path)
}

// Handler echoes r.URL.Header
func helloHandler(rw http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Fprintf(rw, "Header[%q] = %q\n", k, v)
	}
}