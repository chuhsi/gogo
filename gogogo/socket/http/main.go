package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello, you've requsted: %s\n", r.URL.Path)
	})
	http.ListenAndServe(":9090", nil)
}