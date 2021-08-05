package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type hotel struct {
	Name, Address, City string
	Zip                 int
}

type region struct {
	Location string
	Hotels   []hotel
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/**"))
}

func main() {
	r1 := region{
		Location: "Upstate",
		Hotels: []hotel{
			{
				Name:    "Econolodge",
				Address: "123 Jump Street",
				City:    "Rochester",
				Zip:     14610,
			},
			{
				Name:    "Cadillac Hotel",
				Address: "256 Chestnut Street",
				City:    "Rochester",
				Zip:     14628,
			},
		},
	}

	r2 := region{
		Location: "Adirondacks",
		Hotels: []hotel{
			{
				Name:    "Adirondack Lodge",
				Address: "4000 Adirondack Way",
				City:    "Lake Placid",
				Zip:     12946,
			},
			{
				Name:    "Harden Haus",
				Address: "145 Miracle Drive",
				City:    "Lake Placid",
				Zip:     12946,
			},
		},
	}

	r3 := region{
		Location: "Hudson Valley",
		Hotels: []hotel{
			{
				Name:    "Sleepy Hollow Lodge",
				Address: "40 Ichabod Lane",
				City:    "Sleepy Hollow",
				Zip:     10591,
			},
			{
				Name:    "Clovis Hotel",
				Address: "5026 Pumpkin Avenue",
				City:    "Sleepy Hollow",
				Zip:     10591,
			},
		},
	}

	regions := []region{r1, r2, r3}
	err := tpl.ExecuteTemplate(os.Stdout, "body.gohtml", regions)
	if err != nil {
		log.Fatalln("Cannot execute file", err)
	}
}
