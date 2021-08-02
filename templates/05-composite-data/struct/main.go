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

	err := tpl.Execute(os.Stdout, loc1)
	if err != nil {
		log.Fatalln("Cannot execute file", err)
	}
}
