package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	license bool
}

func (p person) speak() {
	fmt.Println(p.first, p.last, `says "Hello."`)
}

func (sa secretAgent) speak() {
	if sa.license != true {
		fmt.Println(sa.first, sa.last, `says "I do not have a license to kill."`)
	} else {
		fmt.Println(sa.first, sa.last, `says "I have a license to kill."`)
	}
}

func main() {
	p1 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		license: true,
	}

	p1.speak()
	sa1.speak()
}
