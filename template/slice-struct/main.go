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

	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	mlk := sage{
		Name:  "Martin Luther King",
		Motto: "Hatred never ceases with hatred but with love alone is healed",
	}

	jesus := sage{
		Name:  "Jesus",
		Motto: "Love all",
	}

	muhammad := sage{
		Name:  "Muhammad",
		Motto: "To overcome evil with good is good, to resist evil by evil is evil",
	}

	sages := []sage{buddha, gandhi, mlk, jesus, muhammad}

	err := tmpl.ExecuteTemplate(os.Stdout, "tmpl.tmpl", sages)
	if err != nil {
		log.Fatalln(err)
	}
}
