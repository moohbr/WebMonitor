package sendMail

import (
	"sync"

	email "github.com/moohbr/WebMonitor/src/providers/mail"
	templates "github.com/moohbr/WebMonitor/src/providers/mail/templates"
)

// Function to send the report mail
func SendMail(To []string, template templates.Mail, wg *sync.WaitGroup) {
	mail := email.NewMail(To, template.Subject, template.Body)
	mail.Send()
	wg.Done()
}
