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
	err := tpl.Execute(os.Stdout, "Embrace focus.")
	if err != nil {
		log.Fatalln("Cannot execute file", err)
	}
}
