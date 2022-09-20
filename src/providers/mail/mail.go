package mail

import (
	"crypto/tls"

	config "github.com/moohbr/WebMonitor/src/infrastructure/config"
	gomail "gopkg.in/mail.v2"
)

type Mail struct {
	From    string
	To      []string
	Subject string
	Body    string
}

func NewMail(to []string, subject string, body string) *Mail {
	return &Mail{To: to, Subject: subject, Body: body}
}

func (mail *Mail) Send() error {
	config.LoadEnv()
	message := gomail.NewMessage()

	message.SetHeader("From", config.GetEnv("SMTP_USER"))
	message.SetHeader("To", mail.To...)

	message.SetHeader("Subject", mail.Subject)

	message.SetBody("text/plain", mail.Body)

	dialer := gomail.NewDialer(config.GetEnv("SMPT_SERVER"), config.ConvertToInt(config.GetEnv("SMTP_PORT")),
		config.GetEnv("SMTP_USER"), config.GetEnv("SMTP_PASSWORD"))

	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dialer.DialAndSend(message); err != nil {
		panic(err)
	}

	return nil
}
