package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

type location struct {
	State string
	City  string
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/**"))
	/*
		tpl = template.Must(template.ParseFiles("tpl.gohtml"))
		tpl = tpl.Funcs(fm)

		This will NOT work as functions need to be loaded BEFORE parsing
	*/
}

func main() {
	loc1 := location{"New York", "Rochester"}
	loc2 := location{"Illinois", "Chicago"}
	loc3 := location{"North Carolina", "Raleigh"}

	locs := []location{loc1, loc2, loc3}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", locs)
	if err != nil {
		log.Fatalln("Cannot execute file", err)
	}
}
