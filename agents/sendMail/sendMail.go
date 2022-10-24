package sendMail

import (
	email "github.com/moohbr/WebMonitor/providers/mail"
	"github.com/moohbr/WebMonitor/providers/mail/templates"
)

// Function to send the report mail
func SendMail(To []string, template templates.Mail) {
	mail := email.NewMail(To, template.Subject, template.Body)
	err := mail.Send()

	if err != nil {
		panic(err)
	}

}
