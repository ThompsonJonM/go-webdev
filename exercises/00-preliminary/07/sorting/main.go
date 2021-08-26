package main

import (
	"fmt"
	"github.com/ThompsonJonM/go-webdev/exercises/00-preliminary/07/importPlayers"
	"math/rand"
	"time"
)

type Player struct {
	Name  string
	Email string
}

func main() {
	ps := ImportPlayers()

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(ps), func(i, j int) {
		ps[i], ps[j] = ps[j], ps[i]
	})

	fmt.Println(ps)
}
