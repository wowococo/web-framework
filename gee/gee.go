package gee

import (
	"fmt"
	"net/http"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine implements the Handler interface
//
//	type Handler interface {
//		ServeHTTP(ResponseWriter, *Request)
//	}
type Engine struct {
	routers map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{
		routers: make(map[string]HandlerFunc),
	}
}

func (engine *Engine) addRoute(method, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.routers[key] = handler
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

// parser the request path, look for the routers map, go to handler or response 404
func (engine *Engine) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.routers[key]; ok {
		handler(rw, req)
	} else {
		fmt.Fprintf(rw, "404 NOT FOUND: %s\n", req.URL)
	}
}
