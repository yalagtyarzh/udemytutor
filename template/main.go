package main

import (
	"html/template"
	"log"
	"os"
)

var tmpl *template.Template

func main() {
	tmpl = template.Must(template.ParseGlob("templates/*.tmpl"))
	err := tmpl.ExecuteTemplate(os.Stdout, "tmpl.tmpl", "xd")
	if err != nil {
		log.Fatalln(err)
	}
}
