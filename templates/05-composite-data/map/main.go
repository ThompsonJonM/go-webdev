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
	locs := map[string]string{
		"New York":       "Rochester",
		"Illinois":       "Chicago",
		"North Carolina": "Raleigh",
	}

	err := tpl.Execute(os.Stdout, locs)
	if err != nil {
		log.Fatalln("Cannot execute file", err)
	}
}
