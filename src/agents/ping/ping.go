package ping

import (
	"log"
	"net/http"
	"strconv"
	"time"

	mail "github.com/moohbr/WebMonitor/src/agents/sendMail"
	"github.com/moohbr/WebMonitor/src/data"
	"github.com/moohbr/WebMonitor/src/infrastructure/database"
	"github.com/moohbr/WebMonitor/src/providers/mail/templates"
)

func PingServer(server data.Server) {

	if server.Monitor {
		start := time.Now()
		response, err := http.Get("https://" + server.URL)
		if err != nil {
			log.Println("[ERROR] Error pinging server: " + server.Name)
			log.Println("[ERROR] Error: " + err.Error())
			server.LastStatus = 0
		} else {
			server.LastStatus = response.StatusCode
		}

		server.LastCheck = time.Now().UTC().Format("2006-01-02 15:04:05")
		server.AvarageResponseTime = time.Since(start)

		log.Println("[SYSTEM] " + server.Name + " | " + server.IP + " | " + server.URL + " | " +
			strconv.FormatInt(server.AvarageResponseTime.Milliseconds(), 10) + "ms | " + server.LastUpdate + " | " + server.LastCheck +
			" | " + strconv.Itoa(server.LastStatus) + " | " + strconv.FormatBool(server.Monitor))

		if server.LastStatus == 200 {
			return
		}

		db := database.OpenDatabase()
		defer db.Close()

		db.UpdateServer(server)
		users := db.GetUsers()

		if len(users) == 0 {
			return
		}

		list_of_emails := []string{}

		for _, user := range users {
			list_of_emails = append(list_of_emails, user.Email)
		}

		mail.SendMail(list_of_emails, templates.ServerDown(server))

	}
}

func PingAllServers() {
	db := database.OpenDatabase()

	defer db.Close()

	servers := db.GetServers()
	for _, server := range servers {
		PingServer(server)
	}
}
