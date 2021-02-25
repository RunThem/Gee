/**
 * Copyright 2021 RunThem. All Rights Reserved.
 *
 * Distributed under MIT license.
 *
 * See file LICENSE for detail or copy at https://opensource.org/licenses/MIT
 *
 * Author: RunThem
 * Email: iccy.fun@outlook.com
 * Created at 2021/2/25 18:09
 */

package gee

import (
	"log"
	"net/http"
)

// The router is a route implemented by Map and supports only static routes
type router struct {
	handler map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handler: make(map[string]HandlerFunc)}
}

// Add routing rules
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handler[key] = handler
}

// Provide HandlerFunc or 404 Not Found for the accessed path
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handler[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 Not Found: %s\n", c.Path)
	}
}
