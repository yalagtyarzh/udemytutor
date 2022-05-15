package main

import (
	"html/template"
	"log"
	"os"
)

type person struct {
	Name string
	Age  int
}

func (p person) SomeProcessing() int {
	return 7
}

func (p person) AgeDbl() int {
	return p.Age * 2
}

func (p person) TakesArg(x int) int {
	return x * 2
}

var tmpl *template.Template

func main() {
	tmpl := template.Must(template.ParseGlob("./templates/*.tmpl"))
	p := person{
		Name: "Iam Fleming",
		Age:  56,
	}

	err := tmpl.Execute(os.Stdout, p)
	if err != nil {
		log.Fatalln(err)
	}
}
