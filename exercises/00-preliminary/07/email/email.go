package email

import (
	"fmt"
	"github.com/ThompsonJonM/go-webdev/exercises/00-preliminary/07/players"
	"log"
	"net/smtp"
	"strings"
)

type smtpServer struct {
	host string
	port string
}

func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

func getMessageString(fromEmail, To, Subject, emailBody string) []byte {
	return []byte("From: " + fromEmail + "\r\n" + "To: " + To + "\r\n" + "Subject: " + Subject + "\r\n" + "MIME-Version: 1.0\r\n" + "Content-Type: text/html; charset=\"utf-8\"\r\n\r\n" + emailBody + "\r\n")
}

func SendEmail(from, pass, file, subject, body string) {
	ps := players.ImportPlayers(file)

	var i int

	for _, v := range ps {
		to := []string{
			v.Email,
		}
		f := strings.Split(v.Name, " ")
		first := strings.Join(f[:1], "")

		s := smtpServer{host: "smtp.gmail.com", port: "587"}
		ms := fmt.Sprintf(body, first)

		b := getMessageString(from, to[0], subject, ms)

		auth := smtp.PlainAuth("", from, pass, s.host)

		err := smtp.SendMail(s.Address(), auth, from, to, b)
		if err != nil {
			log.Fatalln("Could not send e-mail", err)
		}

		i++

		fmt.Printf("%d of %d emails sent.\n", i, len(ps))
	}
}
