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
	err := tpl.Execute(os.Stdout, 42) // Can ONLY pass in a single element of data
	if err != nil {
		log.Fatalln("Cannot execute file", err)
	}
}
