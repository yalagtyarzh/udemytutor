package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

var tmpl *template.Template

var fm = template.FuncMap{
	"fdateMDY": monthDayYear,
}

func monthDayYear(t time.Time) string {
	return t.Format(time.Kitchen)
}

func main() {
	tmpl := template.Must(template.New("").Funcs(fm).ParseGlob("./templates/*.tmpl"))
	err := tmpl.ExecuteTemplate(os.Stdout, "tmpl.tmpl", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
