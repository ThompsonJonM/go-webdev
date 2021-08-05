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
	Name  string
	Menus []menu
}

type restaurants []restaurant

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/**"))
}

func main() {
	r1 := restaurant{
		"Winston's Cafe",
		[]menu{
			{
				"Breakfast",
				[]item{
					{
						"Ham and Eggs",
						7.99,
					},
				},
			},
			{
				"Lunch",
				[]item{
					{
						"Peanut Butter Burger",
						9.99,
					},
				},
			},
			{
				"Dinner",
				[]item{
					{
						"Salmon Filet",
						17.99,
					},
				},
			},
		},
	}

	r2 := restaurant{
		"Nicoles's Kitchen",
		[]menu{
			{
				"Breakfast",
				[]item{
					{
						"Ham and Eggs",
						7.99,
					},
				},
			},
			{
				"Lunch",
				[]item{
					{
						"Caprese Salad",
						7.99,
					},
				},
			},
			{
				"Dinner",
				[]item{
					{
						"Lobster",
						27.99,
					},
				},
			},
		},
	}

	rs := restaurants{r1, r2}

	err := tpl.ExecuteTemplate(os.Stdout, "body.gohtml", rs)
	if err != nil {
		log.Fatalln("Cannot execute file", err)
	}
}
