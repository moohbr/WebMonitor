package templates

import (
	data "github.com/moohbr/WebMonitor/src/data"
)

type Mail struct {
	Subject string
	Body    string
}

// Templates for the mails

var ReportMail = Mail{
	Subject: "WebMonitor Report",
	Body:    "This is the report of the WebMonitor",
}

var TestMail = Mail{
	Subject: "WebMonitor Mail",
	Body:    "This is a test mail",
}

var NewUserMail = Mail{
	Subject: "WebMonitor - Your account was created",
	Body:    "Congratulations! Your account was created.",
}

var ServerDown = func(server data.Server) Mail {
	return Mail{
		Subject: "WebMonitor - Server Down",
		Body:    "The server " + server.Name + " is down",
	}
}
