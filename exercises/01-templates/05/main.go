package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"text/template"
	"time"
)

type record struct {
	Date time.Time
	Open float64
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/body.gohtml"))
}

func main() {
	f, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln("Could not open file", err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalln("Could not close file", err)
		}
	}(f)

	r := csv.NewReader(f)
	rs, err := r.ReadAll()
	if err != nil {
		log.Fatalln("Could not read records", err)
	}

	var recs []record
	for i, rec := range rs {
		if i == 0 {
			continue
		}

		d, err := time.Parse("2006-01-02", rec[0])
		if err != nil {
			log.Fatalln("Could not parse date", err)
		}

		f, err := strconv.ParseFloat(rec[1], 64)
		if err != nil {
			log.Fatalln("Could not parse float", err)
		}

		recs = append(recs, record{
			Date: d,
			Open: f,
		})
	}

	err = tpl.Execute(os.Stdout, recs)
	if err != nil {
		log.Fatalln("Cannot execute file", err)
	}
}
