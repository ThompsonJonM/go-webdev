package main

import (
	"fmt"
)

func main() {
	x := []int{48, 46, 22, 10, 5, 2}

	lowest := x[0]
	for i, v := range x {
		if v <= x[i] {
			lowest = v
		}
	}

	fmt.Println(lowest)
}
