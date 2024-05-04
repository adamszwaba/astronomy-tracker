package main

import (
	"fmt"
	"net/http"
)

func main() {
	// create a simple server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi, World!")
	})

	http.HandleFunc("/v1/", func(w http.ResponseWriter, r * http.Request) {
		fmt.Fprintf(w, "maybe")
	})

	http.ListenAndServe(":8080", nil)
}
