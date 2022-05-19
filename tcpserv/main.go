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

	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
