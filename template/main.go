package main

import (
	"html/template"
	"log"
	"os"
)

type Food struct {
	Name        string
	Description string
	Price       float64
}

type Meal struct {
	Meal  string
	Foods []Food
}

type Menu []Meal

var tmpl *template.Template

func main() {
	tmpl := template.Must(template.ParseGlob("./templates/*.tmpl"))

	m := Menu{
		Meal{
			Meal: "Breakfast",
			Foods: []Food{
				Food{
					Name:        "Oatmeal",
					Description: "yum yum",
					Price:       4.95,
				},
				Food{
					Name:        "Cheerios",
					Description: "American eating food traditional now",
					Price:       3.95,
				},
				Food{
					Name:        "Juice Orange",
					Description: "Delicious drinking in throat squeezed fresh",
					Price:       2.95,
				},
			},
		},
		Meal{
			Meal: "Lunch",
			Foods: []Food{
				Food{
					Name:        "Hamburger",
					Description: "Delicous good eating for you",
					Price:       6.95,
				},
				Food{
					Name:        "Cheese Melted Sandwich",
					Description: "Make cheese bread melt grease hot",
					Price:       3.95,
				},
				Food{
					Name:        "French Fries",
					Description: "French eat potatoe fingers",
					Price:       2.95,
				},
			},
		},
		Meal{
			Meal: "Dinner",
			Foods: []Food{
				Food{
					Name:        "Pasta Bolognese",
					Description: "From Italy delicious eating",
					Price:       7.95,
				},
				Food{
					Name:        "Steak",
					Description: "Dead cow grilled bloody",
					Price:       13.95,
				},
				Food{
					Name:        "Bistro Potatoe",
					Description: "Bistro bar wood American bacon",
					Price:       6.95,
				},
			},
		},
	}

	err := tmpl.Execute(os.Stdout, m)
	if err != nil {
		log.Fatalln(err)
	}
}
