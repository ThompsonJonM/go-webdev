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

type car struct {
	Make  string
	Model string
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

	c1 := car{
		Make:  "Chevrolet",
		Model: "Lumina",
	}

	c2 := car{
		Make:  "MINI",
		Model: "Cooper",
	}

	c3 := car{
		Make:  "Scion",
		Model: "tC",
	}

	locs := []location{loc1, loc2, loc3}
	cars := []car{c1, c2, c3}
	items := struct {
		Locations []location
		Cars      []car
	}{
		Locations: locs,
		Cars:      cars,
	}

	err := tpl.Execute(os.Stdout, items)
	if err != nil {
		log.Fatalln("Cannot execute file", err)
	}
}
