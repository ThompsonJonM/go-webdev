package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	pi := math.Pi
	return pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func printArea(s shape) {
	fmt.Println(s.area())
}

func main() {
	s1 := square{
		side: 21.1,
	}

	c1 := circle{
		radius: 52.5,
	}

	fmt.Println(s1, c1)
	printArea(c1)
	printArea(s1)
}
