package usecases

import (
	"fmt"

	email "github.com/moohbr/WebMonitor/src/providers/mail"
)

type Mail struct {
	Subject string
	Body    string
}

// Templates for the mails

var reportMail = Mail{
	Subject: "WebMonitor Report",
	Body:    "This is the report of the WebMonitor",
}

var TestMail = Mail{
	Subject: "Test Mail",
	Body:    "This is a test mail",
}

// Function to send the report mail
func SendMail(To []string, template Mail) {
	mail := email.NewMail(To, template.Subject, template.Body)
	mail.Send()
	fmt.Println("Mail sent!")
}
