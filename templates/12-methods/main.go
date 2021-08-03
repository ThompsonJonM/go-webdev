package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type person struct {
	Name string
	Age  int
}

func (p person) SomeProcessing() int {
	return 7
}

func (p person) DoubleAge() int {
	return p.Age * 2
}

func (p person) TakesArg(x int) int {
	return x * 2
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/**"))
}

func main() {
	p1 := person{"Jonathan Thompson", 34}

	err := tpl.ExecuteTemplate(os.Stdout, "body.gohtml", p1)
	if err != nil {
		log.Fatalln("Cannot execute file", err)
	}
}
