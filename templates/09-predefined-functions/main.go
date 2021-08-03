package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type person struct {
	First string
	Last  string
	Age   int
	profession
}

type profession struct {
	Career string
	Degree bool
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/**"))
}

func main() {
	p1 := person{
		First: "Matt",
		Last:  "Barkley",
		Age:   42,
		profession: profession{
			Career: "Football Player",
			Degree: true,
		},
	}

	p2 := person{
		First: "Jonathan",
		Last:  "Thompson",
		Age:   34,
		profession: profession{
			Career: "Back End Engineer",
			Degree: false,
		},
	}

	p3 := person{
		First: "Joe",
		Last:  "Schmoe",
		Age:   27,
		profession: profession{
			Career: "",
			Degree: false,
		},
	}

	people := []person{p1, p2, p3}
	err := tpl.Execute(os.Stdout, people)
	if err != nil {
		log.Fatalln("Cannot execute file", err)
	}
}
