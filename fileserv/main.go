package main

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

func main() {
	tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))

	http.HandleFunc("/", index)
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.gohtml", nil)
}
