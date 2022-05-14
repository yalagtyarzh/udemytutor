package main

import (
	"html/template"
	"log"
	"os"
)

var tmpl *template.Template

func main() {
	tmpl = template.Must(template.ParseGlob("./templates/*.tmpl"))

	sages := map[string]string{
		"India":    "Gandhi",
		"America":  "MLK",
		"Meditate": "Buddha",
		"Love":     "Jesus",
		"Prophet":  "Muhammad",
	}

	err := tmpl.ExecuteTemplate(os.Stdout, "tmpl.tmpl", sages)
	if err != nil {
		log.Fatalln(err)
	}
}
