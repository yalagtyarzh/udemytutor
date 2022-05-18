package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "dog dog doggy")
}

func cat(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "cat cat cat")
}

func main() {

	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/cat/", cat)

	http.ListenAndServe(":8080", nil)
}
