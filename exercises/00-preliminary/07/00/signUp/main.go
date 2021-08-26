package main

import (
	"github.com/ThompsonJonM/go-webdev/exercises/00-preliminary/07/email"
)

func main() {
	b := `
Good morning %s

You are receiving this e-mail because you played Fantasy Football at Pendo in 2020. If you are
interested in playing for the 2021 season, head over to Slack channel #ffootball and sign up
using this sheet:

This is an automated e-mail created using Go.
`

	email.SendEmail("../../players", b)
}
