package main

import (
	"fmt"
	"net/http"

	"github.com/nicholasjackson/env"
)

var name = env.String("NAME", false, "Jun", "Hello server")

func main() {

	env.Parse()

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello %s", *name)
	})

	http.ListenAndServe(":9090", nil)
}
