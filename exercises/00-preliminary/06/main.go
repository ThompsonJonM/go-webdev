package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Problem struct {
	Problem  string
	Solution string
}

func main() {
	ps := importProblems()
	var c, inc int

	for _, v := range ps {
		var ipt string
		fmt.Printf("What is the solution to: %s?\n", v.Problem)
		_, err := fmt.Scanln(&ipt)
		if err != nil {
			log.Fatalln("Could not read input", err)
		}

		if ipt != v.Solution {
			fmt.Printf("Incorrect. The correct answer is: %s\n", v.Solution)
			inc++
		} else {
			fmt.Println("Correct!")
			c++
		}
	}
	fmt.Printf("You answered %d correctly and %d incorrectly\n", c, inc)
}

func importProblems() []Problem {
	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalln("Could not open file", err)
	}
	rdr := csv.NewReader(file)
	rows, err := rdr.ReadAll()
	if err != nil {
		log.Fatalln("Could not read file", err)
	}

	var ps []Problem
	for _, v := range rows {
		ps = append(ps, Problem{
			Problem:  v[0],
			Solution: v[1],
		})
	}

	return ps
}
