package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type location struct {
	State string
	City  string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/**"))
}

func main() {
	loc1 := location{
		State: "New York",
		City:  "Rochester",
	}

	loc2 := location{
		State: "Illinois",
		City:  "Chicago",
	}

	loc3 := location{
		State: "North Carolina",
		City:  "Raleigh",
	}

	locs := []location{loc1, loc2, loc3}

	err := tpl.Execute(os.Stdout, locs)
	if err != nil {
		log.Fatalln("Cannot execute file", err)
	}
}
