package templates

import (
	data "github.com/moohbr/WebMonitor/src/data"
)

type Mail struct {
	Subject string
	Body    string
}

// Templates for the mails

var ServerDown = func(server data.Server) Mail {
	return Mail{
		Subject: "WebMonitor - Server Down",
		Body:    "The server " + server.Name + " is down",
	}
}

var ServerUp = func(server data.Server) Mail {
	return Mail{
		Subject: "WebMonitor - Server Up",
		Body:    "The server " + server.Name + " is up",
	}
}

var NewServer = func(server data.Server) Mail {
	return Mail{
		Subject: "WebMonitor - New Server",
		Body:    "The server " + server.Name + " was added",
	}
}
