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
	err := tpl.ExecuteTemplate(os.Stdout, "body.gohtml", "Hello Footer!")
	if err != nil {
		log.Fatalln("Cannot execute file", err)
	}
}
