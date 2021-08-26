package main

import (
	"fmt"
	"github.com/ThompsonJonM/go-webdev/exercises/00-preliminary/07/importPlayers"
	"log"
	"net/smtp"
)

type smtpServer struct {
	host string
	port string
}

func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

func main() {
	players := ImportPlayers()

	for _, v := range players {
		from := ""
		pass := ""
		to := []string{
			v.Email,
		}

		s := smtpServer{host: "smtp.gmail.com", port: "587"}
		ms := fmt.Sprintf(`
Good morning %s

You are receiving this e-mail because you played Fantasy Football at Pendo in 2020. If you are
interested in playing for the 2021 season, head over to Slack channel #ffootball and sign up
using this sheet:

This is an automated e-mail created using Go.
	`, v.Name)

		mess := []byte("Subject: Fantasy Football 2021" + ms)

		auth := smtp.PlainAuth("", from, pass, s.host)

		err := smtp.SendMail(s.Address(), auth, from, to, mess)
		if err != nil {
			log.Fatalln("Could not send e-mail", err)
		}

		fmt.Println("Email sent.")
	}
}
