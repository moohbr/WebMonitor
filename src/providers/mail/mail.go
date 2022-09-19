package mail

import (
	"crypto/tls"

	config "github.com/moohbr/WebMonitor/src/infrastructure/config"
	gomail "gopkg.in/mail.v2"
)

// Mail is the mail struct
type Mail struct {
	From    string
	To      []string
	Subject string
	Body    string
}

// NewMail creates a new mail
func NewMail(to []string, subject string, body string) *Mail {
	return &Mail{To: to, Subject: subject, Body: body}
}

// Send sends the mail
func (mail *Mail) Send() error {
	config.LoadEnv()
	message := gomail.NewMessage()

	message.SetHeader("From", config.GetEnv("SMTP_USER"))
	// Set E-Mail receivers
	message.SetHeader("To", mail.To...)

	// Set E-Mail subject
	message.SetHeader("Subject", mail.Subject)

	// Set E-Mail body. You can set plain text or html with text/html
	message.SetBody("text/plain", mail.Body)

	// Settings for SMTP server
	dialer := gomail.NewDialer(config.GetEnv("SMPT_SERVER"), config.ConvertToInt(config.GetEnv("SMTP_PORT")), config.GetEnv("SMTP_USER"), config.GetEnv("SMTP_PASSWORD"))

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := dialer.DialAndSend(message); err != nil {
		return err
	}

	return nil
}
