package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Alister-key", "this is from Alister")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any code I want in this func</h1>")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
