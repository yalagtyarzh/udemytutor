package main

import (
	"html/template"
	"log"
	"os"
)

type user struct {
	Name  string
	Motto string
	Admin bool
}

var tmpl *template.Template

func main() {
	tmpl := template.Must(template.ParseGlob("./templates/*.tmpl"))

	u1 := user{
		Name:  "Buddha",
		Motto: "Be the change",
		Admin: false,
	}

	u2 := user{
		Name:  "Gandhi",
		Motto: "Be the change",
		Admin: true,
	}

	u3 := user{
		Name:  "",
		Motto: "Nobody",
		Admin: true,
	}

	users := []user{u1, u2, u3}

	err := tmpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}
}
