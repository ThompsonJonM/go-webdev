package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/**"))
}

func main() {
	locs := []string{"Rochester, NY", "Chicago, IL", "Raleigh, NC"}

	err := tpl.Execute(os.Stdout, locs)
	if err != nil {
		log.Fatalln("Cannot execute file", err)
	}
}
