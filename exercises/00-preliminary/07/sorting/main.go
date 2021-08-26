package main

import (
	"fmt"
	"github.com/ThompsonJonM/go-webdev/exercises/00-preliminary/07/players"
	"math/rand"
	"time"
)

func main() {
	ps := players.ImportPlayers("../players.csv")

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(ps), func(i, j int) {
		ps[i], ps[j] = ps[j], ps[i]
	})

	fmt.Println(ps)
}
