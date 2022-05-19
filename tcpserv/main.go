package main

import (
	"io"
	"net/http"
	"text/template"
)

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "dog dog doggy")
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "home")
}

func me(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	io.WriteString(w, "Alister")
}

var tmpl *template.Template

func main() {
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))

	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
