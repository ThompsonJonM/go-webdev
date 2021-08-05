package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"text/template"
	"time"
)

var tpl *template.Template

type record struct {
	Date time.Time
	Open float64
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/**"))
}

func main() {
	data := prs("table.csv")
	err := tpl.ExecuteTemplate(os.Stdout, "body.gohtml", data)
	if err != nil {
		log.Fatalln("Cannot execute template", err)
	}
}

func prs(f string) []record {
	src, err := os.Open(f)
	if err != nil {
		log.Fatalln("Could not open file", err)
	}
	defer func(src *os.File) {
		err := src.Close()
		if err != nil {
			log.Fatalln("Could not close file", err)
		}
	}(src)

	rdr := csv.NewReader(src)
	rec, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln("Could not read file", err)
	}

	var r record
	var recs []record

	for i, v := range rec {
		if i == 0 {
			continue
		}

		date, _ := time.Parse("2006-01-02", v[0])
		open, _ := strconv.ParseFloat(v[1], 64)
		r = record{
			Date: date,
			Open: open,
		}

		recs = append(recs, r)
	}

	return recs
}
