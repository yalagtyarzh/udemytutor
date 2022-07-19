package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func main() {
	tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))

	http.HandleFunc("/", dogs)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080", nil)
}

func dogs(w http.ResponseWriter, r *http.Request) {
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("template didn't execute:", err)
	}
}
