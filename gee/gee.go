package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

// uni handler for all requests
type Engine struct {
	Routers map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{
		Routers: make(map[string]HandlerFunc),
	}
}

func (engine *Engine) AddRoute(method, pattern string, handlerFunc HandlerFunc) {
	key := method + "-" + pattern
	engine.Routers[key] = handlerFunc
}

func (engine *Engine) Get(pattern string, handlerFunc HandlerFunc) {
	engine.AddRoute("GET", pattern, handlerFunc)
}

func (engine *Engine) Post(pattern string, handlerFunc HandlerFunc) {
	engine.AddRoute("POST", pattern, handlerFunc)
}

func (engine *Engine) Run(addr string) {
	http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if hander, ok := engine.Routers[key]; ok {
		hander(rw, req)
	} else {
		fmt.Fprintf(rw, "404 NOT FOUND, %s", req.URL.Path)
	}
}
