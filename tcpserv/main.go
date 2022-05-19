package main

import (
	"net/http"
	"text/template"
)

func dog(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "template.gohtml", "dog")
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "template.gohtml", "home")
}

func me(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "template.gohtml", "Alister")
}

var tmpl *template.Template

func main() {
	tmpl = template.Must(template.ParseGlob("templates/*.gohtml"))

	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}
