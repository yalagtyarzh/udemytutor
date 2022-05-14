package main

import (
	"html/template"
	"log"
	"os"
)

var tmpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func main() {
	tmpl = template.Must(template.ParseGlob("./templates/*.tmpl"))

	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	err := tmpl.ExecuteTemplate(os.Stdout, "tmpl.tmpl", buddha)
	if err != nil {
		log.Fatalln(err)
	}
}
