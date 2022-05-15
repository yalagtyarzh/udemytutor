package main

import (
	"html/template"
	"log"
	"os"
)

type Regions []Region

type Region struct {
	Name   string
	Hotels []Hotel
}

type Hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region
}

var tmpl *template.Template

func main() {
	tmpl := template.Must(template.ParseGlob("./templates/*.tmpl"))

	h := Regions{
		Region{
			Name: "Southern",
			Hotels: []Hotel{
				Hotel{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     "95612",
				},
				Hotel{
					Name:    "H",
					Address: "4",
					City:    "L",
					Zip:     "95612",
				},
			},
		},
		Region{
			Name: "Central",
			Hotels: []Hotel{
				Hotel{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     "95612",
				},
				Hotel{
					Name:    "H",
					Address: "4",
					City:    "L",
					Zip:     "95612",
				},
			},
		},
		Region{
			Name: "Nortern",
			Hotels: []Hotel{
				Hotel{
					Name:    "Hotel California",
					Address: "42 Sunset Boulevard",
					City:    "Los Angeles",
					Zip:     "95612",
				},
				Hotel{
					Name:    "H",
					Address: "4",
					City:    "L",
					Zip:     "95612",
				},
			},
		},
	}

	err := tmpl.Execute(os.Stdout, h)
	if err != nil {
		log.Fatalln(err)
	}
}
