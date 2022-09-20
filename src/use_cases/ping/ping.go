package ping

import (
	"log"
	"net/http"
	"time"

	data "github.com/moohbr/WebMonitor/src/data"
	database "github.com/moohbr/WebMonitor/src/infrastructure/database"
	mail "github.com/moohbr/WebMonitor/src/use_cases/sendMail"
	templates "github.com/moohbr/WebMonitor/src/providers/mail/templates"
)

func PingServer(server data.Server) data.Server {
	resp, err := http.Get(server.URL)
	if err != nil {
		log.Fatal(err)
	}

	server.LastStatus = resp.StatusCode
	server.LastCheck = time.Now()
	db := database.OpenDatabase()

	defer db.Close()

	db.UpdateServer(server)

	if server.LastStatus != 200 {
		log.Println("[SYSTEM] Server " + server.Name + " is down")
		users := db.GetUsers()
		for _, user := range users {
			log.Println("[SYSTEM] Sending email to " + user.Email)
			mail.SendMail([]string{user.Email}, templates.ServerDown(server))
		}
	}
	return server
}

func PingAllServers() {
	db := database.OpenDatabase()

	defer db.Close()

	servers := db.GetServers()
	for _, server := range servers {
		PingServer(server)
	}
}
