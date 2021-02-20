package main

import (
	"fmt"
	"log"
	"net/http"
)

// This is a simple example of a net/http package
func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":8080", engine))
}

// Implementing the Handler interface
type Engine struct{}

// The method required for Handler
func (e *Engine) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(rw, "URL.Path = %q\n", r.URL.Path)
	case "/hello":
		for k, v := range r.Header {
			fmt.Fprintf(rw, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(rw, "404 not found: %s\n", r.URL)
	}
}
