package main

import "net/http"

func main() {

	h := http.FileServer(http.Dir("assets/"))

	http.Handle("/static/", http.StripPrefix("/static/",h))
}