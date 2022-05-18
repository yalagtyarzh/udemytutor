package main

import (
	"log"
	"net/http"
	"text/template"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	tmpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

var tmpl *template.Template

func main() {
	tmpl = template.Must(template.ParseGlob("*.gohtml"))

	var d hotdog
	http.ListenAndServe(":8080", d)
}
