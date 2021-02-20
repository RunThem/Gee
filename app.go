package gee

import (
	"fmt"
	"net/http"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router map[string]HandlerFunc
}

// New is the constructor of gee.Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(rw, r)
	} else {
		fmt.Fprintf(rw, "404 Not Found: %s\n", r.URL)
	}
}
