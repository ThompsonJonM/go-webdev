package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type course struct {
	Number, Name string
	Units        int
}

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/**"))
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{
				{
					Number: "HIS 100",
					Name:   "Intro to History",
					Units:  4,
				},
				{
					Number: "HIS 202",
					Name:   "Historiography",
					Units:  4,
				},
				{
					Number: "HIS 304",
					Name:   "China in the Modern Age",
					Units:  4,
				},
			},
		},
		Spring: semester{
			Term: "Spring",
			Courses: []course{
				{
					"HIS 340",
					"Imperial China",
					4,
				},
				{
					"HIS 230",
					"American History",
					4,
				},
			},
		},
		Summer: semester{
			Term: "Summer",
			Courses: []course{
				{
					"ANT 100",
					"Intro to Anthropology",
					4,
				},
			},
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "body.gohtml", y)
	if err != nil {
		log.Fatalln("Cannot execute file", err)
	}
}
