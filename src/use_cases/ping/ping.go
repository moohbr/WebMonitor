package ping

import (
	"log"
	"net/http"
	"time"

	data "github.com/moohbr/WebMonitor/src/data"
	database "github.com/moohbr/WebMonitor/src/infrastructure/database"
	mail "github.com/moohbr/WebMonitor/src/use_cases/sendMail"
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
			sendMail.SendEmail(user.Email, "Server "+server.Name+" is down", "The server "+server.Name+" is down. Please check it.")
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
