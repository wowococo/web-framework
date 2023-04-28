package main

import (
	"fmt"
	"net/http"

	"gee"
)

func main() {
	engine := gee.New()

	engine.GET("/", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(rw, "req url = %q", req.URL.Path)
	})
	engine.GET("/hello", func(rw http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(rw, "Headers[%q] = %q", k, v)
		}
	})

	engine.Run(":9999")

}
