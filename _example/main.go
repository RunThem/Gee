package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RunThem/gee"
)

// This is a simple example of a gee package
func main() {
	r := gee.New()

	r.GET("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "URL.Path = %q\n", r.URL.Path)
	})

	r.GET("/hello", func(rw http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header{
			fmt.Fprintf(rw, "Header[%q] = %q\n", k, v)
		}
	})

	fmt.Println("Start ...")
	log.Println(r.Run(":8080"))
}