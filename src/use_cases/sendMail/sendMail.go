package sendMail

import (
	"log"

	email "github.com/moohbr/WebMonitor/src/providers/mail"
	templates "github.com/moohbr/WebMonitor/src/providers/mail/templates"
)

// Function to send the report mail
func SendMail(To []string, template templates.Mail) {
	mail := email.NewMail(To, template.Subject, template.Body)
	err := mail.Send()
	if err != nil {
		log.Println(err)
	}
	log.Println("SendMail - The email was sent!")
}
