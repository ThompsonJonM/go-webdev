package main

import (
	"github.com/ThompsonJonM/go-webdev/exercises/00-preliminary/07/email"
)

func main() {
	b := `
<html><body>
Good morning %s,
<br>
<br>
You are receiving this e-mail because you have been assigned to League 1 of the Pendo
2021 Fantasy Football season! Your league manager is: Jonathan Thompson. 
<br>
<br>
League rules are the following:
<li>PPR</li>
<li>6 teams make the playoffs</li>
<li>Playoffs are weeks 15, 16, and 17</li>
<li>Top 2 teams are given a BYE in week 15</li>
<br>
<b>This is an automated e-mail created using Go</b>.
</body></html>
`
	email.SendEmail("../../players.csv", b)
}
