package main

import (
	"html/template"
	"log"
	"os"
)

var tmpl *template.Template

func main() {
	tmpl := template.Must(template.ParseGlob("./templates/*.tmpl"))

	g1 := struct {
		Score1 int
		Score2 int
	}{
		7,
		9,
	}

	err := tmpl.Execute(os.Stdout, g1)
	if err != nil {
		log.Fatalln(err)
	}
}
