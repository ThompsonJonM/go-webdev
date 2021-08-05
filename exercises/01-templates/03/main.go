package main

import (
	"log"
	"os"
	"text/template"
)

type item struct {
	Name  string
	Price float64
}

type menu struct {
	Type  string
	Items []item
}

type restaurant struct {
	Breakfast, Lunch, Dinner menu
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/**"))
}

func main() {
	r := restaurant{
		Breakfast: menu{
			"Breakfast",
			[]item{
				{
					"Ham and Eggs",
					7.99,
				},
				{
					"Bacon and Egg Sandwich",
					6.50,
				},
			},
		},
		Lunch: menu{
			"Lunch",
			[]item{
				{
					"Italian Panini",
					9.99,
				},
				{
					"Garbage Plate",
					10.99,
				},
			},
		},
		Dinner: menu{
			"Dinner",
			[]item{
				{
					"NY Strip Steak",
					20.99,
				},
				{
					"Salmon Filet",
					17.99,
				},
			},
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "body.gohtml", r)
	if err != nil {
		log.Fatalln("Cannot execute file", err)
	}
}
