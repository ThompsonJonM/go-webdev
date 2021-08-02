package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("templates/tpl.gohtml") // Returns a pointer to a template
	if err != nil {
		log.Fatalln("Cannot parse file", err)
	}

	err = tpl.Execute(os.Stdout, nil) // Consumes a template pointer and writes to a source using a writer
	if err != nil {
		log.Fatalln("Cannot execute", err)
	}

	tpl, err = template.ParseFiles("templates/tpl.gohtml", "templates/tpl-two.gohtml")
	if err != nil {
		log.Fatalln("Cannot parse files", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "tpl-two.gohtml", nil)
	if err != nil {
		log.Fatalln("Cannot execute specific file", err)
	}

	tpl, err = template.ParseGlob("templates/**")
	if err != nil {
		log.Fatalln("Cannot parse glob", err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "tpl-two.gohtml", nil)
	if err != nil {
		log.Fatalln("Cannot execute specific file", err)
	}

}
