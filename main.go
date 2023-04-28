package main

import (
	"fmt"
	"net/http"

	"gee"
)

func main() {
	engine := gee.New()

	engine.Get("/", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(rw, "req url = %q", req.URL.Path)
	})
	engine.Get("/hello", func(rw http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(rw, "Headers[%q] = %q", k, v)
		}
	})

	engine.Run(":9999")

}
