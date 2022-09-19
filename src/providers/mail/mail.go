package mail

import (
	"fmt"
	"log"
	"net/smtp"
	"strings"
)

// Mail is the mail struct
type Mail struct {
	From    string
	To      []string
	Subject string
	Body    string
}

// NewMail creates a new mail
func NewMail(from string, to []string, subject string, body string) *Mail {
	return &Mail{from, to, subject, body}
}

// Send sends the mail
func (m *Mail) Send() {
	msg := "From: " + m.From + "\n" +
		"To: " + strings.Join(m.To, ", ") + "\n" +
		"Subject: " + m.Subject + "\n\n" +
		m.Body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", " ", " ", "smtp.gmail.com"),
		m.From, m.To, []byte(msg))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mail sent!")
}
