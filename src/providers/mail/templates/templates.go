package templates

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
